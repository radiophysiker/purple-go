package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", RandomHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
