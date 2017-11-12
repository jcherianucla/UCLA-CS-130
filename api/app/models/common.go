package models

import (
	"github.com/jcherianucla/UCLA-CS-130/api/config/db"
	"github.com/jcherianucla/UCLA-CS-130/api/utilities"
	"sync"
)

// Layer passed back to Controllers
type layer struct {
	User *UserTable
}

var instance *layer
var once sync.Once

// Singleton access to the model layer
func LayerInstance() *layer {
	once.Do(func() {
		// Create DB only once
		db, err := db.New(db.Config{
			utilities.DB_HOST,
			utilities.DB_PORT,
			utilities.DB_USER,
			utilities.DB_PASSWORD,
			utilities.DB_NAME,
		})
		utilities.CheckError(err)
		// Create user table only once
		ut, err := NewUserTable(&db)
		utilities.CheckError(err)
		instance = &layer{User: &ut}
	})
	return instance
}
