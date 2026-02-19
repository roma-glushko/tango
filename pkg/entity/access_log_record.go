package entity

import (
	"strings"
	"time"
)

// AccessLogRecord Type of AccessLog record
type AccessLogRecord struct {
	IP            string
	URI           string
	Time          time.Time
	RequestMethod string
	Protocol      string
	ResponseCode  uint64
	ResponseSize  uint64
	UserAgent     string
	RefererURL    string
}

// SplitIPs splits the IP string into individual IP addresses.
func (r *AccessLogRecord) SplitIPs() []string {
	if r.IP == "" {
		return nil
	}

	if !strings.Contains(r.IP, " ") {
		return []string{r.IP}
	}

	parts := strings.Split(r.IP, " ")
	result := make([]string, 0, len(parts))
	for _, p := range parts {
		if p != "" {
			result = append(result, p)
		}
	}
	return result
}
