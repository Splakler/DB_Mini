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
	Stops     []struct {
		Train    string `json:"train"`
		StopName string `json:"stop_name"`
		StopId   int    `json:"stop_id"`
		DepTime  string `json:"dep_time"`
		ArrTime  string `json:"arr_time"`
	}
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
