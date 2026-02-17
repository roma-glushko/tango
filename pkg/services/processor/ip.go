package processor

import (
	"fmt"
	"net"
	"strings"
	"tango/pkg/entity"
	"tango/pkg/services/config"
)

// IPProcessor removes system IPs from access log records.
type IPProcessor struct {
	systemIPs       map[string]struct{}
	systemIPSubnets []*net.IPNet
}

// NewIPProcessor creates a new IPProcessor from the given processor config.
func NewIPProcessor(processorConfig config.ProcessorConfig) (IPProcessor, error) {
	systemIPPatterns := processorConfig.SystemIPList

	systemIPs := make(map[string]struct{})
	systemIPSubnets := make([]*net.IPNet, 0)

	for _, ipPattern := range systemIPPatterns {
		// IP subnet pattern
		if strings.Contains(ipPattern, "/") {
			_, systemIPNet, err := net.ParseCIDR(ipPattern)
			if err != nil {
				return IPProcessor{}, fmt.Errorf("invalid CIDR pattern %q: %w", ipPattern, err)
			}
			systemIPSubnets = append(systemIPSubnets, systemIPNet)
			continue
		}

		// single IP pattern
		if net.ParseIP(ipPattern) == nil {
			return IPProcessor{}, fmt.Errorf("invalid IP address %q", ipPattern)
		}
		systemIPs[ipPattern] = struct{}{}
	}

	return IPProcessor{
		systemIPs:       systemIPs,
		systemIPSubnets: systemIPSubnets,
	}, nil
}

// Process removes system IPs from the access log record's IP list.
func (f *IPProcessor) Process(accessLogRecord entity.AccessLogRecord) entity.AccessLogRecord {
	if len(f.systemIPs) == 0 && len(f.systemIPSubnets) == 0 {
		return accessLogRecord
	}

	ipList := make([]string, 0)

	// filter system IPs
	for _, accessLogIP := range accessLogRecord.IP {
		filtered := false
		ip := net.ParseIP(accessLogIP)

		// check ip subnet patterns
		// goes first as potentially covers more IPs than single IP pattern
		for _, ipSubnet := range f.systemIPSubnets {
			if ipSubnet.Contains(ip) {
				filtered = true
				break
			}
		}

		// check single ip patterns
		if !filtered {
			if _, ok := f.systemIPs[accessLogIP]; ok {
				filtered = true
			}
		}

		// was IP filtered during checks?
		if !filtered {
			ipList = append(ipList, accessLogIP)
		}
	}

	accessLogRecord.IP = ipList

	return accessLogRecord
}
