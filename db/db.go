package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"

)


var db *sql.DB

func InitDB(infoDB string) {
	var err error
	db, err = sql.Open("postgres", infoDB)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}