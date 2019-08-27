package filter

import (
	"tango/internal/domain/entity"
	"tango/internal/usecase/config"
)

//
type IPFilter struct {
	ipFilters     []string
	keepIpFilters []string
}

//
func NewIPFilter(filterConfig config.FilterConfig) *IPFilter {
	return &IPFilter{
		ipFilters:     filterConfig.IpFilters,
		keepIpFilters: filterConfig.KeepIpFilters,
	}
}

// todo: remove duplicated code
func contains(ipList []string, ip string) bool {
	for _, ipItem := range ipList {
		if ipItem == ip {
			return true
		}
	}
	return false
}

//
func (f *IPFilter) Filter(accessLogRecord entity.AccessLogRecord) bool {
	if len(f.keepIpFilters) == 0 && len(f.ipFilters) == 0 {
		return false
	}

	ipList := accessLogRecord.IP

	// if keep filter is enabled, than keep only specified
	if len(f.keepIpFilters) > 0 {
		for _, keepIP := range f.keepIpFilters {
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
