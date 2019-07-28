package filter

import (
	"tango/internal/domain/entity"
)

//
type UserAgentFilter struct {
}

//
func NewUserAgentFilter() *UserAgentFilter {
	return &UserAgentFilter{}
}

//
func (f *UserAgentFilter) Filter(accessLogRecord entity.AccessLogRecord) bool {
	//ipList := accessLogRecord.IP

	//for _, ip := range ipList {
	//if  {
	// ip filter logic
	//}
	//}

	return false
}
