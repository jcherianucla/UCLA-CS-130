package tests

import (
	"errors"
	"github.com/jcherianucla/UCLA-CS-130/api/app/models"
	"github.com/jcherianucla/UCLA-CS-130/api/config/db"
	"github.com/jcherianucla/UCLA-CS-130/api/utilities"
	"strconv"
	"strings"
	"testing"
	"time"
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

var before time.Time
var after time.Time

// Test assignment model to be used throughout
var ta models.Assignment = models.Assignment{
	Name:        "Project 1",
	Description: "Create a fiery backend in Golang",
	Lang:        utilities.SetLanguage("C++"),
	Class_id:    1,
}

var updated models.Assignment = models.Assignment{
	Name:        "Project 1",
	Description: "Create a fiery backend in Golang",
	Lang:        utilities.SetLanguage("Java"),
	Class_id:    1,
}

// The assignment table reference to act on
var at models.AssignmentTable

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
	// Create a assignment table
	at, _ = models.NewAssignmentTable(&db)
	// Set deadlines
	before, _ = time.Parse(utilities.TIME_FORMAT, "12-09-17 23:55 (PST)")
	ta.Deadline = before
	after, _ = time.Parse(utilities.TIME_FORMAT, "12-10-20 23:55 (PST)")
	updated.Deadline = after
}

// Represents a basic test object that all tests must conform to.
type test struct {
	name        string
	isErr       bool
	expectedErr error
}

// Represents a test on a assignment object.
type assignmentTest struct {
	base       test
	assignment models.Assignment
}

// Represents a test on a query object.
type queryTest struct {
	base  test
	query models.AssignmentQuery
}

func getAssignmentId(assignment models.Assignment) int64 {
	assignments, _ := at.Get(models.AssignmentQuery{Name: assignment.Name}, "")
	return assignments[0].Id
}

// TestAssignmentInsert will test the insert method for assignment models amongst multiple scenarios. It uses table based testing. Total number of tests: 3.
// It takes in a testing framework object to run tests on.
func TestAssignmentInsert(t *testing.T) {
	cases := []assignmentTest{
		{test{"Missing deadline", true, fieldErr}, models.Assignment{Name: "HW 1", Description: "It's lit", Lang: utilities.SetLanguage("C++"), Class_id: 1}},
		{test{"Valid assignment", false, noErr}, ta},
		{test{"Existing assignment", true, existsErr}, ta},
	}

	// Clean
	_ = at.DeleteAll()
	for _, tc := range cases {
		t.Run(tc.base.name, func(t *testing.T) {
			result := models.Assignment{}
			result, err := at.Insert(tc.assignment)
			if tc.base.isErr {
				if !strings.Contains(err.Error(), tc.base.expectedErr.Error()) {
					t.Errorf("Errors do not match up")
				}
			} else {
				if !tc.assignment.Equals(result) {
					t.Errorf("Assignments do not match up")
				}
			}
		})
	}
}

// TestAssignmentGet simply tests query validation. This is because for the most part, it used internally with error handling done in other parts of the model functionality. Total number of tests: 1.
// It takes in a testing framework object to run tests on.
func TestAssignmentGet(t *testing.T) {
	cases := []queryTest{
		{test{"Invalid query", true, invalidQueryErr}, models.AssignmentQuery{}},
	}

	for _, tc := range cases {
		t.Run(tc.base.name, func(t *testing.T) {
			_, err := at.Get(tc.query, "")
			if tc.base.isErr {
				if !strings.Contains(err.Error(), tc.base.expectedErr.Error()) {
					t.Errorf("Errors do not match up")
				}
			}
		})
	}
}

// TestAssignmentUpdate tests the update method for the assignment model. If the update works on any field, then the overall update method works as it is the same logic for each field. Total number of tests: 2.
// It takes in a testing framework object to run tests on.
func TestAssignmentUpdate(t *testing.T) {
	// Get the relevant ID for assignment
	id := getAssignmentId(ta)
	validId := strconv.FormatInt(id, 10)
	invalidId := strconv.FormatInt(id-1, 10)
	cases := []assignmentTest{
		{test{"No assignment", true, noModelErr}, ta},
		{test{"Valid update", false, noErr}, updated},
	}

	for _, tc := range cases {
		t.Run(tc.base.name, func(t *testing.T) {
			result := models.Assignment{}
			if tc.base.isErr {
				_, err := at.Update(invalidId, tc.assignment)
				if !strings.Contains(err.Error(), tc.base.expectedErr.Error()) {
					t.Errorf("Errors do not match up")
				}
			} else {
				result, _ = at.Update(validId, tc.assignment)
				if !tc.assignment.Equals(result) {
					t.Errorf("Updated assignments do not match up")
				}
			}
		})
	}
}

// TestAssignmentDelete tests that the assignment model delete function actually deletes the assignment from the database. Total number of tests: 2.
// It takes in a testing framework object to run tests on.
func TestAssignmentDelete(t *testing.T) {
	// Get the relevant ID for assignment
	id := getAssignmentId(updated)
	validId := strconv.FormatInt(id, 10)
	invalidId := strconv.FormatInt(id-1, 10)
	cases := []assignmentTest{
		{test{"No assignment", true, noModelErr}, ta},
		{test{"Valid delete", false, noErr}, updated},
	}

	for _, tc := range cases {
		t.Run(tc.base.name, func(t *testing.T) {
			if tc.base.isErr {
				err := at.Delete(invalidId)
				if !strings.Contains(err.Error(), tc.base.expectedErr.Error()) {
					t.Errorf("Errors do not match up")
				}
			} else {
				_ = at.Delete(validId)
				assignments, _ := at.Get(models.AssignmentQuery{Id: id}, "")
				if assignments != nil {
					t.Errorf("Did not delete assignment")
				}
			}
		})
	}
}

// TestDeadline tests that the utilities function for comparing deadlines works correctly. Total number of tests: 2.
// It takes in a testing framework object to run tests on.
func TestDeadline(t *testing.T) {
	cases := []assignmentTest{
		{test{"After deadline", false, noErr}, ta},
		{test{"Before deadline", false, noErr}, updated},
	}

	for _, tc := range cases {
		t.Run(tc.base.name, func(t *testing.T) {
			if strings.Contains(tc.base.name, "Before") {
				if !utilities.BeforeDeadline(tc.assignment.Deadline) {
					t.Errorf("Should output true for before deadline")
				}
			} else {
				if utilities.BeforeDeadline(tc.assignment.Deadline) {
					t.Errorf("Should output false for after deadline")
				}
			}
		})
	}
}
