package main

import (
	"log"
	"net/http"
	"socialPostsDashboard/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
