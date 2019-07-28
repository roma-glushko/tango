package filter

import (
	"strings"
	"tango/internal/domain/entity"
)

var webAssetPatterns = []string{
	"/pub/static/",
	"/pub/media/",
}

//
type WebAssetFilter struct {
}

//
func NewWebAssetFilter() *WebAssetFilter {
	return &WebAssetFilter{}
}

//
func (f *WebAssetFilter) Filter(accessLogRecord entity.AccessLogRecord) bool {
	if len(webAssetPatterns) == 0 {
		return false
	}

	uri := accessLogRecord.URI

	for _, webAssetPattern := range webAssetPatterns {
		if strings.HasPrefix(uri, webAssetPattern) {
			return true
		}
	}

	return false
}
