package filter

import (
	"strings"
	"tango/internal/domain/entity"
)

var urlPartList = []string{}

//
type UrlFilter struct {
}

//
func NewUrlFilter() *UrlFilter {
	return &UrlFilter{}
}

//
func (f *UrlFilter) Filter(accessLogRecord entity.AccessLogRecord) bool {
	if len(urlPartList) == 0 {
		return false
	}

	uri := accessLogRecord.URI

	for _, urlPart := range urlPartList {
		if strings.Contains(uri, urlPart) {
			return true
		}
	}

	return false
}
