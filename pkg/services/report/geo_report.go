package report

import (
	"sync"
	"tango/pkg/entity"
)

// Geo Location Data
type GeoData struct {
	Country   string
	City      string
	Continent string
}

// Geolocation report information
type Geolocation struct {
	GeoData       *GeoData
	SampleRequest string
	BrowserAgent  string
	Requests      uint64
}

// GeoLocationProvider provides ability to find geolocation data by IP
type GeoLocationProvider interface {
	GetGeoDataByIP(ip string) *GeoData // keep geoip2 references in the services layer, don't plan to support anything else
	Close() error
}

// GeoReportWriter is an interface for saving geo location reports
type GeoReportWriter interface {
	Save(reportPath string, geolocationReport map[string]*Geolocation)
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
	var geoReport = make(map[string]*Geolocation)
	var mutex sync.Mutex      // TODO: try to use sync.Map
	var geoDBmutex sync.Mutex // TODO: try to use sync.Map

	defer u.geoLocationProvider.Close()

	var waitGroup sync.WaitGroup

	for i := 0; i < 4; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()

			for accessRecord := range logChan {
				for _, ip := range accessRecord.IP {

					if _, found := geoReport[ip]; found {
						mutex.Lock()
						geoReport[ip].Requests++
						mutex.Unlock()
						continue
					}

					geoDBmutex.Lock()
					geoData := u.geoLocationProvider.GetGeoDataByIP(ip)
					geoDBmutex.Unlock()

					mutex.Lock()
					geoReport[ip] = &Geolocation{
						GeoData:       geoData,
						SampleRequest: accessRecord.URI,
						BrowserAgent:  accessRecord.UserAgent,
						Requests:      1,
					}
					mutex.Unlock()
				}
			}
		}()
	}

	waitGroup.Wait()

	u.geoReportWriter.Save(reportPath, geoReport)
}
