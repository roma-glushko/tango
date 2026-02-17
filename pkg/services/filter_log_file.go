package services

import (
	"tango/pkg/entity"
)

// AccessLogFilter defines the interface for filtering access log records.
type AccessLogFilter interface {
	Filter(accessLog entity.AccessLogRecord) bool
}

// FilterAccessLogService applies a chain of filters to access log records.
type FilterAccessLogService struct {
	filters []AccessLogFilter
}

// NewFilterAccessLogService creates a new FilterAccessLogService with the given filters.
func NewFilterAccessLogService(accessLogFilters []AccessLogFilter) FilterAccessLogService {
	return FilterAccessLogService{filters: accessLogFilters}
}

// Filter returns true if any configured filter matches the access log record.
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
