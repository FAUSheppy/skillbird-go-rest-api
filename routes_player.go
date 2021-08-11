package main

import (
    _ "database/sql"
    _ "fmt"
    _ "log"
    _ "github.com/lib/pq"
    _ "github.com/jmoiron/sqlx"
    restful "github.com/emicklei/go-restful"
)

func FindPlayer(request *restful.Request, response *restful.Response){
        id := request.PathParameter(playerIdIdentifier)
        db := GetCon()
        player := Player{}
        db.Get(player, "SELECT * from players where id=$1", id)
        response.WriteEntity(player)
}

func UpdatePlayer(request *restful.Request, response *restful.Response){
        DeletePlayer(request, response)
        AddPlayer(request, response)
}

func AddPlayer(request *restful.Request, response *restful.Response){
        player := new(Player)
		request.ReadEntity(&player)
        db := GetCon()
        db.NamedExec(`INSERT INTO players (id, name, lastgame, wins, mu, sigma, games)
        VALUES (:id, :name, :lastGame, :wins, :mu, :sigma, :games)`, player)
        response.WriteEntity(player)
}

func DeletePlayer(request *restful.Request, response *restful.Response){
        id := request.PathParameter(playerIdIdentifier)
        db := GetCon()
        player := Player{}
        db.MustExec("DELETE from players where id=$1", id)
        response.WriteEntity(player)
}
