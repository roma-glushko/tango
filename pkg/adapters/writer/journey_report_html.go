package writer

import (
	"log"
	"os"
	"tango/pkg/entity"
	"tango/tmpl"
	"text/template"
)

// JourneyReportHTMLWriter writes journey reports to HTML files.
type JourneyReportHTMLWriter struct {
}

// JourneyPlaceHTMLReport represents a place in the journey HTML report.
type JourneyPlaceHTMLReport struct {
	ID    string
	Label string
	Color string
}

// JourneyRoadHTMLReport represents a road between places in the journey HTML report.
type JourneyRoadHTMLReport struct {
	From string
	To   string
}

// JourneyHTMLReport represents a complete journey in the HTML report.
type JourneyHTMLReport struct {
	ID     string
	IP     string
	Places []JourneyPlaceHTMLReport
	Roads  []JourneyRoadHTMLReport
}

// NewJourneyReportHTMLWriter creates a new JourneyReportHTMLWriter instance.
func NewJourneyReportHTMLWriter() *JourneyReportHTMLWriter {
	return &JourneyReportHTMLWriter{}
}

// Save writes a journey report to an HTML file.
func (w *JourneyReportHTMLWriter) Save(filePath string, journeyReportData map[string]*entity.Journey) {
	journeyReportTemplate, err := template.New("journey-report").Parse(tmpl.JourneyReportTemplate)

	if err != nil {
		log.Fatal("Error on loading journey template file", err)
	}

	reportWriter, err := os.Create(filePath)

	if err != nil {
		log.Println("Error on creating report file: ", err)
		return
	}

	defer func() { _ = reportWriter.Close() }()

	err = journeyReportTemplate.Execute(reportWriter, w.getJourneyReportHTMLData(journeyReportData))

	if err != nil {
		log.Println("Error on generating journey report file: ", err)
		return
	}
}

func (w *JourneyReportHTMLWriter) getJourneyReportHTMLData(journeyReport map[string]*entity.Journey) []JourneyHTMLReport {
	journeyHTMLReport := make([]JourneyHTMLReport, 0)

	for _, journey := range journeyReport {
		journeyHTMLReport = append(journeyHTMLReport, JourneyHTMLReport{
			ID:     journey.ID,
			IP:     journey.IP,
			Places: w.getJourneyPlaceHTMLData(journey.Places),
			Roads:  w.getJourneyRoadHTMLData(journey.Places, journey.Roads),
		})
	}

	return journeyHTMLReport
}

func (w *JourneyReportHTMLWriter) getJourneyPlaceHTMLData(journeyPlaces []*entity.JourneyPlace) []JourneyPlaceHTMLReport {
	journeyPlaceHTMLReport := make([]JourneyPlaceHTMLReport, 0)

	for index, journeyPlace := range journeyPlaces {
		color := w.getPlaceColor(index, journeyPlace)

		journeyPlaceHTMLReport = append(journeyPlaceHTMLReport, JourneyPlaceHTMLReport{
			ID:    journeyPlace.ID,
			Label: journeyPlace.Data.URI,
			Color: color,
		})
	}

	return journeyPlaceHTMLReport
}

func (w *JourneyReportHTMLWriter) getJourneyRoadHTMLData(journeyPlaces []*entity.JourneyPlace, journeyRoads map[string][]*entity.JourneyPlace) []JourneyRoadHTMLReport {
	journeyRoadHTMLReport := make([]JourneyRoadHTMLReport, 0)

	for _, journeyPlaceFrom := range journeyPlaces {
		for _, journeyPlaceTo := range journeyRoads[journeyPlaceFrom.ID] {
			journeyRoadHTMLReport = append(journeyRoadHTMLReport, JourneyRoadHTMLReport{
				From: journeyPlaceFrom.ID,
				To:   journeyPlaceTo.ID,
			})
		}
	}

	return journeyRoadHTMLReport
}

func (w *JourneyReportHTMLWriter) getPlaceColor(index int, journeyPlace *entity.JourneyPlace) string {
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

	// beginning of the network will be highlighted by another color
	if index == 0 {
		color = "#1abc9c" // green color
	}

	return color
}
