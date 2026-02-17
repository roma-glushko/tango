package services

import (
	"bytes"
	"sync"
	"tango/pkg/entity"
	"tango/pkg/services/config"
	"tango/pkg/services/mapper"
	"tango/pkg/services/processor"
)

// ReadAccessLogFunc is a callback function for processing access log lines.
type ReadAccessLogFunc func(line []byte, n int)

// AccessLogReader defines the interface for reading access log files.
type AccessLogReader interface {
	Read(filePath string, readAccessLogFunc ReadAccessLogFunc)
}

// MappableReader can memory-map a file for direct zero-copy access.
type MappableReader interface {
	Map(filePath string) (data []byte, cleanup func(), err error)
}

// ReadAccessLogService reads access logs, processes and filters them.
type ReadAccessLogService struct {
	accessLogReader        AccessLogReader
	filterAccessLogService FilterAccessLogService
	ipProcessor            processor.IPProcessor
	pipelineConfig         config.PipelineConfig
	parser                 mapper.AccessLogParser
}

// NewReadAccessLogService creates a new ReadAccessLogService instance.
func NewReadAccessLogService(
	accessLogReader AccessLogReader,
	filterAccessLogService FilterAccessLogService,
	ipProcessor processor.IPProcessor,
	pipelineConfig config.PipelineConfig,
	parser mapper.AccessLogParser,
) *ReadAccessLogService {
	return &ReadAccessLogService{
		accessLogReader:        accessLogReader,
		filterAccessLogService: filterAccessLogService,
		ipProcessor:            ipProcessor,
		pipelineConfig:         pipelineConfig,
		parser:                 parser,
	}
}

// Read parses access logs and streams batches of AccessLogRecords through a merged channel.
func (u *ReadAccessLogService) Read(filePath string) <-chan []entity.AccessLogRecord {
	partitions := u.ReadPartitions(filePath)

	if len(partitions) == 1 {
		return partitions[0]
	}

	// Merge all partition channels into one
	out := make(chan []entity.AccessLogRecord, len(partitions)*4)
	var wg sync.WaitGroup
	for _, ch := range partitions {
		wg.Add(1)
		go func(c <-chan []entity.AccessLogRecord) {
			defer wg.Done()
			for batch := range c {
				out <- batch
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

const resultBatchSize = 256

// ReadPartitions returns one channel per worker for independent parallel consumption.
func (u *ReadAccessLogService) ReadPartitions(filePath string) []<-chan []entity.AccessLogRecord {
	numWorkers := u.pipelineConfig.Workers
	if numWorkers <= 1 {
		return []<-chan []entity.AccessLogRecord{u.readSequential(filePath)}
	}

	if mr, ok := u.accessLogReader.(MappableReader); ok {
		data, cleanup, err := mr.Map(filePath)
		if err == nil {
			return u.readPartitioned(data, cleanup, numWorkers)
		}
	}

	// Fallback: single merged channel
	return []<-chan []entity.AccessLogRecord{u.readChannelBased(filePath, numWorkers)}
}

// readSequential preserves the original single-threaded behavior.
func (u *ReadAccessLogService) readSequential(filePath string) <-chan []entity.AccessLogRecord {
	out := make(chan []entity.AccessLogRecord, 16)

	go func() {
		defer close(out)
		batch := make([]entity.AccessLogRecord, 0, resultBatchSize)

		u.accessLogReader.Read(filePath, func(line []byte, n int) {
			record, err := u.parser(line)
			if err != nil {
				return
			}

			record = u.ipProcessor.Process(record)

			if u.filterAccessLogService.Filter(record) {
				return
			}

			batch = append(batch, record)
			if len(batch) >= resultBatchSize {
				out <- batch
				batch = make([]entity.AccessLogRecord, 0, resultBatchSize)
			}
		})

		if len(batch) > 0 {
			out <- batch
		}
	}()

	return out
}

// readPartitioned splits mmap'd data into N chunks at newline boundaries.
// Returns one channel per worker — each worker parses its chunk independently.
func (u *ReadAccessLogService) readPartitioned(data []byte, cleanup func(), numWorkers int) []<-chan []entity.AccessLogRecord {
	channels := make([]<-chan []entity.AccessLogRecord, 0, numWorkers)

	chunkSize := len(data) / numWorkers

	var workersWg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize

		if i == numWorkers-1 {
			end = len(data)
		} else {
			nl := bytes.IndexByte(data[end:], '\n')
			if nl >= 0 {
				end += nl + 1
			} else {
				end = len(data)
			}
		}

		if i > 0 {
			nl := bytes.IndexByte(data[start:end], '\n')
			if nl >= 0 {
				start += nl + 1
			} else {
				continue
			}
		}

		ch := make(chan []entity.AccessLogRecord, 4)
		channels = append(channels, ch)

		workersWg.Add(1)
		go func(chunk []byte, out chan<- []entity.AccessLogRecord) {
			defer workersWg.Done()
			defer close(out)
			results := make([]entity.AccessLogRecord, 0, resultBatchSize)

			offset := 0
			for offset < len(chunk) {
				nl := bytes.IndexByte(chunk[offset:], '\n')
				var line []byte
				if nl < 0 {
					line = chunk[offset:]
					offset = len(chunk)
				} else {
					line = chunk[offset : offset+nl]
					offset += nl + 1
				}
				if len(line) > 0 && line[len(line)-1] == '\r' {
					line = line[:len(line)-1]
				}
				if len(line) == 0 {
					continue
				}

				record, err := u.parser(line)
				if err != nil {
					continue
				}

				record = u.ipProcessor.Process(record)

				if u.filterAccessLogService.Filter(record) {
					continue
				}

				results = append(results, record)
				if len(results) >= resultBatchSize {
					out <- results
					results = make([]entity.AccessLogRecord, 0, resultBatchSize)
				}
			}

			if len(results) > 0 {
				out <- results
			}
		}(data[start:end], ch)
	}

	// Release mmap after all workers finish
	go func() {
		workersWg.Wait()
		cleanup()
	}()

	return channels
}

// readChannelBased is the fallback fan-out/fan-in pipeline for non-mappable readers.
func (u *ReadAccessLogService) readChannelBased(filePath string, numWorkers int) <-chan []entity.AccessLogRecord {
	const lineBatchSize = 256
	linesCh := make(chan [][]byte, numWorkers*4)
	out := make(chan []entity.AccessLogRecord, numWorkers*4)

	go func() {
		defer close(linesCh)
		batch := make([][]byte, 0, lineBatchSize)

		u.accessLogReader.Read(filePath, func(line []byte, n int) {
			lineCopy := make([]byte, len(line))
			copy(lineCopy, line)
			batch = append(batch, lineCopy)
			if len(batch) >= lineBatchSize {
				linesCh <- batch
				batch = make([][]byte, 0, lineBatchSize)
			}
		})

		if len(batch) > 0 {
			linesCh <- batch
		}
	}()

	var workersWg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		workersWg.Add(1)
		go func() {
			defer workersWg.Done()
			results := make([]entity.AccessLogRecord, 0, resultBatchSize)

			for lines := range linesCh {
				for _, line := range lines {
					record, err := u.parser(line)
					if err != nil {
						continue
					}

					record = u.ipProcessor.Process(record)

					if u.filterAccessLogService.Filter(record) {
						continue
					}

					results = append(results, record)
					if len(results) >= resultBatchSize {
						out <- results
						results = make([]entity.AccessLogRecord, 0, resultBatchSize)
					}
				}
			}

			if len(results) > 0 {
				out <- results
			}
		}()
	}

	go func() {
		workersWg.Wait()
		close(out)
	}()

	return out
}
