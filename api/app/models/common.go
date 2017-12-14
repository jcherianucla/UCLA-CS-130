package models

import (
	"github.com/jcherianucla/UCLA-CS-130/api/config/db"
	"github.com/jcherianucla/UCLA-CS-130/api/utilities"
	"sync"
)

// Represents the layer for the model by exposing the
// different models' tables.
type layer struct {
	User       *UserTable
	Class      *ClassTable
	Assignment *AssignmentTable
	Submission *SubmissionTable
	Enrolled   *EnrolledTable
}

// Singleton reference to the model layer.
var instance *layer

// Lock for running only once.
var once sync.Once

// LayerInstance gets the static singelton reference
// using double check synchronization.
// It returns the reference to the layer.
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
		// Create all the tables
		ut, err := NewUserTable(&db)
		utilities.CheckError(err)
		ct, err := NewClassTable(&db)
		utilities.CheckError(err)
		at, err := NewAssignmentTable(&db)
		utilities.CheckError(err)
		st, err := NewSubmissionTable(&db)
		utilities.CheckError(err)
		et, err := NewEnrolledTable(&db)
		utilities.CheckError(err)
		// Create the layer only once
		instance = &layer{User: &ut, Class: &ct, Assignment: &at, Submission: &st, Enrolled: &et}
	})
	return instance
}
