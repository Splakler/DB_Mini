package data

import (
	"encoding/json"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type StaDa struct {
	Total    int       `json:"total"`
	Stations []Station `json:"result"`
}

type Station struct {
	IsOpen         bool
	ImgUrl         *url.URL
	Num            int    `json:"number"`
	Name           string `json:"name"`
	Category       int    `json:"category"`
	MailingAddress struct {
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
		Street  string `json:"street"`
	} `json:"mailingAddress"`
	EvaNumbers    []evaNumber `json:"evaNumbers"`
	HasDBLounge   bool        `json:"hasDBLounge"`
	DBInformation struct {
		Availability struct {
			Monday    day `json:"monday"`
			Tuesday   day `json:"tuesday"`
			Wednesday day `json:"wednesday"`
			Thursday  day `json:"thursday"`
			Friday    day `json:"friday"`
			Saturday  day `json:"saturday"`
			Sunday    day `json:"sunday"`
			Holiday   day `json:"holiday"`
		} `json:"availability"`
	} `json:"DBInformation"`
}

type evaNumber struct {
	IsMain                bool `json:"isMain"`
	Eva                   int  `json:"number"`
	GeographicCoordinates struct {
		Coordinates [2]float64 `json:"coordinates"`
	} `json:"geographicCoordinates"`
}

type day struct {
	FromTime string `json:"fromTime"`
	ToTime   string `json:"toTime"`
}

func (s StaDa) NewStation() *StaDa {
	return &StaDa{}
}

func (s StaDa) FetchEverything() *StaDa {
	var res StaDa
	return res.ReadJson(*ReqStaDaAll())
}

func (s StaDa) ReadJson(body []byte) *StaDa {
	res := &StaDa{}
	err := json.Unmarshal(body, res)
	if err != nil {
		log.Println("Error in ReadJson\n", err)
	}
	return res
}

func (s StaDa) SearchForName(search string) *[]Station {
	var res []Station = nil
	for _, elem := range s.Stations {
		if strings.Contains(elem.Name, search) {
			res = append(res, elem)
		}
	}
	return &res
}

func (s StaDa) SearchFoNum(search int) *Station {
	for _, elem := range s.Stations {
		if elem.Num == search {
			return &elem
		}
	}
	return nil
}

func (stat Station) HasOpen() bool {
	if stat.DBInformation.Availability.Monday.ToTime != "" {
		currHour, currMinute, _ := time.Now().Clock()
		openTime := strings.Split(stat.DBInformation.Availability.Monday.FromTime, ":")
		closeTime := strings.Split(stat.DBInformation.Availability.Monday.ToTime, ":")
		opHour, _ := strconv.Atoi(openTime[0])
		opMinute, _ := strconv.Atoi(openTime[1])
		clHour, _ := strconv.Atoi(closeTime[0])
		clMinute, _ := strconv.Atoi(closeTime[1])
		if currHour == opHour && currMinute >= opMinute {
			return true
		} else if currHour > opHour && currHour < clHour {
			return true
		} else if currHour == clHour && currMinute < clMinute {
			return true
		} else {
			return false
		}
	}
	return false
}

func (s Station) GetImageUrl() (*url.URL, error) {
	return url.Parse("https://api.railway-stations.org/photos/de/" + strconv.Itoa(s.Num) + "_1.jpg")
}
