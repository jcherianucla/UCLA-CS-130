package models

import (
	"bytes"
	"fmt"
	"github.com/jcherianucla/UCLA-CS-130/api/config/db"
	"github.com/jcherianucla/UCLA-CS-130/api/utilities"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

// The table name as set in the Postgres DB creation.
const (
	SUBMISSION_TABLE = "submissions"
)

// Represents the connection to the DB instance.
type SubmissionTable struct {
	connection *db.Db
}

// Represents an Submission row in the Submission table within the DB.
// Validator and json tags are used for convenient serialization and
// deserialization.
type Submission struct {
	Id            int64     `valid:"-" json:"-"`
	File          []byte    `valid:"required" json:"file"`
	Score         int       `valid:"-" json:"score"`
	Pre_results   []string  `valid:"-" json:"pre_results"`
	Post_results  []string  `valid:"-" json:"post_results"`
	Assignment_id int64     `valid:"required" json:"assignment_id"`
	User_id       int64     `valid:"required" json:"student_id"`
	Time_created  time.Time `valid:"-" json:"-"`
	Time_updated  time.Time `valid:"-" json:"-"`
}

// Represents all fields a submission can be queried over.
type SubmissionQuery struct {
	Id            int64
	Assignment_id int64
	User_id       int64
}

func NewSubmission(r *http.Request) (submission Submission, err error) {
	err = r.ParseMultipartForm(utilities.MAX_STORAGE)
	if err != nil {
		return
	}
	f, _, err := r.FormFile("upload")
	submission.File, err = convertToBytes(f)
	return
}

func (submission *Submission) Equals(other Submission) bool {
	return submission.Assignment_id == other.Assignment_id &&
		submission.User_id == other.User_id &&
		bytes.Equal(submission.File, other.File)
}

func NewSubmissionTable(db *db.Db) (submissionTable SubmissionTable, err error) {
	if db == nil {
		err = errors.New("Invalid database connection")
		return
	}
	submissionTable.connection = db
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s(
			id SERIAL,
			file BYTEA,
			score INT,
			pre_results TEXT[],
			post_results TEXT[],
			assignment_id INT,
			user_id INT,
			time_created TIMESTAMP DEFAULT now(),
			time_updated TIMESTAMP
		);`, SUBMISSION_TABLE)
	// Create the actual table
	if err = submissionTable.connection.CreateTable(query); err != nil {
		err = errors.Wrapf(err, "Could not create table on initialization")
	}
	return
}

func (table *SubmissionTable) Insert(submission Submission) (new Submission, err error) {
	submissionQuery := SubmissionQuery{Assignment_id: submission.Assignment_id, User_id: submission.User_id}
	data, err := table.connection.Insert(SUBMISSION_TABLE, "AND", submission, submissionQuery)
	if err != nil {
		return
	}
	err = utilities.FillStruct(data, &new)
	return
}

func (table *SubmissionTable) Get(submissionQuery SubmissionQuery, op string) (submissions []Submission, err error) {
	allData, err := table.connection.Get(submissionQuery, op, SUBMISSION_TABLE)
	if err != nil {
		return
	}
	for _, data := range allData {
		submission := Submission{}
		err = utilities.FillStruct(data, &submission)
		if err != nil {
			return
		}
		submissions = append(submissions, submission)
	}
	return
}

func (table *SubmissionTable) GetByID(strId string) (submission Submission, err error) {
	data, err := table.connection.GetByID(strId, SUBMISSION_TABLE)
	if err != nil {
		return
	}
	err = utilities.FillStruct(data, &submission)
	return
}

func (table *SubmissionTable) Update(strId string, updates Submission) (updated Submission, err error) {
	data, err := table.connection.Update(strId, SUBMISSION_TABLE, updates)
	if err != nil {
		return
	}
	err = utilities.FillStruct(data, &updated)
	return
}

func (table *SubmissionTable) Delete(strId string) error {
	return table.connection.Delete(strId, SUBMISSION_TABLE)
}

func (table *SubmissionTable) DeleteAll() error {
	return table.connection.DeleteAll(SUBMISSION_TABLE)
}
