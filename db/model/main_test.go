package model

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/nebisin/gowallet/util"
	"log"
	"os"
	"testing"
)

var testRepository *Repository
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testRepository = CreateRepository(testDB)

	os.Exit(m.Run())
}
