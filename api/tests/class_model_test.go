package tests

import (
	"errors"
	"github.com/jcherianucla/UCLA-CS-130/api/app/models"
	"github.com/jcherianucla/UCLA-CS-130/api/config/db"
	"github.com/jcherianucla/UCLA-CS-130/api/utilities"
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

// Test class model to be used throughout
var tc models.Class = models.Class{
	Name:        "CS 132",
	Description: "Compiler Construction at UCLA",
	Creator_id:  1,
}

var updated models.Class = models.Class{
	Name:        "CS 130",
	Description: "Software Engineering at UCLA",
	Creator_id:  1,
}

// The class table reference to act on
var ct models.ClassTable

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
	// Create a class table
	ct, _ = models.NewClassTable(&db)
}

// Represents a basic test object that all tests must conform to.
type test struct {
	name        string
	isErr       bool
	expectedErr error
}

// Represents a test on a class object.
type classTest struct {
	base  test
	class models.Class
}

// Represents a test on a query object.
type queryTest struct {
	base  test
	query models.ClassQuery
}

func getClassId(class models.Class) int64 {
	classes, _ := ct.Get(models.ClassQuery{Name: class.Name}, "")
	return classes[0].Id
}

// TestClassInsert will test the insert method for class models amongst multiple scenarios. It uses table based testing. Total number of tests: 3.
// It takes in a testing framework object to run tests on.
func TestClassInsert(t *testing.T) {
	cases := []classTest{
		{test{"Missing name", true, fieldErr}, models.Class{Description: "Hello world"}},
		{test{"Valid class", false, noErr}, tc},
		{test{"Existing class", true, existsErr}, tc},
	}

	// Clean
	_ = ct.DeleteAll()
	for _, tc := range cases {
		t.Run(tc.base.name, func(t *testing.T) {
			result := models.Class{}
			result, err := ct.Insert(tc.class)
			if tc.base.isErr {
				if !strings.Contains(err.Error(), tc.base.expectedErr.Error()) {
					t.Errorf("Errors do not match up")
				}
			} else {
				if !tc.class.Equals(result) {
					t.Errorf("Classes do not match up")
				}
			}
		})
	}
}

// TestClassGet simply tests query validation. This is because for the most part, it used internally with error handling done in other parts of the model functionality. Total number of tests: 1.
// It takes in a testing framework object to run tests on.
func TestClassGet(t *testing.T) {
	cases := []queryTest{
		{test{"Invalid query", true, invalidQueryErr}, models.ClassQuery{}},
	}

	for _, tc := range cases {
		t.Run(tc.base.name, func(t *testing.T) {
			_, err := ct.Get(tc.query, "")
			if tc.base.isErr {
				if !strings.Contains(err.Error(), tc.base.expectedErr.Error()) {
					t.Errorf("Errors do not match up")
				}
			}
		})
	}
}

// TestClassUpdate tests the update method for the class model. If the update works on any field, then the overall update method works as it is the same logic for each field. Total number of tests: 2.
// It takes in a testing framework object to run tests on.
func TestClassUpdate(t *testing.T) {
	// Get the relevant ID for class
	id := getClassId(tc)
	validId := strconv.FormatInt(id, 10)
	invalidId := strconv.FormatInt(id-1, 10)
	cases := []classTest{
		{test{"No class", true, noModelErr}, tc},
		{test{"Valid update", false, noErr}, updated},
	}

	for _, tc := range cases {
		t.Run(tc.base.name, func(t *testing.T) {
			result := models.Class{}
			if tc.base.isErr {
				_, err := ct.Update(invalidId, tc.class)
				if !strings.Contains(err.Error(), tc.base.expectedErr.Error()) {
					t.Errorf("Errors do not match up")
				}
			} else {
				result, _ = ct.Update(validId, tc.class)
				if !tc.class.Equals(result) {
					t.Errorf("Updated classes do not match up")
				}
			}
		})
	}
}

// TestClassDelete tests that the class model delete function actually deletes the class from the database. Total number of tests: 2.
// It takes in a testing framework object to run tests on.
func TestClassDelete(t *testing.T) {
	// Get the relevant ID for classes
	id := getClassId(updated)
	validId := strconv.FormatInt(id, 10)
	invalidId := strconv.FormatInt(id-1, 10)
	cases := []classTest{
		{test{"No class", true, noModelErr}, tc},
		{test{"Valid delete", false, noErr}, updated},
	}

	for _, tc := range cases {
		t.Run(tc.base.name, func(t *testing.T) {
			if tc.base.isErr {
				err := ct.Delete(invalidId)
				if !strings.Contains(err.Error(), tc.base.expectedErr.Error()) {
					t.Errorf("Errors do not match up")
				}
			} else {
				_ = ct.Delete(validId)
				classes, _ := ct.Get(models.ClassQuery{Id: id}, "")
				if classes != nil {
					t.Errorf("Did not delete class")
				}
			}
		})
	}
}
