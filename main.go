package main

import (
	"log"
	"net/http"

	"github.com/guitarpawat/go-mathapi/api"
)

func main() {
	http.HandleFunc("/", api.APIHandler)
	err := http.ListenAndServe(":8000", nil)
	log.Fatalln(err)
}
