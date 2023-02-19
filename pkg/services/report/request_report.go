package report

import (
	"fmt"
	"net/url"
	"regexp"
	"sync"
	"tango/pkg/entity"
	"tango/pkg/services/mapper"
)

// RequestReportItem
type RequestReportItem struct {
	Path         string
	Requests     uint64
	ResponseCode uint64
	RefererURLs  map[string]bool
}

type RequestReport struct {
	report map[string]*RequestReportItem
	mu     sync.Mutex
}

func (r *RequestReport) AddRequest(path string, refererURL string, logRecord entity.AccessLogRecord) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if requestRecord, exists := r.report[path]; exists {
		requestRecord.Requests++

		// collect referer URLs
		if _, found := requestRecord.RefererURLs[refererURL]; !found {
			requestRecord.RefererURLs[refererURL] = true
		}

		return
	}

	r.report[path] = &RequestReportItem{
		Path:         path,
		Requests:     1,
		ResponseCode: logRecord.ResponseCode,
		RefererURLs:  map[string]bool{refererURL: true},
	}
}

func (r *RequestReport) Report() map[string]*RequestReportItem {
	return r.report
}

func NewRequestReport() *RequestReport {
	return &RequestReport{
		report: make(map[string]*RequestReportItem),
		mu:     sync.Mutex{},
	}
}

// RequestReportWriter
type RequestReportWriter interface {
	Save(reportPath string, requestReport *RequestReport)
}

// RequestReportService
type RequestReportService struct {
	logMapper           *mapper.AccessLogMapper
	requestReportWriter RequestReportWriter
}

//
func NewRequestReportService(logMapper *mapper.AccessLogMapper, requestReportWriter RequestReportWriter) *RequestReportService {
	return &RequestReportService{
		logMapper:           logMapper,
		requestReportWriter: requestReportWriter,
	}
}

// GenerateReport processes access logs and collect request reports
func (s *RequestReportService) GenerateReport(reportPath string, logChan <-chan entity.AccessLogRecord) {
	requestReport := NewRequestReport()

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

	var waitGroup sync.WaitGroup

	for i := 0; i < 4; i++ {
		waitGroup.Add(1)

		go func() {
			defer waitGroup.Done()

			for accessRecord := range logChan {
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

				requestReport.AddRequest(path, refererURL, accessRecord)
				s.logMapper.Recycle(accessRecord)
			}
		}()
	}

	waitGroup.Wait()

	fmt.Println("💃 saving the request report...")
	s.requestReportWriter.Save(reportPath, requestReport)
}
