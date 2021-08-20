package main

import (
    _ "database/sql"
    _ "fmt"
    _ "log"
    _ "github.com/lib/pq"
    _ "github.com/jmoiron/sqlx"
    restful "github.com/emicklei/go-restful"
)

func GetPlayerRankByRating(request *restful.Request, response *restful.Response){
        ts := request.PathParameter(playerRating)
        db := GetCon()
        rank := Rank{}
        db.Get(rank, "SELECT COUNT(*) from players where (mu-2*sigma) > (?-2*?)", ts)
        response.WriteEntity(rank)
}
