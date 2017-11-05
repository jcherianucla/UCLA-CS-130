package utilities

import (
	"fmt"
	"os"
)

const (
	DB_USER     = "gradeportal"
	DB_PASSWORD = "cs130gp"
	// TODO: Make this dependent on the application environment
	DB_NAME = "gp_development"
)

var DB_INFO = func() string {

	dbInfo := os.Getenv("DB_INFO")

	if dbInfo == "" {
		dbInfo = fmt.Sprintf(
			"user=%s password=%s dbname=%s sslmode=disable",
			DB_USER,
			DB_PASSWORD,
			DB_NAME,
		)
		os.Setenv("DB_INFO", dbInfo)
	}
	return dbInfo
}()
