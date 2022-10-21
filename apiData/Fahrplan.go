package apiData

import (
	"encoding/json"
	"log"
	"strings"
	"time"
)

type ArrivalData struct {
	Arrivals []struct {
		Name      string `json:"name"`
		DateTime  string `json:"date_time"`
		Origin    string `json:"origin"`
		Track     string `json:"track"`
		DetailsId string `json:"details_id"`
	}
}

type DepartureData struct {
	Departures []struct {
		Name      string `json:"name"`
		DateTime  string `json:"date_time"`
		Direction string `json:"direction"`
		Track     string `json:"track"`
		DetailsId string `json:"detailsId"`
	}
}

type JourneyData struct {
	TrainName string
	Stops     []struct {
		Train    string `json:"train"`
		StopName string `json:"stop_name"`
		StopId   int    `json:"stop_id"`
		DepTime  string `json:"dep_time"`
		ArrTime  string `json:"arr_time"`
	}
}

func (a ArrivalData) GetArrivalsFor(eva int) {
	a = *a.ReadJson(*ReqFahrplanArr(eva, getDateFromTime(time.Now())))
}

func (a ArrivalData) ReadJson(body []byte) *ArrivalData {
	res := &ArrivalData{}
	err := json.Unmarshal(body, res)
	if err != nil {
		log.Println("Error in ReadJson\n", err)
	}
	return res
}

func (d DepartureData) GetDeparturesFor(eva int) {
	d = *d.ReadJson(*ReqFahrplanDep(eva, getDateFromTime(time.Now())))
}

func (d DepartureData) ReadJson(body []byte) *DepartureData {
	res := &DepartureData{}
	err := json.Unmarshal(body, res)
	if err != nil {
		log.Println("Error in ReadJson\n", err)
	}
	return res
}

func (j JourneyData) GetJourneyDetailsFor(jId string) {
	j = *j.ReadJson(*ReqFahrplanJourney(jId))
}

func (j JourneyData) ReadJson(body []byte) *JourneyData {
	res := &JourneyData{}
	err := json.Unmarshal(body, res)
	if err != nil {
		log.Println("Error in ReadJson\n", err)
	}
	return res
}

func getDateFromTime(t time.Time) string {
	strList := strings.Split(t.String(), " ")
	return strList[0] + "T" + strings.Split(strList[1], ".")[0]
}
