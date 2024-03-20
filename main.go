package main

import (
	"database/sql"
	"log"
	"moneytracker/api"
	db "moneytracker/db/sqlc"
	"moneytracker/db/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
	defer conn.Close()

	queries := db.New(conn)
	server, err := api.NewServer(config, queries)
	if err != nil {
		log.Fatal("Cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
