package report

import (
	"fmt"
	"net/url"
	"regexp"
	"tango/pkg/entity"
)

// RequestReportItem
type RequestReportItem struct {
	Path         string
	Requests     uint64
	ResponseCode uint64
	RefererURLs  map[string]bool
}

// RequestReportWriter
type RequestReportWriter interface {
	Save(reportPath string, browserReport map[string]*RequestReportItem)
}

// RequestReportService
type RequestReportService struct {
	requestReportWriter RequestReportWriter
}

//
func NewRequestReportService(requestReportWriter RequestReportWriter) *RequestReportService {
	return &RequestReportService{
		requestReportWriter: requestReportWriter,
	}
}

// GenerateReport processes access logs and collect request reports
func (u *RequestReportService) GenerateReport(reportPath string, accessRecords []entity.AccessLogRecord) {
	var requestReport = make(map[string]*RequestReportItem)

	// todo: move to configs
	queryPatterns := []string{
		"/key(.*)/",
		"/referer(.*)/",
	}

	// init additional query filters
	pathFilters := make([]*regexp.Regexp, 0)

	for _, pattern := range queryPatterns {
		filter, err := regexp.Compile(pattern)

		if err != nil {
			fmt.Println(err)
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
