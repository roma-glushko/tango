package filter

import (
	"strings"
	"tango/pkg/entity"
	"tango/pkg/services/config"
)

//
type AssetFilter struct {
	assetFilters []string
}

//
func NewAssetFilter(filterConfig config.FilterConfig) *AssetFilter {
	return &AssetFilter{
		assetFilters: filterConfig.AssetFilters,
	}
}

//
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
