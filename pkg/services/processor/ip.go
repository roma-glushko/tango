package processor

import (
	"net"
	"strings"
	"tango/pkg/domain/entity"
	"tango/pkg/services/config"
)

//
type IPProcessor struct {
	systemIPs       []string
	systemIPSubnets []*net.IPNet
}

//
func NewIPProcessor(processorConfig config.ProcessorConfig) IPProcessor {
	systemIpPatterns := processorConfig.SystemIpList

	systemIPs := make([]string, 0)
	systemIPSubnets := make([]*net.IPNet, 0)

	for _, ipPattern := range systemIpPatterns {
		// IP subnet pattern
		if strings.Contains(ipPattern, "/") {
			_, systemIpNet, _ := net.ParseCIDR(ipPattern)
			systemIPSubnets = append(systemIPSubnets, systemIpNet)
			continue
		}

		// single IP pattern
		systemIPs = append(systemIPs, ipPattern)
	}

	return IPProcessor{
		systemIPs:       systemIPs,
		systemIPSubnets: systemIPSubnets,
	}
}

// Process list of parsed IPs for access log record and remove system IPs
func (f *IPProcessor) Process(accessLogRecord entity.AccessLogRecord) entity.AccessLogRecord {
	if len(f.systemIPs) == 0 && len(f.systemIPSubnets) == 0 {
		return accessLogRecord
	}

	ipList := make([]string, 0)

	// filter system IPs
	for _, accessLogIp := range accessLogRecord.IP {
		filtered := false
		ip := net.ParseIP(accessLogIp)

		// check ip subnet patterns
		// goes first as potencially covers more IPs than singe IP pattern
		for _, ipSubnet := range f.systemIPSubnets {
			if ipSubnet.Contains(ip) {
				filtered = true
				break
			}
		}

		// check single ip patterns
		if !filtered {
			for _, systemIP := range f.systemIPs {
				if accessLogIp == systemIP {
					filtered = true
					break
				}
			}
		}

		// was IP filtered during checks?
		if !filtered {
			ipList = append(ipList, accessLogIp)
		}
	}

	accessLogRecord.IP = ipList

	return accessLogRecord
}
