package saver

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"tango/internal/usecase"
)

// Save GeoLocation Report to CSV file
func Save(filePath string, geolocationReport map[string]*usecase.Geolocation) {
	file, err := os.Create(filePath)

	if err != nil {
		log.Fatal("Error on writing geolocation report: ", err)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Header
	writer.Write([]string{
		"IP",
		"Country",
		"City",
		"Count of Requests",
	})

	// Body
	for ip, geoLocation := range geolocationReport {
		err := writer.Write([]string{
			ip,
			geoLocation.Country,
			geoLocation.City,
			strconv.FormatUint(geoLocation.Requests, 10),
		})

		if err != nil {
			log.Fatal("Error on writing geolocation report: ", err)
		}
	}
}
