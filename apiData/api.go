package apiData

import (
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const DBApisUrl = "https://apis.deutschebahn.com/db-api-marketplace/apis/"
const FahrplanUrl = DBApisUrl + "fahrplan/v1/"
const StaDaUrl = DBApisUrl + "station-data/v2/stations"
const ArrivalUrl = FahrplanUrl + "arrivalBoard/"
const DepartureUrl = FahrplanUrl + "departureBoard/"
const JourneyUrl = FahrplanUrl + "journeyDetails/"

var Cid = ""
var ApiKey = ""

func ReqStaDaAll() *[]byte {
	body, err := ioutil.ReadAll(GetUrl(StaDaUrl).Body)
	CatchError(err, "Error in ioUtil.readall")
	return &body
}

func ReqFahrplanArr(id int, date string) *[]byte {
	body, err := ioutil.ReadAll(GetUrl(ArrivalUrl + strconv.Itoa(id) + "?date=" + date).Body)
	CatchError(err, "Error in ioUtil.readall")
	return &body
}

func ReqFahrplanDep(id int, date string) *[]byte {
	body, err := ioutil.ReadAll(GetUrl(DepartureUrl + strconv.Itoa(id) + "?date=" + date).Body)
	CatchError(err, "Error in ioUtil.readall")
	return &body
}

func ReqFahrplanJourney(jId string) *[]byte {
	body, err := ioutil.ReadAll(GetUrl(JourneyUrl + jId).Body)
	CatchError(err, "Error in ioUtil.readall")
	return &body
}

func GetUrl(url string) *http.Response {
	req, err := http.NewRequest("GET", url, nil)
	CatchError(err, "Error in Request")

	viper.AutomaticEnv()
	req.Header.Add("DB-Client-Id", Cid)
	req.Header.Add("DB-Api-Key", ApiKey)
	req.Header.Add("accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	CatchError(err, "Request Failed or Invalid")
	return res
}

func CatchError(err error, msg string) {
	if err != nil {
		log.Fatal(msg+"\n", err)
	}
}

func ReqKeys() {
	viper.AutomaticEnv()
	Cid = viper.GetString("Cid")
	ApiKey = viper.GetString("API_KEY")
}

//func ReadJson(obj interface{}, body []byte) interface{} {
//	res := &obj
//	err := json.Unmarshal(body, res)
//	if err != nil {
//		log.Println("Error in ReadJson\n", err)
//	}
//	return &res
//}
