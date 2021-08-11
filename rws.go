package main

import (
    _ "database/sql"
    _ "fmt"
    _ "log"
    _ "github.com/lib/pq"
    _ "github.com/jmoiron/sqlx"
    restful "github.com/emicklei/go-restful"
)

var NO_PATTERN string = ""

var playerIdIdentifier string =  "user-id"
var playerIdPattern    string = "/{" + playerIdIdentifier + "}"


func PlayerService() *restful.WebService {
	service := new(restful.WebService)
	service.
		Path("/players").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)

	service.Route(service.GET(playerIdPattern).To(FindPlayer))
	service.Route(service.POST("").To(UpdatePlayer))
	service.Route(service.PUT(playerIdPattern).To(AddPlayer))
	service.Route(service.DELETE(playerIdPattern).To(DeletePlayer))

	return service
}

var historicalTimestamp string =  "timestamp"
var historicalTimestampPattern string = "/{" + historicalTimestamp + "}"

func HistoricalDataService() *restful.WebService {
	service := new(restful.WebService)
	service.
		Path("/historical").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)

	service.Route(service.GET(historicalTimestampPattern).To(GetHistoricalData))
	service.Route(service.POST(NO_PATTERN).To(SubmitHistoricalData))
	service.Route(service.DELETE(historicalTimestampPattern).To(DeleteHistoricalData))

	return service
}
