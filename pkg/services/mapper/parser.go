package mapper

import (
	"fmt"
	"tango/pkg/entity"
)

// AccessLogParser parses a raw log line into an AccessLogRecord.
type AccessLogParser func(line []byte) (entity.AccessLogRecord, error)

// Supported log format names.
const (
	FormatApacheCombined = "apache-combined"
	FormatApacheCommon   = "apache-common"
	DefaultLogFormat     = FormatApacheCombined
)

// NewAccessLogParser returns a parser function for the given log format name.
func NewAccessLogParser(format string) (AccessLogParser, error) {
	switch format {
	case FormatApacheCombined:
		return ParseApacheCombined, nil
	case FormatApacheCommon:
		return ParseApacheCommon, nil
	default:
		return nil, fmt.Errorf("unsupported log format: %q (supported: %s, %s)", format, FormatApacheCombined, FormatApacheCommon)
	}
}
