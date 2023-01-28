package filter

import (
	"strings"
	"tango/pkg/domain/entity"
	"tango/pkg/services/config"
)

//
type UserAgentFilter struct {
	keepUaFilters []string
	uaFilters     []string
}

//
func NewUserAgentFilter(filterConfig config.FilterConfig) *UserAgentFilter {
	keepUaFilters := filterConfig.KeepUaFilters
	uaFilters := filterConfig.UaFilters

	return &UserAgentFilter{
		keepUaFilters: keepUaFilters,
		uaFilters:     uaFilters,
	}
}

//
func (f *UserAgentFilter) Filter(accessLogRecord entity.AccessLogRecord) bool {
	if len(f.keepUaFilters) == 0 && len(f.uaFilters) == 0 {
		return false
	}

	userAgent := accessLogRecord.UserAgent

	// if keep filter is enabled, than keep only specified
	if len(f.keepUaFilters) > 0 {
		for _, keepUserAgent := range f.keepUaFilters {
			if strings.Contains(userAgent, keepUserAgent) {
				return false
			}
		}

		return true
	}

	// if keep filter is not enabled, then try to filter user agents
	if len(f.uaFilters) > 0 {
		for _, userAgentFilter := range f.uaFilters {
			if strings.Contains(userAgent, userAgentFilter) {
				return true
			}
		}
	}

	return false
}
