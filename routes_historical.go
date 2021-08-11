package main

import (
    _ "database/sql"
    _ "fmt"
    _ "log"
    _ "github.com/lib/pq"
    _ "github.com/jmoiron/sqlx"
    restful "github.com/emicklei/go-restful"
)

func GetHistoricalData(request *restful.Request, response *restful.Response){
        ts := request.PathParameter(historicalTimestamp)
        db := GetCon()
        entry := HistoricalEntry{}
        db.Get(entry, "SELECT * from playersHistoricalData where timestamp=$1", ts)
        response.WriteEntity(entry)
}

func SubmitHistoricalData(request *restful.Request, response *restful.Response){
        entry := new(HistoricalEntry)
        request.ReadEntity(&entry)
        db := GetCon()
        db.NamedExec(`INSERT INTO playersHistoricalData (id, timestamp, mu, sigma)
        VALUES (:id, :name, :timestamp, :mu, :sigma)`, entry)
        response.WriteEntity(entry)
}

func DeleteHistoricalData(request *restful.Request, response *restful.Response){
        ts := request.PathParameter(historicalTimestamp)
        db := GetCon()
        entry := HistoricalEntry{}
        db.Get(entry, "DELETE from playersHistoricalData where timestamp=$1", ts)
}
