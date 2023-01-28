package services

import (
	"tango/pkg/domain/entity"
	"tango/pkg/services/mapper"
	"tango/pkg/services/processor"
)

//
type ReadAccessLogFunc func(accessLogRecord string, bytes int)

//
type AccessLogReader interface {
	Read(filePath string, readAccessLogFunc ReadAccessLogFunc)
}

//
type ReadAccessLogService struct {
	accessLogReader        AccessLogReader
	filterAccessLogService FilterAccessLogService
	ipProcessor            processor.IPProcessor
}

// NewReadAccessLogService Create a new ReadAccessLogService
func NewReadAccessLogService(accessLogReader AccessLogReader, filterAccessLogService FilterAccessLogService, ipProcessor processor.IPProcessor) *ReadAccessLogService {
	return &ReadAccessLogService{
		accessLogReader:        accessLogReader,
		filterAccessLogService: filterAccessLogService,
		ipProcessor:            ipProcessor,
	}
}

// Read Access Logs and convert them to array of AccessLogRecord-s
func (u *ReadAccessLogService) Read(filePath string) []entity.AccessLogRecord {
	accessRecords := make([]entity.AccessLogRecord, 0)

	u.accessLogReader.Read(filePath, func(accessLogRecord string, bytes int) {
		accessRecord := mapper.MapAccessLogRecord(accessLogRecord)

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
