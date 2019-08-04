package filter

import (
	"strings"
	"tango/internal/domain/entity"
	"tango/internal/usecase/config"
)

//
type UrlFilter struct {
	uriFilters     []string
	keepUriFilters []string
}

//
func NewUrlFilter(filterConfig config.FilterConfig) *UrlFilter {
	uriFilters := filterConfig.UriFilters
	keepUriFilters := filterConfig.KeepUriFilters

	return &UrlFilter{
		uriFilters:     uriFilters,
		keepUriFilters: keepUriFilters,
	}
}

//
func (f *UrlFilter) Filter(accessLogRecord entity.AccessLogRecord) bool {
	if len(f.uriFilters) == 0 && len(f.keepUriFilters) == 0 {
		return false
	}

	uri := accessLogRecord.URI

	// if keep filter is enabled, than keep only specified
	if len(f.keepUriFilters) > 0 {
		for _, urlPart := range f.keepUriFilters {
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
