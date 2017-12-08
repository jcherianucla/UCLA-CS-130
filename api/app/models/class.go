// The models package houses the Model layer within the MVC architecture design.
// This is the lowest layer in the architecture, which directly communicates with
// the database layer. The models represent an abstraction to the DB object relations.
package models

import (
	"encoding/json"
	"fmt"
	"github.com/jcherianucla/UCLA-CS-130/api/config/db"
	"github.com/jcherianucla/UCLA-CS-130/api/utilities"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"time"
)

// The table name as set in the Postgres DB creation.
const (
	CLASS_TABLE = "classes"
)

// Represents the connection to the DB instance.
type ClassTable struct {
	connection *db.Db
}

// Represents a Class row in the Class table within the DB.
// Validator and json tags are used for convenient serialization and
// deserialization.
type Class struct {
	Id           int64     `valid:"-" json:"-"`
	Name         string    `valid:"required" json:"name"`
	Description  string    `valid:"-" json:"description"`
	Quarter      string    `valid:"required" json:"quarter"`
	Year         string    `valid:"required" json:"year"`
	Creator_id   int64     `valid:"-" json:"creator_id"`
	Time_created time.Time `valid:"-" json:"-"`
}

// Represents all fields a class can be queried over.
type ClassQuery struct {
	Id         int64
	Name       string
	Quarter    string
	Year       string
	Creator_id int64
}

// NewClass is used to create a new class object from an incoming HTTP request.
// It takes in the HTTP request in JSON format.
// It returns the class constructed and an error if one exists.
func NewClass(r *http.Request) (class Class, err error) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = errors.Wrapf(err, "Couldn't read request body")
		return
	}
	// Converts JSON to class
	json.Unmarshal(b, &class)
	return
}

// Equals is a custom comparator for two class objects on non-auto parameter fields.
// It takes in a class object representing the other value.
// It returns a boolean indicating the equality
func (class *Class) Equals(other Class) bool {
	return class.Name == other.Name &&
		class.Quarter == other.Quarter &&
		class.Year == other.Year &&
		class.Creator_id == other.Creator_id
}

// NewClassTable creates a new table within the database for housing
// all class objects.
// It takes in a reference to an open database connection.
// It returns the constructed table, and an error if one exists.
func NewClassTable(db *db.Db) (classTable ClassTable, err error) {
	// Ensure connection is alive
	if db == nil {
		err = errors.New("Invalid database connection")
		return
	}
	classTable.connection = db
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s(
			id SERIAL,
			name TEXT,
			description TEXT,
			quarter TEXT,
			year TEXT,
			creator_id INT,
			time_created TIMESTAMP DEFAULT now()
		);`, CLASS_TABLE)
	// Create the actual table
	if err = classTable.connection.CreateTable(query); err != nil {
		err = errors.Wrapf(err, "Could not create table on initialization")
	}
	return
}

// Insert will put a new class row within the table in the DB, verifying
// all fields are valid.
// It takes in the class object to insert and the user id to associate to the creator.
// It returns the new class as in the table and an error if one exists.
func (table *ClassTable) Insert(class Class) (new Class, err error) {
	classQuery := ClassQuery{Name: class.Name, Quarter: class.Quarter, Year: class.Year, Creator_id: class.Creator_id}
	data, err := table.connection.Insert(CLASS_TABLE, "AND", class, classQuery)
	if err != nil {
		return
	}
	utilities.Sugar.Infof("Class Data: %v", data)
	err = utilities.FillStruct(data, &new)
	return
}

// Get attempts to provide a generalized search through the class table based on the
// provided queries.
// It takes a class query for the queryable fields, and an operator such as "AND" or "OR" to
// define the context of the search.
// It returns all the found classes and an error if one exists.
func (table *ClassTable) Get(classQuery ClassQuery, op string) (classes []Class, err error) {
	allData, err := table.connection.Get(classQuery, op, CLASS_TABLE)
	if err != nil {
		return
	}
	for _, data := range allData {
		class := Class{}
		err = utilities.FillStruct(data, &class)
		if err != nil {
			return
		}
		classes = append(classes, class)
	}
	return
}

// GetByID uses the internal get mechanism for the table to find a class given an id to search on.
// It takes an ID as a string to convert to an integer to then search on.
// It returns the found class and an error if one exists.
func (table *ClassTable) GetByID(strId string) (class Class, err error) {
	data, err := table.connection.GetByID(strId, CLASS_TABLE)
	if err != nil {
		return
	}
	err = utilities.FillStruct(data, &class)
	return
}

// Update will update the class row in the table based on the incoming class object.
// It takes in an id to identify the class in the DB, and updates as a class object.
// It returns the updated class as in the DB, and an error if one exists.
func (table *ClassTable) Update(strId string, updates Class) (updated Class, err error) {
	data, err := table.connection.Update(strId, CLASS_TABLE, updates)
	if err != nil {
		return
	}
	err = utilities.FillStruct(data, &updated)
	return
}

// Delete permanently removes the class object from the table.
// It takes in an id for the class we wish to delete.
// It returns an error if one exists.
func (table *ClassTable) Delete(strId string) error {
	return table.connection.Delete(strId, CLASS_TABLE)
}

// DeleteAll permanently removes all class objects from the table.
// It returns an error if one exists.
func (table *ClassTable) DeleteAll() error {
	return table.connection.DeleteAll(CLASS_TABLE)
}
