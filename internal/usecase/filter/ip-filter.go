package filter

import (
	"tango/internal/domain/entity"
)

var keepIpList = []string{
	"127.0.0.1",
}

//
type IPFilter struct {
}

//
func NewIPFilter() *IPFilter {
	return &IPFilter{}
}

//
func contains(ipList []string, ip string) bool {
	for _, ipItem := range ipList {
		if ipItem == ip {
			return true
		}
	}
	return false
}

//
func (f *IPFilter) Filter(accessLogRecord entity.AccessLogRecord) bool {
	if len(keepIpList) == 0 {
		return false
	}

	ipList := accessLogRecord.IP

	for _, keepIP := range keepIpList {
		if contains(ipList, keepIP) {
			return true
		}
	}

	return false
}
