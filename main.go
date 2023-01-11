package main

import (
	"log"
	"net/http"
)

func main() {
	topMux := http.NewServeMux()
	topMux.Handle("/", http.FileServer(http.Dir("static/home")))
	log.Fatal(http.ListenAndServe(":8080", topMux))
}
