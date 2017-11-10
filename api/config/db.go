package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Db struct {
	Pool *sql.DB
	cfg  Config
}

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

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
		err = errors.Wrapf(err, "Couldn't open connection to postgres database")
		return
	}
	// Test
	if err = pqDb.Ping(); err != nil {
		err = errors.Wrapf(err, "Unable to ping database")
		return
	}
	db.Pool = pqDb
	return
}

func (db *Db) Close() (err error) {
	if db.Pool == nil {
		return
	}
	if err = db.Pool.Close(); err != nil {
		err = errors.Wrapf(err, "Could not close postgres db")
	}
	return
}
