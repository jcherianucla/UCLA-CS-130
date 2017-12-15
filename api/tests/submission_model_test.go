package tests

import (
	"errors"
	"github.com/jcherianucla/UCLA-CS-130/api/app/models"
	"github.com/jcherianucla/UCLA-CS-130/api/config/db"
	"github.com/jcherianucla/UCLA-CS-130/api/utilities"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
)

const (
	// Name of the test Postgres database
	DB_NAME = "gp_test"
)

var (
	noErr           = errors.New("")
	fieldErr        = errors.New("has invalid fields")
	existsErr       = errors.New("already exists")
	noModelErr      = errors.New("Failed to find")
	invalidQueryErr = errors.New("Get query preparation failed")
)

var grade_script []byte

// Test submission model to be used throughout
var ts models.Submission = models.Submission{
	User_id:       1,
	Score:         90,
	Assignment_id: 1,
}

var updated models.Submission = models.Submission{
	User_id:       1,
	Score:         100,
	Assignment_id: 1,
}

// The submission table reference to act on
var st models.SubmissionTable

// init will begin the setup of the test database before tests are run.
func init() {
	// Open testing db connection
	db, _ := db.New(db.Config{
		utilities.DB_HOST,
		utilities.DB_PORT,
		utilities.DB_USER,
		utilities.DB_PASSWORD,
		DB_NAME,
	})
	// Create a submission table
	st, _ = models.NewSubmissionTable(&db)
	cwd, _ := os.Getwd()
	// Read in file for submission
	f, _ := ioutil.ReadFile(cwd + "/files/correct_submission.cpp")
	// Set submission files
	ts.File = f
	updated.File = f
	grade_script, _ = ioutil.ReadFile(cwd + "/files/grade.sh")
}

// Represents a basic test object that all tests must conform to.
type test struct {
	name        string
	isErr       bool
	expectedErr error
}

// Represents a test on a submission object.
type submissionTest struct {
	base       test
	submission models.Submission
}

// Represents a test on a query object.
type queryTest struct {
	base  test
	query models.SubmissionQuery
}

func getSubmissionId(submission models.Submission) int64 {
	uid := strconv.FormatInt(submission.User_id, 10)
	aid := strconv.FormatInt(submission.Assignment_id, 10)
	s, _ := st.GetByID(uid, aid)
	return s.Id
}

// TestSubmissionInsert will test the insert method for submission models amongst multiple scenarios. It uses table based testing. Total number of tests: 3.
// It takes in a testing framework object to run tests on.
func TestSubmissionInsert(t *testing.T) {
	cases := []submissionTest{
		{test{"Missing file", true, fieldErr}, models.Submission{User_id: 1, Assignment_id: 1}},
		{test{"Valid submission", false, noErr}, ts},
		{test{"Existing submission", true, existsErr}, ts},
	}

	// Clean
	_ = st.DeleteAll()
	for _, tc := range cases {
		t.Run(tc.base.name, func(t *testing.T) {
			result := models.Submission{}
			result, err := st.Insert(tc.submission)
			if tc.base.isErr {
				if !strings.Contains(err.Error(), tc.base.expectedErr.Error()) {
					t.Errorf("Errors do not match up")
				}
			} else {
				if !tc.submission.Equals(result) {
					t.Logf("Result: %v", result)
					t.Logf("Actual: %v", tc.submission)
					t.Errorf("Submissions do not match up")
				}
			}
		})
	}
}

// TestSubmissionGet simply tests query validation. This is because for the most part, it used internally with error handling done in other parts of the model functionality. Total number of tests: 1.
// It takes in a testing framework object to run tests on.
func TestSubmissionGet(t *testing.T) {
	cases := []queryTest{
		{test{"Invalid query", true, invalidQueryErr}, models.SubmissionQuery{}},
	}

	for _, tc := range cases {
		t.Run(tc.base.name, func(t *testing.T) {
			_, err := st.Get(tc.query, "")
			if tc.base.isErr {
				if !strings.Contains(err.Error(), tc.base.expectedErr.Error()) {
					t.Errorf("Errors do not match up")
				}
			}
		})
	}
}

// TestSubmissionUpdate tests the update method for the submission model. If the update works on any field, then the overall update method works as it is the same logic for each field. Total number of tests: 2.
// It takes in a testing framework object to run tests on.
func TestSubmissionUpdate(t *testing.T) {
	// Get the relevant ID for submission
	id := getSubmissionId(ts)
	validId := strconv.FormatInt(id, 10)
	invalidId := strconv.FormatInt(id-1, 10)
	cases := []submissionTest{
		{test{"No submission", true, noModelErr}, ts},
		{test{"Valid update", false, noErr}, updated},
	}

	for _, tc := range cases {
		t.Run(tc.base.name, func(t *testing.T) {
			result := models.Submission{}
			if tc.base.isErr {
				_, err := st.Update(invalidId, tc.submission)
				if !strings.Contains(err.Error(), tc.base.expectedErr.Error()) {
					t.Errorf("Errors do not match up")
				}
			} else {
				result, _ = st.Update(validId, tc.submission)
				if !tc.submission.Equals(result) {
					t.Errorf("Updated submissions do not match up")
				}
			}
		})
	}
}

// TestSubmissionDelete tests that the submission model delete function actually deletes the submission from the database. Total number of tests: 2.
// It takes in a testing framework object to run tests on.
func TestSubmissionDelete(t *testing.T) {
	// Get the relevant ID for submission
	id := getSubmissionId(updated)
	validId := strconv.FormatInt(id, 10)
	invalidId := strconv.FormatInt(id-1, 10)
	cases := []submissionTest{
		{test{"No submission", true, noModelErr}, ts},
		{test{"Valid delete", false, noErr}, updated},
	}

	for _, tc := range cases {
		t.Run(tc.base.name, func(t *testing.T) {
			if tc.base.isErr {
				err := st.Delete(invalidId)
				if !strings.Contains(err.Error(), tc.base.expectedErr.Error()) {
					t.Errorf("Errors do not match up")
				}
			} else {
				_ = st.Delete(validId)
				submissions, _ := st.Get(models.SubmissionQuery{Id: id}, "")
				if submissions != nil {
					t.Errorf("Did not delete submission")
				}
			}
		})
	}
}

// TestExecution runs a minimal test to ensure that execution of script with submission takes place. Total number of tests: 1.
// It takes in a testing framework object to run tests on.
func TestExecution(t *testing.T) {
	cases := []submissionTest{
		{test{"Valid execution", false, noErr}, ts},
	}

	for _, tc := range cases {
		t.Run(tc.base.name, func(t *testing.T) {
			s, _, _ := models.Exec(grade_script, tc.submission.File, utilities.SetLanguage("C++"))
			if s != 100 {
				t.Errorf("Did not get correct score")
			}
		})
	}
}
