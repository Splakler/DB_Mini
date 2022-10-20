package data

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type Ort struct {
	IsEmpty   bool
	Search    string
	Locations []Location
}
type Location struct {
	Name string  `json:"name"`
	Lon  float64 `json:"lon"`
	Lat  float64 `json:"lat"`
	Id   int     `json:"id"`
}

func (o Ort) ReadJson(body *[]byte) *Ort {
	res := &[]Location{}
	err := json.Unmarshal(*body, res)
	CatchError(err, "Error in ReadJson")
	erg := Ort{false, "", *res}
	if erg.Locations == nil {
		erg.IsEmpty = true
	}
	return &erg
}

func (o Ort) Test() {
	return
}

func SearchFor(search string) *[]byte {
	var body *[]byte
	_, err := strconv.Atoi(search)
	if err != nil {
		body = ReqFahrplan("location", search)
	} else {
		year, month, day := time.Now().Date()
		body = ReqFahrplan("arrivalBoard", search+"?date="+strconv.Itoa(year)+"-"+string(month)+"-"+strconv.Itoa(day))
	}
	return body
}

func (o Ort) CleanData(search string) *Ort {
	if o.IsEmpty {
		return &o
	}

	var temp Ort
	temp.Search = search
	for _, Elem := range o.Locations {
		if strings.Contains(strings.ToUpper(Elem.Name), strings.ToUpper(search)) {
			temp.Locations = append(temp.Locations, Elem)
		}
	}
	if temp.Locations == nil {
		temp.IsEmpty = true
	}
	return &temp
}
