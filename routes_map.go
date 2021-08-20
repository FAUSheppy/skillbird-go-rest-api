package main

import (
    _ "database/sql"
    _ "fmt"
    _ "log"
    _ "github.com/lib/pq"
    _ "github.com/jmoiron/sqlx"
    restful "github.com/emicklei/go-restful"
)

func GetMap(request *restful.Request, response *restful.Response){
        id := request.PathParameter(roundTimestampIdentifier)
        db := GetCon()
        round := Round{}
        db.Get(round, "SELECT * from maps where name=$1", id)
        response.WriteEntity(round)
}
