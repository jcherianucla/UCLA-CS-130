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
	Id            int64              `valid:"-" json:"id"`
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

// convertToBytes converts a file object received from the client
// into a byte slice as needed by the database.
// It takes in the stream of the file object.
// It returns the file as a byte slice and an error if one exists.
func convertToBytes(r io.Reader) (f []byte, err error) {
	buf := bytes.NewBuffer(nil)
	if _, err = io.Copy(buf, r); err != nil {
		err = errors.Wrapf(err, "Failed to copy over file contents to DB")
		return
	}
	f = buf.Bytes()
	return
}

// storeFile will get the file from the request, and store it into the data
// object used to construct the assignment struct.
// It takes in a request object to grab the file from, and a reference to the
// map of data, and the key to store into the map.
// It returns an error if one exists.
func storeFile(r *http.Request, m *map[string]interface{}, key string) error {
	f, _, err := r.FormFile(key)
	if err != nil {
		return err
	}
	if f != nil {
		defer f.Close()
		(*m)[strings.Title(key)], err = convertToBytes(f)
	}
	return err
}

// NewAssignment creates a new assignment object based on a request, assuming the request holding multipart form-data.
// It takes in the request to analyze.
// It returns the constructed assignment and an error if one exists.
func NewAssignment(r *http.Request) (assignment Assignment, err error) {
	err = r.ParseMultipartForm(utilities.MAX_STORAGE)
	utilities.Sugar.Infof("Form: %v", r.PostForm)
	if err != nil {
		return
	}
	m := make(map[string]interface{})
	// Fill up basic information
	for k, v := range r.PostForm {
		k = strings.Title(k)
		if k == "Deadline" {
			t, _ := time.Parse(utilities.TIME_FORMAT, v[0])
			m[k] = t
		} else if k == "Language" {
			m["Lang"] = utilities.SetLanguage(v[0])
		} else {
			m[k] = v[0]
		}
	}
	storeFile(r, &m, "grade_script")
	storeFile(r, &m, "sanity_script")
	// Create assignment object
	err = utilities.FillStruct(m, &assignment)
	utilities.Sugar.Infof("Assignment object: %v", assignment)
	return
}

// Equals is a custom comparator for two assignment objects on non-auto parameter fields.
// It takes in an assignment object representing the other value.
// It returns a boolean indicating the equality
func (assign *Assignment) Equals(other Assignment) bool {
	return assign.Name == other.Name &&
		assign.Lang == other.Lang &&
		assign.Class_id == other.Class_id
}

// NewAssignmentTable creates a new table within the database for housing all assignment objects.
// It takes in a reference to an open database connection.
// It returns the constructed table, and an error if one exists.
func NewAssignmentTable(db *db.Db) (assignTable AssignmentTable, err error) {
	// Ensure connection is alive
	if db == nil {
		err = errors.New("Invalid database connection")
		return
	}
	assignTable.connection = db
	// Construct query
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

// Insert will put a new assignment row within the table in the DB, verifying all fields are valid.
// It takes in the assignment object to insert.
// It returns the new assignment as in the table and an error if one exists.
func (table *AssignmentTable) Insert(assign Assignment) (new Assignment, err error) {
	// Query to ensure the same object doesn't already exist
	assignQuery := AssignmentQuery{Name: assign.Name, Lang: assign.Lang, Class_id: assign.Class_id}
	data, err := table.connection.Insert(ASSIGNMENT_TABLE, "AND", assign, assignQuery)
	if err != nil {
		return
	}
	err = utilities.FillStruct(data, &new)
	return
}

// Get attempts to provide a generalized search through the assignment table based on the provided queries.
// It takes a assignment query for the queryable fields, and an operator such as "AND" or "OR" to define the context of the search.
// It returns all the found assignments and an error if one exists.
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

// GetByID finds the assignment with the specified assignment id.
// It takes an ID as a string to convert to an integer to then search on.
// It returns the found assignment and an error if one exists.
func (table *AssignmentTable) GetByID(strId string) (assign Assignment, err error) {
	data, err := table.connection.GetByID(strId, ASSIGNMENT_TABLE)
	if err != nil {
		return
	}
	err = utilities.FillStruct(data, &assign)
	return
}

// Update will update the assignment row in the table based on the incoming assignment object.
// It takes in an id to identify the assignment in the DB, and updates as a assignment object.
// It returns the updated assignment as in the DB, and an error if one exists.
func (table *AssignmentTable) Update(strId string, updates Assignment) (updated Assignment, err error) {
	data, err := table.connection.Update(strId, ASSIGNMENT_TABLE, updates)
	if err != nil {
		return
	}
	err = utilities.FillStruct(data, &updated)
	return
}

// Delete permanently removes the assignment object from the table, along with any submissions made to the assignment.
// It takes in an id for the assignment we wish to delete.
// It returns an error if one exists.
func (table *AssignmentTable) Delete(strId string) (err error) {
	aid, _ := strconv.ParseInt(strId, 10, 64)
	// Find all submissions related to the assignment
	submissions, err := LayerInstance().Submission.Get(SubmissionQuery{Assignment_id: aid}, "")
	if err != nil {
		return
	}
	// Delete all submissions related to the assignment
	for _, submission := range submissions {
		sid := strconv.FormatInt(submission.Id, 10)
		err = LayerInstance().Submission.Delete(sid)
		if err != nil {
			return
		}
	}
	// Delete assignment itself
	err = table.connection.Delete(strId, ASSIGNMENT_TABLE)
	return
}

// DeleteAll will perform a relational delete on all assignments within the database by calling delete on the individual assignment.
// It returns an error if one exists.
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
