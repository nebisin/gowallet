package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/nebisin/gowallet/api"
	"github.com/nebisin/gowallet/db/model"
	"log"
)

func main() {
	conn, err := sql.Open("postgres", "postgres://root:password@localhost:5432/go_wallet?sslmode=disable")
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := model.NewStore(conn)
	server := api.NewServer(store)

	if err := server.Start("0.0.0.0:8080"); err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
