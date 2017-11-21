// The tests package houses all the tests for all aspects of the backend application.
// It focuses on running negative tests to ensure the system before correctly in situations of error. It also runs basic positive tests to make sure the right data is acquired.
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
	fieldErr        = errors.New("User model has invalid fields")
	emailErr        = errors.New("User model has invalid fields: email")
	emailErr2       = errors.New("Please proved a valid email address")
	existsErr       = errors.New("User already exists")
	passErr         = errors.New("Provided password does not match")
	noUserErr       = errors.New("Failed to find user")
	invalidQueryErr = errors.New("Get query preparation failed")
)

// Test user model to be used throughout
var tu models.User = models.User{
	Is_professor: false,
	Email:        "jcherian@ucla.edu",
	First_name:   "Jahan",
	Last_name:    "Cherian",
	Password:     []byte("password"),
}

// Test updated user model to be used in updates
var updated models.User = models.User{
	Is_professor: true,
	Email:        "jcherian@ucla.edu",
	First_name:   "Jahan",
	Last_name:    "Kuruvilla Cherian",
	Password:     []byte("newpassword"),
}

// The user table reference to act on
var ut models.UserTable

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
	// Create a user table
	ut, _ = models.NewUserTable(&db)
}

// Represents a basic test object that all tests must conform to.
type test struct {
	name        string
	isErr       bool
	expectedErr error
}

// Represents a test on a user object.
type userTest struct {
	base test
	user models.User
}

// Represents a test on a query object.
type queryTest struct {
	base  test
	query models.UserQuery
}

// getId returns the corresponding ID for a user.
// It takes in a user object.
// It returns the corresponding int ID.
func getId(user models.User) int {
	users, _ := ut.Get(models.UserQuery{Email: user.Email}, "")
	return users[0].Id
}

// TestUserInsert will test the insert method for user models amongst multiple scenarios. It uses table based testing. Total number of tests: 4.
// It takes in a testing framework object to run tests on.
func TestUserInsert(t *testing.T) {
	cases := []userTest{
		{test{"Missing fields", true, fieldErr}, models.User{Email: tu.Email, First_name: tu.First_name, Last_name: tu.Last_name}},
		{test{"Invalid email", true, emailErr}, models.User{Email: "jcherianulca.edu", Is_professor: tu.Is_professor, First_name: tu.First_name, Last_name: tu.Last_name, Password: tu.Password}},
		{test{"Valid user", false, noErr}, tu},
		{test{"Existing user", true, existsErr}, tu},
	}

	// Clean
	_ = ut.DeleteAll()
	for _, tc := range cases {
		t.Run(tc.base.name, func(t *testing.T) {
			result := models.User{}
			result, err := ut.Insert(tc.user)
			if tc.base.isErr {
				if !strings.Contains(err.Error(), tc.base.expectedErr.Error()) {
					t.Errorf("Errors do not match up")
				}
			} else {
				if !tc.user.Equals(result) {
					t.Errorf("Users do not match up")
				}
			}
		})
	}
}

// TestUserLogin will test the login method for the user models. It relies on the database having the testUser already inserted. Total number of tests: 4
// It takes in a testing framework object to run tests on.
func TestUserLogin(t *testing.T) {
	cases := []userTest{
		{test{"Invalid email", true, emailErr2}, models.User{Email: "jcherianucla.edu", Password: tu.Password}},
		{test{"Password equality", true, passErr}, models.User{Email: tu.Email, Password: []byte("passwor")}},
		{test{"No user", true, noUserErr}, models.User{Email: "oozgur217@gmail.com", Password: tu.Password}},
		{test{"Valid Login", false, noErr}, models.User{Email: tu.Email, Password: tu.Password}},
	}
	for _, tc := range cases {
		t.Run(tc.base.name, func(t *testing.T) {
			result := models.User{}
			result, err := ut.Login(tc.user)
			if tc.base.isErr {
				if !strings.Contains(err.Error(), tc.base.expectedErr.Error()) {
					t.Errorf("Errors do not match up")
				}
			} else {
				if !tu.Equals(result) {
					t.Errorf("Users do not match up")
				}
			}
		})
	}
}

// TestUserGet simply tests query validation. This is because for the most part, it used internally with error handling done in other parts of the model functionality. Total number of tests: 1.
// It takes in a testing framework object to run tests on.
func TestUserGet(t *testing.T) {
	cases := []queryTest{
		{test{"Invalid query", true, invalidQueryErr}, models.UserQuery{}},
	}

	for _, tc := range cases {
		t.Run(tc.base.name, func(t *testing.T) {
			_, err := ut.Get(tc.query, "")
			if tc.base.isErr {
				if !strings.Contains(err.Error(), tc.base.expectedErr.Error()) {
					t.Errorf("Errors do not match up")
				}
			}
		})
	}
}

// TestUserUpdate tests the update method for the user model. If the update works on any field, then the overall update method works as it is the same logic for each field. Total number of tests: 2.
// It takes in a testing framework object to run tests on.
func TestUserUpdate(t *testing.T) {
	// Get the relevant ID for user
	id := getId(tu)
	validId := strconv.Itoa(id)
	invalidId := strconv.Itoa(id - 1)
	cases := []userTest{
		{test{"No user", true, noUserErr}, tu},
		{test{"Valid update", false, noErr}, updated},
	}

	for _, tc := range cases {
		t.Run(tc.base.name, func(t *testing.T) {
			result := models.User{}
			if tc.base.isErr {
				_, err := ut.Update(invalidId, tc.user)
				if !strings.Contains(err.Error(), tc.base.expectedErr.Error()) {
					t.Errorf("Errors do not match up")
				}
			} else {
				result, _ = ut.Update(validId, tc.user)
				if !tc.user.Equals(result) {
					t.Errorf("Updated users do not match up")
				}
			}
		})
	}
}

// TestUserDelete tests that the user model delete function actually deletes the user from the database. Total number of tests: 2.
// It takes in a testing framework object to run tests on.
func TestUserDelete(t *testing.T) {
	// Get the relevant ID for user
	id := getId(updated)
	validId := strconv.Itoa(id)
	invalidId := strconv.Itoa(id - 1)
	cases := []userTest{
		{test{"No user", true, noUserErr}, tu},
		{test{"Valid delete", false, noErr}, updated},
	}

	for _, tc := range cases {
		t.Run(tc.base.name, func(t *testing.T) {
			if tc.base.isErr {
				err := ut.Delete(invalidId)
				if !strings.Contains(err.Error(), tc.base.expectedErr.Error()) {
					t.Errorf("Errors do not match up")
				}
			} else {
				_ = ut.Delete(validId)
				users, _ := ut.Get(models.UserQuery{Id: id}, "")
				if users != nil {
					t.Errorf("Did not delete user")
				}
			}
		})
	}
}
