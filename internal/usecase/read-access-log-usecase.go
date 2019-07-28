package usecase

import (
	"tango/internal/domain/entity"
	"tango/internal/usecase/mapper"
)

//
type ReadAccessLogFunc func(accessLogRecord string, bytes int)

//
type AccessLogReader interface {
	Read(filePath string, readAccessLogFunc ReadAccessLogFunc)
}

//
type ReadAccessLogUsecase struct {
	accessLogReader        AccessLogReader
	filterAccessLogUsecase FilterAccessLogUsecase
}

// Create a new ReadAccessLogUsecase
func NewReadAccessLogUsecase(accessLogReader AccessLogReader, filterAccessLogUsecase FilterAccessLogUsecase) *ReadAccessLogUsecase {
	return &ReadAccessLogUsecase{
		accessLogReader:        accessLogReader,
		filterAccessLogUsecase: filterAccessLogUsecase,
	}
}

// Read Access Logs and convert them to array of AccessLogRecord-s
func (u *ReadAccessLogUsecase) Read(filePath string) []entity.AccessLogRecord {
	accessRecords := make([]entity.AccessLogRecord, 0)

	u.accessLogReader.Read(filePath, func(accessLogRecord string, bytes int) {
		accessRecord := mapper.MapAccessLogRecord(accessLogRecord)

		if u.filterAccessLogUsecase.Filter(accessRecord) {
			return
		}

		accessRecords = append(
			accessRecords,
			accessRecord,
		)
	})

	return accessRecords
}
