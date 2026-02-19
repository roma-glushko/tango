package filter

import (
	"strings"
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

func containsIP(ipField string, ip string) bool {
	// Fast path: single IP (no space separator)
	if ipField == ip {
		return true
	}

	for _, part := range strings.Split(ipField, " ") {
		if part == ip {
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

	ip := accessLogRecord.IP

	// if keep filter is enabled, than keep only specified
	if len(f.keepIPFilters) > 0 {
		for _, keepIP := range f.keepIPFilters {
			if containsIP(ip, keepIP) {
				return false
			}
		}

		return true
	}

	// if keep filter is not enabled, then try to filter ips
	for _, filterIP := range f.ipFilters {
		if containsIP(ip, filterIP) {
			return true
		}
	}

	return false
}
