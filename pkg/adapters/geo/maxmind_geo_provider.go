package geo

import (
	"log"
	"net"
	"sync"
	"tango/pkg/services/report"

	"github.com/oschwald/geoip2-golang"
)

// MaxMindGeoProvider
type MaxMindGeoProvider struct {
	maxmindCityDatabase *geoip2.Reader
	mu                  sync.Mutex
}

// NewMaxMindGeoProvider creates a new instance of MaxMindGeoProvider
func NewMaxMindGeoProvider(maxmindGeoLibPath string) *MaxMindGeoProvider {
	maxmindCityDatabase, err := geoip2.Open(maxmindGeoLibPath)

	if err != nil {
		log.Fatal(err)
	}

	return &MaxMindGeoProvider{
		maxmindCityDatabase: maxmindCityDatabase,
		mu:                  sync.Mutex{},
	}
}

// GetGeoDataByIP provides geo location data by IP
func (p *MaxMindGeoProvider) GetGeoDataByIP(IP string) *report.GeoData {
	p.mu.Lock()
	defer p.mu.Unlock()

	parsedIP := net.ParseIP(IP)
	geoLocation, err := p.maxmindCityDatabase.City(parsedIP)

	if err != nil {
		// todo: would be nice to log this errors
		return &report.GeoData{
			Country:   "N/A",
			City:      "N/A",
			Continent: "N/A",
		}
	}

	return &report.GeoData{
		Country:   geoLocation.Country.Names["en"],
		City:      geoLocation.City.Names["en"],
		Continent: geoLocation.Continent.Names["en"],
	}
}

// Close
func (p *MaxMindGeoProvider) Close() error {
	return p.maxmindCityDatabase.Close()
}
