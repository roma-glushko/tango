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

	ipList := accessLogRecord.SplitIPs()
	filtered := make([]string, 0, len(ipList))

	for _, accessLogIP := range ipList {
		isSystem := false
		ip := net.ParseIP(accessLogIP)

		// check ip subnet patterns
		for _, ipSubnet := range f.systemIPSubnets {
			if ipSubnet.Contains(ip) {
				isSystem = true
				break
			}
		}

		// check single ip patterns
		if !isSystem {
			if _, ok := f.systemIPs[accessLogIP]; ok {
				isSystem = true
			}
		}

		if !isSystem {
			filtered = append(filtered, accessLogIP)
		}
	}

	accessLogRecord.IP = strings.Join(filtered, " ")

	return accessLogRecord
}
