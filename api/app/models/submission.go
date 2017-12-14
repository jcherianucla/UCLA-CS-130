package models

import (
	"bytes"
	"fmt"
	"github.com/jcherianucla/UCLA-CS-130/api/config/db"
	"github.com/jcherianucla/UCLA-CS-130/api/utilities"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
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
	Score         int64     `valid:"-" json:"score"`
	Pre_results   []string  `valid:"-" json:"pre_results"`
	Post_results  []string  `valid:"-" json:"post_results"`
	Assignment_id int64     `valid:"required" json:"assignment_id"`
	User_id       int64     `valid:"required" json:"student_id"`
	Time_created  time.Time `valid:"-" json:"-"`
	Time_updated  time.Time `valid:"-" json:"updated_at"`
}

// Represents all fields a submission can be queried over.
type SubmissionQuery struct {
	Id            int64
	Assignment_id int64
	User_id       int64
}

// Exec acts as the execution engine for running the submitted grading/sanity scripts on the submitted file. Automatically runs in a sandboxed environment as it is housed within a docker container, isolated from the system.
// It takes in the script to run and the file to run on as byte slices, and the language of the associated submitted file.
// It returns a score (if grading script) and a set of results, and an error if one exists.
func Exec(grade_script, s_file []byte, lang utilities.Language) (score int64, results []string, err error) {
	// Create script file from contents
	err = ioutil.WriteFile("script.sh", grade_script, 0666)
	if err != nil {
		err = errors.Wrapf(err, "Failed to create script file")
		return
	}
	s_filepath := "submitted." + strings.ToLower(utilities.GetLanguage(lang))
	// Create submission file from contents
	err = ioutil.WriteFile(s_filepath, s_file, 0777)
	if err != nil {
		err = errors.Wrapf(err, "Failed to create submitted file")
		return
	}
	// Run script on submitted file
	out, err := exec.Command("bash", "script.sh", s_filepath).CombinedOutput()
	if err != nil {
		err = errors.Wrapf(err, "Couldn't execute script on submitted file")
		return
	}
	utilities.Sugar.Infof("Output: %v", out)
	outStr := string(out)
	res := strings.Split(outStr, "\n")
	// Expect score to be first line in output
	score, _ = strconv.ParseInt(res[0][len("Score: "):], 10, 64)
	results = res[1:]
	return

}

// NewSubmission creates a new submission object based on a request, assuming the request holding multipart form-data.
// It takes in the request to analyze.
// It returns the constructed submission and an error if one exists.
func NewSubmission(r *http.Request) (submission Submission, err error) {
	err = r.ParseMultipartForm(utilities.MAX_STORAGE)
	if err != nil {
		return
	}
	// Get the uploaded file
	f, _, err := r.FormFile("upload")
	submission.File, err = convertToBytes(f)
	return
}

// Equals is a custom comparator for two submission objects on non-auto parameter fields.
// It takes in an submission object representing the other value.
// It returns a boolean indicating the equality
func (submission *Submission) Equals(other Submission) bool {
	return submission.Assignment_id == other.Assignment_id &&
		submission.User_id == other.User_id &&
		bytes.Equal(submission.File, other.File)
}

// NewSubmissionTable creates a new table within the database for housing all submission objects.
// It takes in a reference to an open database connection.
// It returns the constructed table, and an error if one exists.
func NewSubmissionTable(db *db.Db) (submissionTable SubmissionTable, err error) {
	// Ensure connection is alive
	if db == nil {
		err = errors.New("Invalid database connection")
		return
	}
	submissionTable.connection = db
	// Construct query
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s(
			id SERIAL,
			file BYTEA,
			score INT,
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

// Insert will put a new submission row within the table in the DB, verifying all fields are valid.
// It takes in the submission object to insert.
// It returns the new submission as in the table and an error if one exists.
func (table *SubmissionTable) Insert(submission Submission) (new Submission, err error) {
	// Query to ensure the same object doesn't already exist
	submissionQuery := SubmissionQuery{Assignment_id: submission.Assignment_id, User_id: submission.User_id}
	data, err := table.connection.Insert(SUBMISSION_TABLE, "AND", submission, submissionQuery)
	if err != nil {
		return
	}
	err = utilities.FillStruct(data, &new)
	utilities.Sugar.Infof("Insert submission: %v", new)
	return
}

// Get attempts to provide a generalized search through the submission table based on the provided queries.
// It takes a submission query for the queryable fields, and an operator such as "AND" or "OR" to define the context of the search.
// It returns all the found submissions and an error if one exists.
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

// GetByID finds the submission given the user id and the assignment id.
// It takes in a user id and assignment id to identify the submission.
// It returns the found submission or an error if one exists.
func (table *SubmissionTable) GetByID(uid, aid string) (submission Submission, err error) {
	_uid, _ := strconv.ParseInt(uid, 10, 64)
	_aid, _ := strconv.ParseInt(aid, 10, 64)
	objs, err := table.Get(SubmissionQuery{User_id: _uid, Assignment_id: _aid}, "AND")
	if err != nil {
		return
	}
	if len(objs) == 0 {
		err = errors.New("Couldn't find object")
	} else {
		submission = objs[0]
	}
	return
}

// Update will update the submission row in the table based on the incoming submission object.
// It takes in an id to identify the submission in the DB, and updates as a submission object.
// It returns the updated submission as in the DB, and an error if one exists.
func (table *SubmissionTable) Update(strId string, updates Submission) (updated Submission, err error) {
	data, err := table.connection.Update(strId, SUBMISSION_TABLE, updates)
	if err != nil {
		return
	}
	err = utilities.FillStruct(data, &updated)
	return
}

// Delete permanently removes the specified submission from the database.
// It takes in the submission id to delete.
// It returns an error if one exists.
func (table *SubmissionTable) Delete(strId string) error {
	return table.connection.Delete(strId, SUBMISSION_TABLE)
}

// DeleteAll permanently removes all submissions from the database.
// It returns an error if one exists.
func (table *SubmissionTable) DeleteAll() error {
	return table.connection.DeleteAll(SUBMISSION_TABLE)
}
