package report

import (
	"fmt"
	"net/url"
	"regexp"
	"tango/pkg/entity"
)

// RequestReportItem represents a single request entry in the request report.
type RequestReportItem struct {
	Path         string
	Requests     uint64
	ResponseCode uint64
	RefererURLs  map[string]bool
}

// RequestReportWriter is an interface for saving request reports.
type RequestReportWriter interface {
	Save(reportPath string, browserReport map[string]*RequestReportItem)
}

// DefaultQueryStripPatterns are regex patterns stripped from request paths by default
var DefaultQueryStripPatterns = []string{
	"/key(.*)/",
	"/referer(.*)/",
}

// RequestReportService knows how to generate request reports.
type RequestReportService struct {
	queryStripPatterns  []string
	requestReportWriter RequestReportWriter
}

// NewRequestReportService creates a new RequestReportService instance.
func NewRequestReportService(requestReportWriter RequestReportWriter) *RequestReportService {
	return &RequestReportService{
		queryStripPatterns:  DefaultQueryStripPatterns,
		requestReportWriter: requestReportWriter,
	}
}

// GenerateReport processes access logs and collect request reports
func (u *RequestReportService) GenerateReport(reportPath string, accessRecords []entity.AccessLogRecord) {
	var requestReport = make(map[string]*RequestReportItem)

	pathFilters := make([]*regexp.Regexp, 0, len(u.queryStripPatterns))

	for _, pattern := range u.queryStripPatterns {
		filter, err := regexp.Compile(pattern)

		if err != nil {
			fmt.Println(err)
			continue
		}

		pathFilters = append(pathFilters, filter)
	}

	for _, accessRecord := range accessRecords {
		requestURI := accessRecord.URI
		refererURL := accessRecord.RefererURL

		parsedURI, err := url.Parse(requestURI)

		path := ""

		if err != nil {
			// during security scans it's possible to submit a request which triggers a panic in url.Parse()
			// in that case, just use the original URI
			path = requestURI
		} else {
			path = parsedURI.Path
		}

		for _, filter := range pathFilters {
			path = filter.ReplaceAllString(path, "")
		}

		if _, ok := requestReport[path]; ok {
			requestReport[path].Requests++

			// collect referer URLs
			if _, found := requestReport[path].RefererURLs[refererURL]; !found {
				requestReport[path].RefererURLs[refererURL] = true
			}

			continue
		}

		requestReport[path] = &RequestReportItem{
			Path:         path,
			Requests:     1,
			ResponseCode: accessRecord.ResponseCode,
			RefererURLs:  map[string]bool{refererURL: true},
		}
	}

	u.requestReportWriter.Save(reportPath, requestReport)
}
