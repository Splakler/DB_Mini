package data

import (
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
)

const DBApisUrl = "https://apis.deutschebahn.com/db-api-marketplace/apis/"
const StationDataUrl = DBApisUrl + "station-data/v2/stations/"
const FahrplanUrl = DBApisUrl + "fahrplan/v1/"
const StaDaUrl = DBApisUrl + "station-data/v2/stations"
const ArrivalUrl = DBApisUrl

var Cid = ""
var ApiKey = ""

func ReqFahrplan(specification, search string) *[]byte {
	req, err := http.NewRequest("GET", FahrplanUrl+specification+"/"+search, nil)
	CatchError(err, "Error in Request")

	viper.AutomaticEnv()
	req.Header.Add("DB-Client-Id", Cid)
	req.Header.Add("DB-Api-Key", ApiKey)
	req.Header.Add("accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	CatchError(err, "Request Failed or Invalid")

	body, err := ioutil.ReadAll(res.Body)
	CatchError(err, "Error in ioUtil.readall")
	return &body
}

func ReqStaDaAll() *[]byte {
	req, err := http.NewRequest("GET", StaDaUrl, nil)
	CatchError(err, "Error in Request")

	viper.AutomaticEnv()
	req.Header.Add("DB-Client-Id", Cid)
	req.Header.Add("DB-Api-Key", ApiKey)
	req.Header.Add("accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	CatchError(err, "Request Failed or Invalid")

	body, err := ioutil.ReadAll(res.Body)
	CatchError(err, "Error in ioUtil.readall")
	return &body
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
