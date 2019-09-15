package geo

import (
	"log"
	"net"

	"github.com/oschwald/geoip2-golang"
)

// MaxMindGeoProvider
type MaxMindGeoProvider struct {
	maxmindCityDatabase *geoip2.Reader
}

// NewMaxMindGeoProvider creates a new instance of MaxMindGeoProvider
func NewMaxMindGeoProvider(maxmindGeoLibPath string) *MaxMindGeoProvider {
	maxmindCityDatabase, err := geoip2.Open(maxmindGeoLibPath)

	if err != nil {
		log.Fatal(err)
	}

	return &MaxMindGeoProvider{
		maxmindCityDatabase: maxmindCityDatabase,
	}
}

// GetGeoLocationByIP provides geo location data by IP
func (p *MaxMindGeoProvider) GetGeoLocationByIP(ip string) *geoip2.City {
	parsedIP := net.ParseIP(ip)
	geoLocation, err := p.maxmindCityDatabase.City(parsedIP)

	if err != nil {
		log.Fatal(err)
	}

	return geoLocation
}

// Close
func (p *MaxMindGeoProvider) Close() error {
	return p.maxmindCityDatabase.Close()
}
