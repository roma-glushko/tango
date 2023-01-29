package report

import (
	"sync"
	"tango/pkg/entity"
)

// GeoData Location Data
type GeoData struct {
	Country   string
	City      string
	Continent string
}

// GeoRecord report information
type GeoRecord struct {
	GeoData       *GeoData
	SampleRequest string
	BrowserAgent  string
	Requests      uint64
}

// GeoReport
type GeoReport struct {
	report map[string]*GeoRecord
	mu     sync.Mutex
}

// AddRequest
func (r *GeoReport) AddRequest(logRecord entity.AccessLogRecord, IP string, geoData *GeoData) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if geoRecord, found := r.report[IP]; found {
		geoRecord.Requests += 1
		return
	}

	r.report[IP] = &GeoRecord{
		GeoData:       geoData,
		SampleRequest: logRecord.URI,
		BrowserAgent:  logRecord.UserAgent,
		Requests:      1,
	}
}

func (r *GeoReport) Report() map[string]*GeoRecord {
	return r.report
}

// NewGeoReport
func NewGeoReport() *GeoReport {
	return &GeoReport{
		report: make(map[string]*GeoRecord),
		mu:     sync.Mutex{},
	}
}

// GeoLocationProvider provides ability to find geolocation data by IP
type GeoLocationProvider interface {
	GetGeoDataByIP(ip string) *GeoData // keep geoip2 references in the services layer, don't plan to support anything else
	Close() error
}

// GeoReportWriter is an interface for saving geo location reports
type GeoReportWriter interface {
	Save(reportPath string, geoReport *GeoReport)
}

// GeoReportService knows how to generate geo reports
type GeoReportService struct {
	geoLocationProvider GeoLocationProvider
	geoReportWriter     GeoReportWriter
}

// NewGeoReportService
func NewGeoReportService(geoLocationProvider GeoLocationProvider, geoReportWriter GeoReportWriter) *GeoReportService {
	return &GeoReportService{
		geoLocationProvider: geoLocationProvider,
		geoReportWriter:     geoReportWriter,
	}
}

// GenerateReport processes access logs and collect geo reports
func (u *GeoReportService) GenerateReport(reportPath string, logChan <-chan entity.AccessLogRecord) {
	geoReport := NewGeoReport()

	defer u.geoLocationProvider.Close()

	var waitGroup sync.WaitGroup

	for i := 0; i < 4; i++ {
		waitGroup.Add(1)

		go func() {
			defer waitGroup.Done()

			for accessRecord := range logChan {
				for _, ip := range accessRecord.IP {
					geoData := u.geoLocationProvider.GetGeoDataByIP(ip)
					geoReport.AddRequest(accessRecord, ip, geoData)
				}
			}
		}()
	}

	waitGroup.Wait()

	u.geoReportWriter.Save(reportPath, geoReport)
}
