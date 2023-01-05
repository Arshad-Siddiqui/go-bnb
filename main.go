package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		index(w, r)
	default:
		http.NotFound(w, r)
	}
}

var staticPath = "./static/"

func index(w http.ResponseWriter, r *http.Request) {
	fileName := "index.html"
	t, err := template.ParseFiles(staticPath + fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.ExecuteTemplate(w, fileName, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
