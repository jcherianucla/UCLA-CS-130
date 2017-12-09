package models

import (
	"encoding/csv"
	"fmt"
	"github.com/jcherianucla/UCLA-CS-130/api/config/db"
	"github.com/jcherianucla/UCLA-CS-130/api/utilities"
	"github.com/pkg/errors"
	"io"
	"strconv"
	"strings"
	"time"
)

// The table name as set in the Postgres DB creation.
const (
	ENROLLED_TABLE = "enrolled"
)

// Represents the connection to the DB instance.
type EnrolledTable struct {
	connection *db.Db
}

// Represents a Student row in the Enrolled table within the DB
// Validator tags are used for convenient serialization and
// deserialization.
type Student struct {
	Id int64 `valid:"-"`
	// Distinguishes privileges between a student and professor
	User_id      int64     `valid:"required"`
	Class_id     int64     `valid:"required"`
	Time_created time.Time `valid:"-"`
}

func (student *Student) Equals(other Student) bool {
	return student.User_id == other.User_id &&
		student.Class_id == other.Class_id
}

// NewStudent is used to create a new user object from a CSV record.
// It takes in the record mapping the column names to values.
// It returns the constructed user, and an error if it exists.
func NewStudent(record map[string]interface{}) (user User, err error) {
	err = utilities.FillStruct(record, &user)
	// Students aren't professors
	user.Is_professor = false
	user.Password = utilities.DEFAULT_PASSWORD
	return
}

func NewEnrolledTable(db *db.Db) (enrolledTable EnrolledTable, err error) {
	// Ensure connection is alive
	if db == nil {
		err = errors.New("Invalid database connection")
		return
	}
	enrolledTable.connection = db
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s(
			id SERIAL,
			user_id INT,
			class_id INT,
			time_created TIMESTAMP DEFAULT now()
		);`, ENROLLED_TABLE)
	// Create the actual table
	if err = enrolledTable.connection.CreateTable(query); err != nil {
		err = errors.Wrapf(err, "User Table creation query failed")
	}
	return
}

func (table *EnrolledTable) Insert(classId string, r io.Reader) (err error) {
	record := make(map[string]interface{})
	rows := csv.NewReader(r)
	student := Student{}
	var cols []string
	i := 0
	for {
		curr, err := rows.Read()
		if err == io.EOF {
			break
		}
		if i == 0 {
			// First row is the columns
			cols = curr
			for _, key := range curr {
				record[strings.Title(strings.ToLower(key))] = nil
			}
		} else {
			// Fill up map with vals
			for i, val := range curr {
				record[strings.Title(strings.ToLower(cols[i]))] = val
			}
			// Create user from csv record
			u, err := NewStudent(record)
			// Check for existence
			found, err := LayerInstance().User.GetByEmail(u.Email)
			if err != nil && strings.Contains(err.Error(), "Couldn't find user with email") {
				// Create the user if it doesn't exist
				u, err = LayerInstance().User.Insert(u)
			} else {
				u = found
			}
			cid, _ := strconv.ParseInt(classId, 10, 64)
			student = Student{User_id: u.Id, Class_id: cid}
			// Insert Student
			_, err = table.connection.Insert(ENROLLED_TABLE, "AND", student, student)
		}
		i++
	}
	return
}

func (table *EnrolledTable) GetStudents(classId string) (students []User, err error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE class_id=$1", ENROLLED_TABLE)

	utilities.Sugar.Infof("SQL Query: %s", query)
	utilities.Sugar.Infof("Value: %v", classId)

	stmt, err := table.connection.Pool.Prepare(query)
	if err != nil {
		err = errors.Wrapf(err, "Get query preparation failed")
		return
	}
	rows, err := stmt.Query(classId)
	if err != nil {
		err = errors.Wrapf(err, "Get query failed to execute")
		return
	}
	for rows.Next() {
		var student Student
		err = rows.Scan(&student.Id, &student.User_id, &student.Class_id)
		if err != nil {
			err = errors.Wrapf(err, "Get query failed to execute")
			return
		}
		user, err := LayerInstance().User.GetByID(strconv.FormatInt(student.User_id, 10))
		if err == nil {
			students = append(students, user)
		}
	}
	return
}

func (table *EnrolledTable) Unenroll(userId string) error {
	return table.connection.DeleteByCol("user_id", userId, ENROLLED_TABLE)
}

func (table *EnrolledTable) DeleteClass(classId string) error {
	return table.connection.DeleteByCol("class_id", classId, ENROLLED_TABLE)
}
