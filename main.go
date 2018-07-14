package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/guitarpawat/go-mathapi/api"
)

var host = flag.String("host", "localhost", "Specify running host of API")
var port = flag.Int("port", 8000, "Specify running host of API")

func main() {
	flag.Parse()
	http.HandleFunc("/", api.APIHandler)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", *host, *port), nil)
	log.Fatalln(err)
}
