package report

import (
	"log"
	"net"
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

// GeoReportWriter is an interface for saving geo location reports
type GeoReportWriter interface {
	Save(reportPath string, geolocationReport map[string]*Geolocation)
}

// GeoReportUsecase knows how to generate geo reports
type GeoReportUsecase struct {
	geoReportWriter GeoReportWriter
}

//
func NewGeoReportUsecase(geoReportWriter GeoReportWriter) *GeoReportUsecase {
	return &GeoReportUsecase{
		geoReportWriter: geoReportWriter,
	}
}

// GenerateReport processes access logs and collect geo reports
func (u *GeoReportUsecase) GenerateReport(reportPath string, accessRecords []entity.AccessLogRecord) {
	var geoReport = make(map[string]*Geolocation)

	db, err := geoip2.Open("assets/GeoLite2-City.mmdb") // todo: move working with GeoLite to infrastructure layer
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for _, accessRecord := range accessRecords {
		for _, ip := range accessRecord.IP {

			if _, ok := geoReport[ip]; ok {
				geoReport[ip].Requests++
				continue
			}

			ipRecord := net.ParseIP(ip)
			geoLocation, err := db.City(ipRecord)

			if err != nil {
				log.Fatal(err)
			}

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
