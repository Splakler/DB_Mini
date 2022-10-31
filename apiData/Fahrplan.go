package apiData

import (
	"encoding/json"
	"log"
	"strings"
	"time"
)

type ArrivalData []struct {
	Name      string `json:"name"`
	DateTime  string `json:"dateTime"`
	Origin    string `json:"origin"`
	Track     string `json:"track"`
	DetailsId string `json:"detailsId"`
}

type DepartureData []struct {
	Name      string `json:"name"`
	DateTime  string `json:"dateTime"`
	Direction string `json:"direction"`
	Track     string `json:"track"`
	DetailsId string `json:"detailsId"`
}

type JourneyData struct {
	TrainName string
	Stops     Stop
}

type Stop []struct {
	Train    string `json:"train"`
	StopName string `json:"stopName"`
	StopId   int    `json:"stopId"`
	DepTime  string `json:"depTime"`
	ArrTime  string `json:"arrTime"`
}

func GetArrivalsFor(eva int) *ArrivalData {
	return ArrivalData{}.ReadJson(*ReqFahrplanArr(eva, getDateFromTime(time.Now())))
}

func (a ArrivalData) ReadJson(body []byte) *ArrivalData {
	res := &ArrivalData{}
	err := json.Unmarshal(body, res)
	if err != nil {
		log.Println("Error in ReadJson\n", err)
	}
	return res
}

func GetDeparturesFor(eva int) *DepartureData {
	return DepartureData{}.ReadJson(*ReqFahrplanDep(eva, getDateFromTime(time.Now())))
}

func (d DepartureData) ReadJson(body []byte) *DepartureData {
	res := &DepartureData{}
	err := json.Unmarshal(body, res)
	if err != nil {
		log.Println("Error in ReadJson\n", err)
	}
	return res
}

func GetJourneyDetailsFor(jId string) *JourneyData {
	var stops Stop
	var res JourneyData
	stops = *stops.ReadJson(*ReqFahrplanJourney(jId))
	res.Stops = stops
	if stops != nil {
		res.TrainName = stops[0].Train
	}
	res = *CleanJourneyData(res)
	return &res
}

func (j Stop) ReadJson(body []byte) *Stop {
	res := &Stop{}
	err := json.Unmarshal(body, res)
	if err != nil {
		log.Println("Error in ReadJson\n", err)
	}
	return res
}

func CleanJourneyData(j JourneyData) *JourneyData {
	for idx, element := range j.Stops {
		element.StopName = strings.Replace(element.StopName, "&#x0028;", " (", -1)
		element.StopName = strings.Replace(element.StopName, "&#x0029;", ") ", -1)
		element.StopName = strings.Replace(element.StopName, "  ", " ", -1)
		j.Stops[idx] = element
	}
	return &j
}

func getDateFromTime(t time.Time) string {
	strList := strings.Split(t.String(), " ")
	return strList[0] + "T" + strings.Split(strList[1], ".")[0]
}
