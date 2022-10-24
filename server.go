package main

import (
	"DB_Mini/apiData"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

type Client struct {
	StationInfos apiData.Station
	Arrivals     apiData.ArrivalData
	Departures   apiData.DepartureData
}

func startServ() {
	port := "8080"
	fileServer := http.FileServer(http.Dir("./static"))
	fmt.Println("Downloading data... Please stand by...")
	t := time.Now()
	stationsList := apiData.FetchEverything()
	fmt.Println("Finished Fetching data. Took:", time.Now().Sub(t))
	http.Handle("/", fileServer)
	http.HandleFunc("/Search", SearchHandler(stationsList))
	http.HandleFunc("/Station", StationHandler(stationsList))
	fmt.Println("Server Started!")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

}

func SearchHandler(stationsList *apiData.StaDa) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := url.Parse(r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		params := u.Query()
		searchQuery := params.Get("q")

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tmpl := template.Must(template.ParseFiles("./static/SearchTemplate.html"))

		fmt.Fprintf(os.Stdout, "POST request successful\n")
		name := r.FormValue("name")

		var res apiData.StaDa
		if name == "" {
			res.Stations = *stationsList.SearchForName(searchQuery)
			res.Search = searchQuery
		} else {
			res.Stations = *stationsList.SearchForName(name)
			res.Search = name
		}

		err = tmpl.Execute(w, res)
		apiData.CatchError(err, "template.Execute Error!")

		fmt.Println(res.Stations)
	}
}

func StationHandler(stationsList *apiData.StaDa) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := url.Parse(r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		params := u.Query()
		searchQuery, _ := strconv.Atoi(params.Get("q"))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tmpl := template.Must(template.ParseFiles("./static/StationTemplate.html"))

		fmt.Fprintf(os.Stdout, "STATION-POST request successful\n")

		var res Client
		res.StationInfos = *stationsList.SearchFoNum(searchQuery)
		res.StationInfos.IsOpen = res.StationInfos.HasOpen()
		res.StationInfos.ImgUrl, _ = res.StationInfos.GetImageUrl()
		res.Arrivals = *apiData.GetArrivalsFor(res.StationInfos.GetMainEva())
		res.Departures = *apiData.GetDeparturesFor(res.StationInfos.GetMainEva())

		err = tmpl.Execute(w, res)
		apiData.CatchError(err, "template.Execute Error!")

		fmt.Println(res)
	}
}
