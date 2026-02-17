package filter

import (
	"strings"
	"tango/pkg/entity"
	"tango/pkg/services/config"
)

// AssetFilter filters access log records by asset URI patterns.
type AssetFilter struct {
	assetFilters []string
}

// NewAssetFilter creates a new AssetFilter from the given filter config.
func NewAssetFilter(filterConfig config.FilterConfig) *AssetFilter {
	return &AssetFilter{
		assetFilters: filterConfig.AssetFilters,
	}
}

// Filter returns true if the access log record matches an asset pattern.
func (f *AssetFilter) Filter(accessLogRecord entity.AccessLogRecord) bool {
	if len(f.assetFilters) == 0 {
		return false
	}

	uri := accessLogRecord.URI

	for _, assetPattern := range f.assetFilters {
		if strings.HasPrefix(uri, assetPattern) {
			return true
		}
	}

	return false
}
