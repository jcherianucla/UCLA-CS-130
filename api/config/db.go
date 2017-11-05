package config

import (
	"database/sql"
	"github.com/jcherianucla/UCLA-CS-130/api/utilities"
	_ "github.com/lib/pq"
)

type DB struct{}

func (*DB) Init() {
	// Start a postgres db with provided information
	psqlDb, err := sql.Open("postgres", utilities.DB_INFO)
	defer psqlDb.Close()
	utilities.CheckError(err)
}
