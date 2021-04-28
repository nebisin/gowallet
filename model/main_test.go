package model

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testRepository *Repository
var testSQLRepository *SQLRepository
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open("postgres", "postgres://root:password@localhost:5432/go_wallet?sslmode=disable")
	if err != nil {

		log.Fatal("cannot connect to db:", err)
	}

	testRepository = CreateRepository(testDB)
	testSQLRepository = NewSQLRepository(testDB)

	os.Exit(m.Run())
}
