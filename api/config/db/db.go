// db package houses the very lowest layer in the overall application
// architecture that communicates with the PostgreSQL database pool.
// This connection is used upstream by the model layer.
package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// Represents the database pool with the DB specific
// configurations.
type Db struct {
	Pool *sql.DB
	cfg  Config
}

// Represents the configuration parameters needed to
// open up a connection to the database.
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

// Close ends the connection to the specific database.
// It returns the error if one exists.
func (db *Db) Close() (err error) {
	if db.Pool == nil {
		return
	}
	if err = db.Pool.Close(); err != nil {
		err = errors.Wrapf(err, "Could not close postgres db")
	}
	return
}

// New creates a new connection to the existing database on the
// host system.
// It takes in a configuration to open the specific database.
// It returns the open database and an error if one exists.
func New(cfg Config) (db Db, err error) {
	if cfg.Host == "" || cfg.Port == "" || cfg.User == "" ||
		cfg.Password == "" || cfg.Database == "" {
		err = errors.New("Provide all fields for config")
		return
	}
	db.cfg = cfg

	pqDb, err := sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.User, cfg.Password, cfg.Database, cfg.Host, cfg.Port))
	if err != nil {
		err = errors.Wrapf(err, "Failed to open a connection to the database")
		return
	}
	// Ping the connection to ensure its open
	if err = pqDb.Ping(); err != nil {
		err = errors.Wrapf(err, "Unable to ping database")
		return
	}
	db.Pool = pqDb
	return
}
