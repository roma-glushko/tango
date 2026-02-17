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
}

// NewReadAccessLogService creates a new ReadAccessLogService instance.
func NewReadAccessLogService(
	accessLogReader AccessLogReader,
	filterAccessLogService FilterAccessLogService,
	ipProcessor processor.IPProcessor,
	pipelineConfig config.PipelineConfig,
) *ReadAccessLogService {
	return &ReadAccessLogService{
		accessLogReader:        accessLogReader,
		filterAccessLogService: filterAccessLogService,
		ipProcessor:            ipProcessor,
		pipelineConfig:         pipelineConfig,
	}
}

// Read parses access logs and converts them to AccessLogRecord slices.
func (u *ReadAccessLogService) Read(filePath string) []entity.AccessLogRecord {
	numWorkers := u.pipelineConfig.Workers
	if numWorkers <= 1 {
		return u.readSequential(filePath)
	}

	return u.readConcurrent(filePath, numWorkers)
}

// readSequential preserves the original single-threaded behavior.
func (u *ReadAccessLogService) readSequential(filePath string) []entity.AccessLogRecord {
	accessRecords := make([]entity.AccessLogRecord, 0, 1024)

	u.accessLogReader.Read(filePath, func(accessLogRecord string, bytes int) {
		accessRecord, err := mapper.MapAccessLogRecord(accessLogRecord)
		if err != nil {
			// skip unparseable lines (e.g. malformed log entries)
			return
		}

		// process parsed access log record
		accessRecord = u.ipProcessor.Process(accessRecord)

		// filter/skip parsed access log record if needed
		if u.filterAccessLogService.Filter(accessRecord) {
			return
		}

		accessRecords = append(
			accessRecords,
			accessRecord,
		)
	})

	return accessRecords
}

// readConcurrent uses a fan-out/fan-in pipeline with goroutines and channels.
func (u *ReadAccessLogService) readConcurrent(filePath string, numWorkers int) []entity.AccessLogRecord {
	linesCh := make(chan string, numWorkers*64)
	resultsCh := make(chan entity.AccessLogRecord, numWorkers*64)

	// Stage 1: Reader goroutine - reads file and sends raw lines to workers
	go func() {
		defer close(linesCh)
		u.accessLogReader.Read(filePath, func(line string, bytes int) {
			linesCh <- line
		})
	}()

	// Stage 2: Worker goroutines - parse, process, and filter records
	var workersWg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		workersWg.Add(1)
		go func() {
			defer workersWg.Done()
			for line := range linesCh {
				record, err := mapper.MapAccessLogRecord(line)
				if err != nil {
					continue
				}

				record = u.ipProcessor.Process(record)

				if u.filterAccessLogService.Filter(record) {
					continue
				}

				resultsCh <- record
			}
		}()
	}

	// Close results channel when all workers finish
	go func() {
		workersWg.Wait()
		close(resultsCh)
	}()

	// Stage 3: Aggregator - collect processed records
	accessRecords := make([]entity.AccessLogRecord, 0, 1024)
	for record := range resultsCh {
		accessRecords = append(accessRecords, record)
	}

	return accessRecords
}
