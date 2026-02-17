package filter

import (
	"strings"
	"tango/pkg/entity"
	"tango/pkg/services/config"
)

// URLFilter filters access log records by URI patterns.
type URLFilter struct {
	uriFilters     []string
	keepURIFilters []string
}

// NewURLFilter creates a new URLFilter from the given filter config.
func NewURLFilter(filterConfig config.FilterConfig) *URLFilter {
	uriFilters := filterConfig.URIFilters
	keepURIFilters := filterConfig.KeepURIFilters

	return &URLFilter{
		uriFilters:     uriFilters,
		keepURIFilters: keepURIFilters,
	}
}

// Filter returns true if the access log record should be filtered by URI.
func (f *URLFilter) Filter(accessLogRecord entity.AccessLogRecord) bool {
	if len(f.uriFilters) == 0 && len(f.keepURIFilters) == 0 {
		return false
	}

	uri := accessLogRecord.URI

	// if keep filter is enabled, than keep only specified
	if len(f.keepURIFilters) > 0 {
		for _, urlPart := range f.keepURIFilters {
			if strings.Contains(uri, urlPart) {
				return false
			}
		}

		return true
	}

	// if keep filter is not enabled, then try to filter user agents
	for _, urlPart := range f.uriFilters {
		if strings.Contains(uri, urlPart) {
			return true
		}
	}

	return false
}
