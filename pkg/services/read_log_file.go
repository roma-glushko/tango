package services

import (
	"sync"
	"tango/pkg/entity"
	"tango/pkg/services/config"
	"tango/pkg/services/mapper"
	"tango/pkg/services/processor"
)

// ReadAccessLogFunc is a callback function for processing access log lines.
type ReadAccessLogFunc func(accessLogRecord string, bytes int)

// AccessLogReader defines the interface for reading access log files.
type AccessLogReader interface {
	Read(filePath string, readAccessLogFunc ReadAccessLogFunc)
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

// Read parses access logs and streams AccessLogRecords through a channel.
func (u *ReadAccessLogService) Read(filePath string) <-chan entity.AccessLogRecord {
	numWorkers := u.pipelineConfig.Workers
	if numWorkers <= 1 {
		return u.readSequential(filePath)
	}

	return u.readConcurrent(filePath, numWorkers)
}

// readSequential preserves the original single-threaded behavior.
func (u *ReadAccessLogService) readSequential(filePath string) <-chan entity.AccessLogRecord {
	out := make(chan entity.AccessLogRecord, 1024)

	go func() {
		defer close(out)

		u.accessLogReader.Read(filePath, func(accessLogRecord string, bytes int) {
			record, err := u.parser(accessLogRecord)
			if err != nil {
				return
			}

			record = u.ipProcessor.Process(record)

			if u.filterAccessLogService.Filter(record) {
				return
			}

			out <- record
		})
	}()

	return out
}

const lineBatchSize = 256
const resultBatchSize = 256

// readConcurrent uses a fan-out/fan-in pipeline with goroutines and channels.
// Both input lines and output records are batched to reduce channel contention.
func (u *ReadAccessLogService) readConcurrent(filePath string, numWorkers int) <-chan entity.AccessLogRecord {
	linesCh := make(chan []string, numWorkers*4)
	batchesCh := make(chan []entity.AccessLogRecord, numWorkers*4)
	out := make(chan entity.AccessLogRecord, 1024)

	// Stage 1: Reader goroutine - reads file and sends batches of raw lines to workers
	go func() {
		defer close(linesCh)
		batch := make([]string, 0, lineBatchSize)

		u.accessLogReader.Read(filePath, func(line string, bytes int) {
			batch = append(batch, line)
			if len(batch) >= lineBatchSize {
				linesCh <- batch
				batch = make([]string, 0, lineBatchSize)
			}
		})

		if len(batch) > 0 {
			linesCh <- batch
		}
	}()

	// Stage 2: Worker goroutines - parse, process, filter, and send result batches
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
						batchesCh <- results
						results = make([]entity.AccessLogRecord, 0, resultBatchSize)
					}
				}
			}

			if len(results) > 0 {
				batchesCh <- results
			}
		}()
	}

	// Close batches channel when all workers finish
	go func() {
		workersWg.Wait()
		close(batchesCh)
	}()

	// Stage 3: Unbatcher - flatten batches into individual records for consumers
	go func() {
		defer close(out)
		for batch := range batchesCh {
			for _, record := range batch {
				out <- record
			}
		}
	}()

	return out
}
