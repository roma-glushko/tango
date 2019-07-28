package usecase

import "tango/internal/domain/entity"

//
type AccessLogFilter interface {
	Filter(accessLog entity.AccessLogRecord) bool
}

//
type FilterAccessLogUsecase struct {
	filters []AccessLogFilter
}

//
func NewFilterAccessLogUsecase(accessLogFilters []AccessLogFilter) FilterAccessLogUsecase {
	return FilterAccessLogUsecase{filters: accessLogFilters}
}

//
func (u *FilterAccessLogUsecase) Filter(accessLogRecord entity.AccessLogRecord) bool {

	if len(u.filters) == 0 {
		return false
	}

	for _, filter := range u.filters {
		if filter.Filter(accessLogRecord) {
			return true
		}
	}

	return false
}
