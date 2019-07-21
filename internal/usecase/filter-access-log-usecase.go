package usecase

import "tango/internal/domain/entity"

//
type AccessLogFilter interface {
}

//
type FilterAccessLogUsecase struct {
	filters map[string][]AccessLogFilter
}

//
func NewAccessLogUsecase(accessLogFilters map[string][]AccessLogFilter) *FilterAccessLogUsecase {
	return &FilterAccessLogUsecase{filters: accessLogFilters}
}

//
func (u *FilterAccessLogUsecase) Filter(accessLogRecords []entity.AccessLogRecord) []entity.AccessLogRecord {

	return accessLogRecords
}
