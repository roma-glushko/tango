package processor

import (
	"net"
	"tango/internal/domain/entity"
	"tango/internal/usecase/config"
)

//
type IPProcessor struct {
	systemIPs []string
}

//
func NewIPProcessor(processorConfig config.ProcessorConfig) IPProcessor {
	return IPProcessor{
		systemIPs: processorConfig.SystemIpList,
	}
}

// Process list of parsed IPs for access log record and remove system IPs
func (f *IPProcessor) Process(accessLogRecord entity.AccessLogRecord) entity.AccessLogRecord {
	if len(f.systemIPs) == 0 {
		return accessLogRecord
	}

	ipList := make([]string, 0)

	// filter system IPs
	for _, accessLogIp := range accessLogRecord.IP {
		filtered := false
		ip := net.ParseIP(accessLogIp)

		for _, ipPattern := range f.systemIPs {
			// it's possible to specify subnet and test if current IP belongs to needed subnet
			_, systemIpNet, _ := net.ParseCIDR(ipPattern)

			if systemIpNet.Contains(ip) {
				filtered = true
				break
			}
		}

		if !filtered {
			ipList = append(ipList, accessLogIp)
		}
	}

	accessLogRecord.IP = ipList

	return accessLogRecord
}
