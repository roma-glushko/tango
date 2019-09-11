package report

import (
	"tango/internal/domain/entity"

	"github.com/oschwald/geoip2-golang"
)

// Geolocation report information
type Geolocation struct {
	Country       string
	City          string
	Continent     string
	SampleRequest string
	BrowserAgent  string
	Requests      uint64
}

// GeoLocationProvider provides ability to find geolocation data by IP
type GeoLocationProvider interface {
	GetGeoLocationByIP(ip string) *geoip2.City // keep geoip2 references in the usecase layer, don't plan to support anything else
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

			geoLocation := u.geoLocationProvider.GetGeoLocationByIP(ip)

			geoReport[ip] = &Geolocation{
				Country:       geoLocation.Country.Names["en"],
				City:          geoLocation.City.Names["en"],
				Continent:     geoLocation.Continent.Names["en"],
				SampleRequest: accessRecord.URI,
				BrowserAgent:  accessRecord.UserAgent,
				Requests:      1,
			}
		}
	}

	u.geoReportWriter.Save(reportPath, geoReport)
}
