package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"time"
)

const staticURL string = "/static/"
const staticRoot string = "static/"

type data struct {
	LiveData string
}

func main() {
	http.HandleFunc("/livedata/", liveDataHandler)
	http.HandleFunc(staticURL, StaticHandler)
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func liveDataHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("livedata.html")
	if err != nil {
		log.Fatal("liveDataHandler: ", err)
	}
	data := &data{LiveData: "Sample Data"}
	err = t.Execute(w, data)
	if err != nil {
		log.Fatal("liveDataHandler: ", err)
	}
}

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	staticFile := r.URL.Path[len(staticURL):]
	if len(staticFile) != 0 {
		f, err := http.Dir(staticRoot).Open(staticFile)
		if err == nil {
			content := io.ReadSeeker(f)
			http.ServeContent(w, r, staticFile, time.Now(), content)
			return
		}
	}
	http.NotFound(w, r)
}
