package filter

import (
	"tango/internal/domain/entity"
)

var ip = []string{
	// todo ip filter data
}

//
type IPFilter struct {
}

//
func NewIPFilter() *IPFilter {
	return &IPFilter{}
}

//
func (f *IPFilter) Filter(accessLogRecord entity.AccessLogRecord) bool {
	//ipList := accessLogRecord.IP

	//for _, ip := range ipList {
	//if  {
	// ip filter logic
	//}
	//}

	return false
}
