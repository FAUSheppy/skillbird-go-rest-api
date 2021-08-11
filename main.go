package main

import (
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
)

func main() {
	restful.Add(PlayerService())
	restful.Add(HistoricalDataService())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
