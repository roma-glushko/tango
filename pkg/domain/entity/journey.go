package entity

import (
	"fmt"
)

// JourneyPlace represents one URL in visitor history
// Some nodes are not directly logged to the access logs due to caching layers.
// In that case, we will replicate missing information based on referer URL and access log records
type JourneyPlace struct {
	ID        string
	WasLogged bool
	Data      *AccessLogRecord
}

// Print a debug information about journey nodes
func (n *JourneyPlace) Print() string {
	return fmt.Sprintf("%v", n.Data.URI)
}

// Journey represents visitor access history and consists of URLs they visited
// Infrastructure can include caching layers which prevents some records to be processed and logged by server engine.
// Our goal is to reproduce visitor journey as long as possible based on referer URLs
type Journey struct {
	ID     string
	IP     string
	Places []*JourneyPlace
	Roads  map[JourneyPlace][]*JourneyPlace
}

// AddPlace adds a journey Places to the journey
func (jg *Journey) AddPlace(n *JourneyPlace) *JourneyPlace {
	jg.Places = append(jg.Places, n)

	return n
}

// GetLastPlace
func (jg *Journey) GetLastPlace() *JourneyPlace {
	if len(jg.Places) == 0 {
		return nil
	}

	return jg.Places[len(jg.Places)-1]
}

// AddRoad adds an road between journey Places
func (jg *Journey) AddRoad(n1, n2 *JourneyPlace) {
	if jg.Roads == nil {
		jg.Roads = make(map[JourneyPlace][]*JourneyPlace)
	}

	jg.Roads[*n1] = append(jg.Roads[*n1], n2)
}

// Print a debug info about journey graph structure
func (jg *Journey) Print() {
	s := ""

	for i := 0; i < len(jg.Places); i++ {
		s += jg.Places[i].Print() + " -> "
		near := jg.Roads[*jg.Places[i]]

		for j := 0; j < len(near); j++ {
			s += near[j].Print() + " "
		}

		s += "\n"
	}

	fmt.Println(s)
}

// FindPlace finds a place in journey by place URI
func (jg *Journey) FindPlace(uri string) *JourneyPlace {
	for i := 0; i < len(jg.Places); i++ {
		place := jg.Places[i]

		if place.Data.URI == uri {
			return place
		}
	}

	return nil
}

// FindPlace finds a place in journey by place URI
func (jg *Journey) FindLastPlace(uri string) *JourneyPlace {
	for i := len(jg.Places) - 1; i >= 0; i-- {
		place := jg.Places[i]

		if place.Data.URI == uri {
			return place
		}
	}

	return nil
}
