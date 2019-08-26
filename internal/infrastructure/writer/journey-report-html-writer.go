package writer

import (
	"log"
	"os"
	"tango/internal/domain/entity"
	"text/template"
)

//
type JourneyReportHtmlWriter struct {
}

//
type JourneyPlaceHtmlReport struct {
	ID    string
	Label string
	Color string
}

// JourneyRoadHtmlReport
type JourneyRoadHtmlReport struct {
	From string
	To   string
}

type JourneyHtmlReport struct {
	ID     string
	IP     string
	Places []JourneyPlaceHtmlReport
	Roads  []JourneyRoadHtmlReport
}

// NewJourneyReportHtmlWriter inits a new html report writer
func NewJourneyReportHtmlWriter() *JourneyReportHtmlWriter {
	return &JourneyReportHtmlWriter{}
}

// Save journey report to html file
func (w *JourneyReportHtmlWriter) Save(filePath string, journeyReport map[string]*entity.Journey) {
	reportTemplate, err := template.ParseFiles("template/journey-report.tmpl")

	if err != nil {
		log.Fatal("Error on reading journey report template: ", err)
	}

	reportWriter, err := os.Create(filePath)

	if err != nil {
		log.Println("Error on opening report file: ", err)
		return
	}

	err = reportTemplate.Execute(reportWriter, w.getJourneyReportHTMLData(journeyReport))

	if err != nil {
		log.Println("Error on generating journey report file: ", err)
		return
	}

	reportWriter.Close()
}

// getJourneyReportHTMLData
func (w *JourneyReportHtmlWriter) getJourneyReportHTMLData(journeyReport map[string]*entity.Journey) []JourneyHtmlReport {
	journeyHtmlReport := make([]JourneyHtmlReport, 0)

	for _, journey := range journeyReport {
		journeyHtmlReport = append(journeyHtmlReport, JourneyHtmlReport{
			ID:     journey.ID,
			IP:     journey.IP,
			Places: w.getJourneyPlaceHTMLData(journey.Places),
			Roads:  w.getJourneyRoadHTMLData(journey.Places, journey.Roads),
		})
	}

	return journeyHtmlReport
}

// getJourneyPlaceHTMLData
func (w *JourneyReportHtmlWriter) getJourneyPlaceHTMLData(journeyPlaces []*entity.JourneyPlace) []JourneyPlaceHtmlReport {
	journeyPlaceHtmlReport := make([]JourneyPlaceHtmlReport, 0)

	for index, journeyPlace := range journeyPlaces {
		color := "#3498db"

		if !journeyPlace.WasLogged {
			color = "#95a5a6"
		}

		// begining of the network will be highlighted by another color
		if index == 0 {
			color = "#1abc9c"
		}

		journeyPlaceHtmlReport = append(journeyPlaceHtmlReport, JourneyPlaceHtmlReport{
			ID:    journeyPlace.ID,
			Label: journeyPlace.Data.URI,
			Color: color,
		})
	}

	return journeyPlaceHtmlReport
}

// getJourneyRoadHTMLData
func (w *JourneyReportHtmlWriter) getJourneyRoadHTMLData(journeyPlaces []*entity.JourneyPlace, journeyRoads map[entity.JourneyPlace][]*entity.JourneyPlace) []JourneyRoadHtmlReport {
	journeyRoadHtmlReport := make([]JourneyRoadHtmlReport, 0)

	for _, journeyPlaceFrom := range journeyPlaces {
		for _, journeyPlaceTo := range journeyRoads[*journeyPlaceFrom] {
			journeyRoadHtmlReport = append(journeyRoadHtmlReport, JourneyRoadHtmlReport{
				From: journeyPlaceFrom.ID,
				To:   journeyPlaceTo.ID,
			})
		}
	}

	return journeyRoadHtmlReport
}
