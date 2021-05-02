package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nebisin/gowallet/api"
	"github.com/nebisin/gowallet/db/model"
	"github.com/nebisin/gowallet/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	fmt.Println(config)
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := model.NewStore(conn)
	server := api.NewServer(store)

	if err := server.Start(config.ServerAddress); err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
