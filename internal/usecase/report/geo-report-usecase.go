package report

import (
	"tango/internal/domain/entity"
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
	GetGeoDataByIP(ip string) *GeoData // keep geoip2 references in the usecase layer, don't plan to support anything else
	Close() error
}

// GeoReportWriter is an interface for saving geo location reports
type GeoReportWriter interface {
	Save(reportPath string, geolocationReport map[string]*Geolocation)
}

// GeoReportUsecase knows how to generate geo reports
type GeoReportUsecase struct {
	geoLocationProvider GeoLocationProvider
	geoReportWriter     GeoReportWriter
}

// NewGeoReportUsecase
func NewGeoReportUsecase(geoLocationProvider GeoLocationProvider, geoReportWriter GeoReportWriter) *GeoReportUsecase {
	return &GeoReportUsecase{
		geoLocationProvider: geoLocationProvider,
		geoReportWriter:     geoReportWriter,
	}
}

// GenerateReport processes access logs and collect geo reports
func (u *GeoReportUsecase) GenerateReport(reportPath string, accessRecords []entity.AccessLogRecord) {
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
