package services

import (
	"tango/pkg/entity"
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
}

// NewReadAccessLogService creates a new ReadAccessLogService instance.
func NewReadAccessLogService(
	accessLogReader AccessLogReader,
	filterAccessLogService FilterAccessLogService,
	ipProcessor processor.IPProcessor,
) *ReadAccessLogService {
	return &ReadAccessLogService{
		accessLogReader:        accessLogReader,
		filterAccessLogService: filterAccessLogService,
		ipProcessor:            ipProcessor,
	}
}

// Read parses access logs and converts them to AccessLogRecord slices.
func (u *ReadAccessLogService) Read(filePath string) []entity.AccessLogRecord {
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
