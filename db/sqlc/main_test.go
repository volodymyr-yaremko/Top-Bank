package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:virtualrelay@localhost:5432/top_bank_db?sslmode=disable"
)

var ()

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {

	var err error

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
