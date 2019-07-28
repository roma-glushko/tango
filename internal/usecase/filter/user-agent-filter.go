package filter

import (
	"strings"
	"tango/internal/domain/entity"
)

var keepUserAgentList = []string{}

//
type UserAgentFilter struct {
}

//
func NewUserAgentFilter() *UserAgentFilter {
	return &UserAgentFilter{}
}

//
func (f *UserAgentFilter) Filter(accessLogRecord entity.AccessLogRecord) bool {
	if len(keepUserAgentList) == 0 {
		return false
	}

	userAgent := accessLogRecord.UserAgent

	for _, keepUserAgent := range keepUserAgentList {
		if !strings.Contains(userAgent, keepUserAgent) {
			return true
		}
	}

	return false
}
