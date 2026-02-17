package filter

import (
	"tango/pkg/entity"
	"tango/pkg/services/config"
)

// IPFilter filters access log records by IP address.
type IPFilter struct {
	ipFilters     []string
	keepIPFilters []string
}

// NewIPFilter creates a new IPFilter from the given filter config.
func NewIPFilter(filterConfig config.FilterConfig) *IPFilter {
	return &IPFilter{
		ipFilters:     filterConfig.IPFilters,
		keepIPFilters: filterConfig.KeepIPFilters,
	}
}

func contains(ipList []string, ip string) bool {
	for _, ipItem := range ipList {
		if ipItem == ip {
			return true
		}
	}
	return false
}

// Filter returns true if the access log record should be filtered by IP.
func (f *IPFilter) Filter(accessLogRecord entity.AccessLogRecord) bool {
	if len(f.keepIPFilters) == 0 && len(f.ipFilters) == 0 {
		return false
	}

	ipList := accessLogRecord.IP

	// if keep filter is enabled, than keep only specified
	if len(f.keepIPFilters) > 0 {
		for _, keepIP := range f.keepIPFilters {
			if contains(ipList, keepIP) {
				return false
			}
		}

		return true
	}

	// if keep filter is not enabled, then try to filter ips
	for _, ip := range f.ipFilters {
		if contains(ipList, ip) {
			return true
		}
	}

	return false
}
