package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/siddheshRajendraNimbalkar/intern/api"
	db "github.com/siddheshRajendraNimbalkar/intern/db/sqlc"
	"github.com/siddheshRajendraNimbalkar/intern/util"
)

func main() {

	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("[Config]::error in env", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("[ERROR IN Main_Test]::While connecting db", err)
	}

	store := db.NewStore(conn)

	server := api.NewServer(config, store)

	err = server.Start(config.Addr)

	if err != nil {
		log.Fatal("[ERROR OCCURE WHILE CONNECTING THE PORT]::", err.Error())
	}
}
