package filter

import (
	"strings"
	"tango/internal/domain/entity"
)

var assetPatterns = []string{}

//
type AssetFilter struct {
}

//
func NewAssetFilter() *AssetFilter {
	return &AssetFilter{}
}

//
func (f *AssetFilter) Filter(accessLogRecord entity.AccessLogRecord) bool {
	if len(assetPatterns) == 0 {
		return false
	}

	uri := accessLogRecord.URI

	for _, webAssetPattern := range assetPatterns {
		if strings.HasPrefix(uri, webAssetPattern) {
			return true
		}
	}

	return false
}
