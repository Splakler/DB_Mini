package main

import (
	"DB_Mini/data"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

func startServ() {
	port := "8080"
	fileServer := http.FileServer(http.Dir("./static"))
	fmt.Println("Downloading Data... Please stand by...")
	t := time.Now()
	Data := data.FetchEverything()
	fmt.Println("Finished Fetching Data. Took:", time.Now().Sub(t))
	http.Handle("/", fileServer)
	http.HandleFunc("/Search", SearchHandler(Data))
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/Station", StationHandler(Data))
	fmt.Println("Server Started!")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	url.Parse("/form.html")
}

func SearchHandler(Data *data.StaDa) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tmpl := template.Must(template.ParseFiles("./static/SearchTemplate.html"))

		fmt.Fprintf(os.Stdout, "POST request successful\n")
		name := r.FormValue("name")
		//address := r.FormValue("address")
		//specs := r.FormValue("specs")

		var res data.StaDa
		res.Stations = *Data.SearchForName(name)
		res.Search = name

		err := tmpl.Execute(w, res)
		data.CatchError(err, "template.Execute Error!")

		fmt.Println(res.Stations)
	}
}

func StationHandler(Data *data.StaDa) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := url.Parse(r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		params := u.Query()
		searchQuery, _ := strconv.Atoi(params.Get("q"))
		//

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tmpl := template.Must(template.ParseFiles("./static/StationTemplate.html"))

		fmt.Fprintf(os.Stdout, "STATION-POST request successful\n")
		var res data.Station
		res = *Data.SearchFoNum(searchQuery)
		res.IsOpen = res.HasOpen()
		res.ImgUrl, _ = res.GetImageUrl()

		err = tmpl.Execute(w, res)
		data.CatchError(err, "template.Execute Error!")

		fmt.Println(res)
	}
}
