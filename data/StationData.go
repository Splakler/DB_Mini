package data

import (
	"encoding/json"
	"log"
)

type Station struct {
	Result []Results `json:"result"`
}

type Results struct {
	Name           string `json:"name"`
	Category       int    `json:"category"`
	MailingAddress struct {
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
		Street  string `json:"street"`
	} `json:"mailingAddress"`
}

func (s Station) NewStation() *Station {
	return &Station{}
}

func (s Station) ReadJson(body []byte) *Station {
	res := &Station{}
	err := json.Unmarshal(body, res)
	if err != nil {
		log.Println("Error in ReadJson\n", err)
	}
	return res
}
