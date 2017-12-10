package models

import (
	"fmt"
	"github.com/gorilla/schema"
	"github.com/jcherianucla/UCLA-CS-130/api/config/db"
	"github.com/jcherianucla/UCLA-CS-130/api/utilities"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

// The table name as set in the Postgres DB creation.
const (
	ASSIGNMENT_TABLE = "assignments"
)

// Represents the connection to the DB instance.
type AssignmentTable struct {
	connection *db.Db
}

// Represents an enum for languages
type Language int

const (
	Cpp Language = iota
	C
	Java
)

func GetLanguage(lang Language) string {
	switch lang {
	case Cpp:
		return "C++"
	case C:
		return "C"
	case Java:
		return "Java"
	default:
		return "Unknown"
	}
}

func SetLanguage(lang string) Language {
	switch lang {
	case "C++":
		return Cpp
	case "C":
		return C
	case "Java":
		return Java
	default:
		return -1
	}
}

// Represents an Assignment row in the Assignment table within the DB.
// Validator and json tags are used for convenient serialization and
// deserialization.
type Assignment struct {
	Id            int64     `valid:"-" json:"-"`
	Name          string    `valid:"required" json:"name"`
	Description   string    `valid:"-" json:"description"`
	Deadline      time.Time `valid:"required" json:"deadline"`
	Lang          Language  `valid:"required" json:"language"`
	Grade_script  []byte    `valid:"-" json:"-"`
	Sanity_script []byte    `valid:"-" json:"-"`
	Class_id      int64     `valid:"-" json:"class_id"`
	Time_created  time.Time `valid:"-" json:"-"`
}

// Represents all fields an assignment can be queried over.
type AssignmentQuery struct {
	Id       int64
	Name     string
	Lang     Language
	Class_id int64
}

func NewAssignment(r *http.Request) (assignment Assignment, err error) {
	err = r.ParseMultipartForm(utilities.MAX_STORAGE)
	if err != nil {
		return
	}
	decoder := schema.NewDecoder()
	err = decoder.Decode(&assignment, r.PostForm)
	return
}

func (assign *Assignment) Equals(other Assignment) bool {
	return assign.Name == other.Name &&
		assign.Lang == other.Lang &&
		assign.Class_id == other.Class_id
}

func NewAssignmentTable(db *db.Db) (assignTable AssignmentTable, err error) {
	if db == nil {
		err = errors.New("Invalid database connection")
		return
	}
	assignTable.connection = db
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s(
			id SERIAL,
			name TEXT,
			description TEXT,
			deadline TIMESTAMP,
			lang INT,
			grade_script BYTEA,
			sanity_script BYTEA,
			class_id INT,
			time_created TIMESTAMP DEFAULT now()
		);`, ASSIGNMENT_TABLE)
	// Create the actual table
	if err = assignTable.connection.CreateTable(query); err != nil {
		err = errors.Wrapf(err, "Could not create table on initialization")
	}
	return
}

func (table *AssignmentTable) Insert(assign Assignment) (new Assignment, err error) {
	assignQuery := AssignmentQuery{Name: assign.Name, Lang: assign.Lang, Class_id: assign.Class_id}
	data, err := table.connection.Insert(ASSIGNMENT_TABLE, "AND", assign, assignQuery)
	if err != nil {
		return
	}
	err = utilities.FillStruct(data, &new)
	return
}

func (table *AssignmentTable) Get(assignQuery AssignmentQuery, op string) (assigns []Assignment, err error) {
	allData, err := table.connection.Get(assignQuery, op, ASSIGNMENT_TABLE)
	if err != nil {
		return
	}
	for _, data := range allData {
		assign := Assignment{}
		err = utilities.FillStruct(data, &assign)
		if err != nil {
			return
		}
		assigns = append(assigns, assign)
	}
	return

}
func (table *AssignmentTable) GetByID(strId string) (assign Assignment, err error) {
	data, err := table.connection.GetByID(strId, ASSIGNMENT_TABLE)
	if err != nil {
		return
	}
	err = utilities.FillStruct(data, &assign)
	return
}

func (table *AssignmentTable) Update(strId string, updates Assignment) (updated Assignment, err error) {
	data, err := table.connection.Update(strId, ASSIGNMENT_TABLE, updates)
	if err != nil {
		return
	}
	err = utilities.FillStruct(data, &updated)
	return
}

func (table *AssignmentTable) Delete(strId string) error {
	return table.connection.Delete(strId, ASSIGNMENT_TABLE)
}

func (table *AssignmentTable) DeleteAll() error {
	return table.connection.DeleteAll(ASSIGNMENT_TABLE)
}
