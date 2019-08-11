package report

import (
	"fmt"
	"net/url"
	"regexp"
	"tango/internal/domain/entity"
)

//
type RequestReportItem struct {
	Path     string
	Requests uint64
}

//
type RequestReportWriter interface {
	Save(reportPath string, browserReport map[string]*RequestReportItem)
}

//
type RequestReportUsecase struct {
	requestReportWriter RequestReportWriter
}

//
func NewRequestReportUsecase(requestReportWriter RequestReportWriter) *RequestReportUsecase {
	return &RequestReportUsecase{
		requestReportWriter: requestReportWriter,
	}
}

// Process access logs and collect request reports
func (u *RequestReportUsecase) GenerateReport(reportPath string, accessRecords []entity.AccessLogRecord) {
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

		parsedURI, err := url.Parse(requestURI)

		if err != nil {
			panic(err)
		}

		path := parsedURI.Path

		for _, filter := range pathFilters {
			path = filter.ReplaceAllString(path, "")
		}

		if _, ok := requestReport[path]; ok {
			requestReport[path].Requests++
			continue
		}

		requestReport[path] = &RequestReportItem{
			Path:     path,
			Requests: 1,
		}
	}

	u.requestReportWriter.Save(reportPath, requestReport)
}
