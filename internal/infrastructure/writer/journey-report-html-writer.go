package writer

import (
	"log"
	"os"
	"tango/internal/domain/entity"
	"text/template"

	"github.com/gobuffalo/packr/v2"
)

// JourneyReportHtmlWriter
type JourneyReportHtmlWriter struct {
}

// JourneyPlaceHtmlReport
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

// JourneyHtmlReport
type JourneyHtmlReport struct {
	ID     string
	IP     string
	Places []JourneyPlaceHtmlReport
	Roads  []JourneyRoadHtmlReport
}

// NewJourneyReportHTMLWriter inits a new html report writer
func NewJourneyReportHTMLWriter() *JourneyReportHtmlWriter {
	return &JourneyReportHtmlWriter{}
}

// Save journey report to html file
func (w *JourneyReportHtmlWriter) Save(filePath string, journeyReportData map[string]*entity.Journey) {
	templateBox := packr.New("template-box", "../../../template")
	journewReportContent, err := templateBox.FindString("journey-report.tmpl")

	if err != nil {
		log.Fatal("Error on loading journey template file", err)
	}

	journeyReportTemplate, err := template.New("journey-report").Parse(journewReportContent)

	if err != nil {
		log.Fatal("Error on loading journey template file", err)
	}

	reportWriter, err := os.Create(filePath)

	if err != nil {
		log.Println("Error on creating report file: ", err)
		return
	}

	defer reportWriter.Close()

	err = journeyReportTemplate.Execute(reportWriter, w.getJourneyReportHTMLData(journeyReportData))

	if err != nil {
		log.Println("Error on generating journey report file: ", err)
		return
	}
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
		color := w.getPlaceColor(index, journeyPlace)

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

// getPlaceColor retrieves color of journey place color
func (w *JourneyReportHtmlWriter) getPlaceColor(index int, journeyPlace *entity.JourneyPlace) string {
	color := "#3498db" // blue color

	if !journeyPlace.WasLogged {
		color = "#95a5a6" // gray color
	}

	responseCode := journeyPlace.Data.ResponseCode

	if responseCode >= 400 && responseCode < 500 {
		color = "#f1c40f" // yellow color
	}

	if responseCode >= 500 {
		color = "#e74c3c" // red color
	}

	// begining of the network will be highlighted by another color
	if index == 0 {
		color = "#1abc9c" // green color
	}

	return color
}
