package services

import "tango/pkg/domain/entity"

//
type AccessLogFilter interface {
	Filter(accessLog entity.AccessLogRecord) bool
}

//
type FilterAccessLogService struct {
	filters []AccessLogFilter
}

//
func NewFilterAccessLogService(accessLogFilters []AccessLogFilter) FilterAccessLogService {
	return FilterAccessLogService{filters: accessLogFilters}
}

//
func (u *FilterAccessLogService) Filter(accessLogRecord entity.AccessLogRecord) bool {

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
