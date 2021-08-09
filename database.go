package main

import (
    _ "database/sql"
    _ "fmt"
    "log"

    _ "github.com/lib/pq"
    "github.com/jmoiron/sqlx"
)

var schema = `
CREATE TABLE players(
	id INTEGER,
	name TEXT,
	lastGame TEXT,
	wins INTEGER,
	mu INTEGER,
	sigma INTEGER,
	games INTEGER);

create TABLE playerHistoricalData (id TEXT, timestamp TEXT, mu REAL, sima REAL);
create TABLE live (id text, time INTEGER, duration INTEGER, players TEXT);
create TABLE rounds (
    timestamp TEXT PRIMARY KEY,
    winners BLOB,
    losers BLOB,
    winnerSide INTEGER,
    map TEXT,
    duration REAL,
    prediction INTEGER,
    confidence REAL
);`

type Player struct {
    id       int    `db:"id"`
    name     string `db:"name"`
    lastGame string `db:"lastGame"`
    wins     int    `db:"wins"`
    mu       int    `db:"mu"`
    sigma    int    `db:"sigma"`
    games    int    `db:"games"`
}

type PlayerHistoricalDataEntry struct {
    id          string  `db:"id"`
    timestamp   string  `db:"timestamp"`
    mu          float64 `db:"mu"`
    sigma       float64 `db:"sigma"`
}

func Init(){
    db := GetCon()
    db.MustExec(schema)
}

func GetCon() *sqlx.DB{
    db, err := sqlx.Connect("postgres", "user=foo dbname=bar sslmode=disable")
    if err != nil {
        log.Fatalln(err)
    }

    return db;
}

func CloseAndCommit(*sqlx.DB){
}
