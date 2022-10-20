package main

import (
	"DB_Mini/data"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
)

func runServer() {
	port := "8080"
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/Search", SearchHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/Station", StationHandler)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	url.Parse("/form.html")
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl := template.Must(template.ParseFiles("./static/template.html"))

	fmt.Fprintf(os.Stdout, "POST request successful\n")
	name := r.FormValue("name")
	//address := r.FormValue("address")
	//specs := r.FormValue("specs")

	res := data.Ort{}.ReadJson(data.SearchFor(name))
	res = res.CleanData(name)

	err := tmpl.Execute(w, res)
	data.CatchError(err, "template.Execute Error!")

	fmt.Println(res.IsEmpty)
	fmt.Println(res.Locations)
}

func StationHandlerfunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := url.Parse(r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		params := u.Query()
		searchQuery := params.Get("q")
	}
}
