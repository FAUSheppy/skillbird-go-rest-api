package main

import (
    _ "database/sql"
    _ "fmt"
    _ "log"
    _ "github.com/lib/pq"
    _ "github.com/jmoiron/sqlx"
    restful "github.com/emicklei/go-restful"
)

func GetRound(request *restful.Request, response *restful.Response){
        id := request.PathParameter(roundTimestampIdentifier)
        db := GetCon()
        round := Round{}
        db.Get(round, "SELECT * from rounds where id=$1", id)
        response.WriteEntity(round)
}

func SubmitRound(request *restful.Request, response *restful.Response){
        round := new(Round)
		request.ReadEntity(&round)
        db := GetCon()
        db.NamedExec(`INSERT INTO players (id, name, lastgame, wins, mu, sigma, games)
        VALUES (:timestamp, :winners, :losers, :winnerSide, :map, :duration, :prediction,
                :confidence)`, round)
}

func DeleteRound(request *restful.Request, response *restful.Response){
        ts := request.PathParameter(roundTimestampIdentifier)
        db := GetCon()
        round := Round{}
        db.Get(round, "DELETE from playersHistoricalData where timestamp=$1", ts)
}
