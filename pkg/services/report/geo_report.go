package report

import (
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
func (u *GeoReportService) GenerateReport(reportPath string, accessRecords []entity.AccessLogRecord) {
	var geoReport = make(map[string]*Geolocation)

	defer u.geoLocationProvider.Close()

	for _, accessRecord := range accessRecords {
		for _, ip := range accessRecord.IP {

			if _, ok := geoReport[ip]; ok {
				geoReport[ip].Requests++
				continue
			}

			geoData := u.geoLocationProvider.GetGeoDataByIP(ip)

			geoReport[ip] = &Geolocation{
				GeoData:       geoData,
				SampleRequest: accessRecord.URI,
				BrowserAgent:  accessRecord.UserAgent,
				Requests:      1,
			}
		}
	}

	u.geoReportWriter.Save(reportPath, geoReport)
}
