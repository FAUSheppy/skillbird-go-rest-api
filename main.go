package main

import (
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
)

func main() {
	restful.Add(PlayerService())
	restful.Add(HistoricalDataService())
	restful.Add(RoundService())
	restful.Add(MapService())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
