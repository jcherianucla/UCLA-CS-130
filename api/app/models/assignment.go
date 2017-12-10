package models

import (
	"bytes"
	"fmt"
	"github.com/jcherianucla/UCLA-CS-130/api/config/db"
	"github.com/jcherianucla/UCLA-CS-130/api/utilities"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"strconv"
	"strings"
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

// Represents an Assignment row in the Assignment table within the DB.
// Validator and json tags are used for convenient serialization and
// deserialization.
type Assignment struct {
	Id            int64              `valid:"-" json:"-"`
	Name          string             `valid:"required" json:"name"`
	Description   string             `valid:"-" json:"description"`
	Deadline      time.Time          `valid:"required" json:"deadline"`
	Lang          utilities.Language `valid:"required" json:"language"`
	Grade_script  []byte             `valid:"-" json:"-"`
	Sanity_script []byte             `valid:"-" json:"-"`
	Class_id      int64              `valid:"-" json:"class_id"`
	Time_created  time.Time          `valid:"-" json:"-"`
}

// Represents all fields an assignment can be queried over.
type AssignmentQuery struct {
	Id       int64
	Name     string
	Lang     utilities.Language
	Class_id int64
}

func convertToBytes(r io.Reader) (f []byte, err error) {
	buf := bytes.NewBuffer(nil)
	if _, err = io.Copy(buf, r); err != nil {
		err = errors.Wrapf(err, "Failed to copy over file contents to DB")
		return
	}
	f = buf.Bytes()
	return
}

func storeFile(r *http.Request, m *map[string]interface{}, key string) error {
	f, _, err := r.FormFile(key)
	if err != nil {
		return err
	}
	defer f.Close()
	(*m)[strings.Title(key)], err = convertToBytes(f)
	return err
}

func NewAssignment(r *http.Request) (assignment Assignment, err error) {
	err = r.ParseMultipartForm(utilities.MAX_STORAGE)
	if err != nil {
		return
	}
	m := make(map[string]interface{})
	// Fill up basic information
	for k, v := range r.PostForm {
		k = strings.Title(k)
		if k == "Deadline" {
			m[k], err = time.Parse("01-02-06 15:04", v[0])
		} else if k == "Language" {
			m["Lang"] = utilities.SetLanguage(v[0])
		} else {
			m[k] = v[0]
		}
	}
	storeFile(r, &m, "grade_script")
	storeFile(r, &m, "sanity_script")
	err = utilities.FillStruct(m, &assignment)
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

func (table *AssignmentTable) DeleteAll() (err error) {
	query := fmt.Sprintf("SELECT id FROM %s", ASSIGNMENT_TABLE)

	utilities.Sugar.Infof("SQL Query: %v", query)

	rows, err := table.connection.Pool.Query(query)
	if err != nil {
		err = errors.Wrapf(err, "Delete all query failed")
		return
	}
	// Delete all the assignments by calling the relational delete
	for rows.Next() {
		var id int64
		if err = rows.Scan(&id); err != nil {
			err = errors.Wrapf(err, "Failed to scan into id")
			return
		}
		if err = table.Delete(strconv.FormatInt(id, 10)); err != nil {
			return
		}
	}
	return
}
