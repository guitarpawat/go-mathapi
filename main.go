package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/guitarpawat/go-mathapi/api"
)

var host = flag.String("host", "localhost", "Specify running host of API")
var port = flag.Int("port", 0, "Specify running host of API")

func main() {
	flag.Parse()
	var err error
	if *port == 0 {
		*port, err = strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			panic(err)
		}
	}
	http.HandleFunc("/", api.APIHandler)
	err = http.ListenAndServe(fmt.Sprintf("%s:%d", *host, *port), nil)
	log.Fatalln(err)
}
