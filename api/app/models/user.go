// The models package houses the Model layer within the MVC architecture design.
// This is the lowest layer in the architecture, which directly communicates with
// the database layer. It utilizes the ORM for basic model functionality such as CRUD. The models represent an abstraction to the DB object relations.
package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/dgrijalva/jwt-go"
	"github.com/jcherianucla/UCLA-CS-130/api/config/db"
	"github.com/jcherianucla/UCLA-CS-130/api/utilities"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"time"
)

// The table name as set in the Postgres DB creation.
const (
	USER_TABLE = "users"
)

// Represents the connection to the DB instance.
type UserTable struct {
	connection *db.Db
}

// Represents a User row in the User table within the DB
// Validator and json tags are used for convenient serialization and
// deserialization.
type User struct {
	Id int64 `valid:"-" json:"-"`
	// Distinguishes privileges between a student and professor
	Is_professor bool      `valid:"-" json:"is_professor"`
	Email        string    `valid:"email,required" json:"email"`
	First_name   string    `valid:"required" json:"first_name"`
	Last_name    string    `valid:"required" json:"last_name"`
	Password     []byte    `valid:"required" json:"password"`
	Time_created time.Time `valid:"-" json:"-"`
}

// Represents all fields a user can be queried over.
type UserQuery struct {
	Id           int
	Is_professor bool
	Email        string
	First_name   string
	Last_name    string
}

// NewUser is used to create a new user object from an incoming HTTP request.
// It takes in the HTTP request in JSON format.
// It returns the user constructed and an error if one exists.
func NewUser(r *http.Request) (user User, err error) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = errors.Wrapf(err, "Couldn't read request body")
	}
	utilities.Sugar.Infof("Body: %s", string(b[:]))
	// Converts JSON to user
	json.Unmarshal(b, &user)
	return
}

// GenerateJWT creates a JSON Web Token for a user based on the id,
// with an expiration time of 1 day
// It returns the token string
func (user *User) GenerateJWT() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	utilities.Sugar.Infof("Toke: %v", token)
	tokenString, err := token.SignedString(utilities.SECRET)
	utilities.CheckError(err)
	utilities.Sugar.Infof("Token String: %s", tokenString)
	return tokenString
}

// Equals is a custom comparator for two user objects on non-auto parameter fields.
// It takes in a user object representing the other value.
// It returns a boolean indicating the equality
func (user *User) Equals(other User) bool {
	return user.Is_professor == other.Is_professor &&
		user.Email == other.Email &&
		user.First_name == other.First_name &&
		user.Last_name == other.Last_name &&
		bytes.Equal(user.Password, other.Password)
}

// NewUserTable creates a new table within the database for housing
// all user objects.
// It takes in a reference to an open database connection.
// It returns the constructed table, and an error if one exists.
func NewUserTable(db *db.Db) (userTable UserTable, err error) {
	// Ensure connection is alive
	if db == nil {
		err = errors.New("Invalid database connection")
		return
	}
	userTable.connection = db
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s(
			id SERIAL,
			is_professor BOOLEAN,
			email TEXT,
			first_name TEXT,
			last_name TEXT,
			password BYTEA,
			time_created TIMESTAMP DEFAULT now()
		);`, USER_TABLE)
	// Create the actual table
	if err = userTable.connection.CreateTable(query); err != nil {
		err = errors.Wrapf(err, "User Table creation query failed")
	}
	return
}

// Login will try and retrieve the user based on provided email for the table
// if one exists, and compare the passwords to ensure the same user is logging in.
// It takes in a user object to try and login.
// It return the found user and an error if any exist.
func (table *UserTable) Login(user User) (found User, err error) {
	if !govalidator.IsEmail(user.Email) {
		err = errors.New("Please proved a valid email address")
		return
	} else if len(user.Password) == 0 {
		err = errors.New("Password can't be blank")
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE email=$1", USER_TABLE)

	utilities.Sugar.Infof("SQL Query: %s", query)
	utilities.Sugar.Infof("Value: %s", user.Email)

	stmt, err := table.connection.Pool.Prepare(query)
	if err != nil {
		err = errors.Wrapf(err, "Login query preparation failed")
	}
	row := stmt.QueryRow(user.Email)
	err = row.Scan(&found.Id, &found.Is_professor, &found.Email, &found.First_name, &found.Last_name, &found.Password, &found.Time_created)
	if err != nil {
		err = errors.Wrapf(err, "Failed to find user")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		err = errors.Wrapf(err, "Password hash failed")
		return
	}
	// Compare incoming password with db password
	err = bcrypt.CompareHashAndPassword(hash, found.Password)
	if err != nil {
		err = errors.Wrapf(err, "Provided password does not match")
	}
	return
}

// Insert will put a new user row within the table in the DB, verifying
// all fields are valid.
// It takes in the user object to insert.
// It returns the new user as in the table and an error if one exists.
func (table *UserTable) Insert(user User) (new User, err error) {
	utilities.Sugar.Infof("Password: %v", user.Password)
	data, err := table.connection.Insert(USER_TABLE, "", user, UserQuery{Email: user.Email})
	if err != nil {
		return
	}
	err = utilities.FillStruct(data, &new)
	return
}

// Get attempts to provide a generalized search through the user table based on the
// provided queries.
// It takes a user query for the queryable fields, and an operator such as "AND" or "OR" to
// define the context of the search.
// It returns all the found users and an error if one exists.
func (table *UserTable) Get(userQuery UserQuery, op string) (users []User, err error) {
	allData, err := table.connection.Get(userQuery, op, USER_TABLE)
	if err != nil {
		return
	}
	for _, data := range allData {
		user := User{}
		err = utilities.FillStruct(data, &user)
		if err != nil {
			return
		}
		users = append(users, user)
	}
	return
}

// GetByID uses the internal get mechanism for the table to find a user given an id to search on.
// It takes an ID as a string to convert to an integer to then search on.
// It returns the found user and an error if one exists.
func (table *UserTable) GetByID(strId string) (user User, err error) {
	data, err := table.connection.GetByID(strId, USER_TABLE)
	if err != nil {
		return
	}
	err = utilities.FillStruct(data, &user)
	return
}

// Update will update the user row in the table based on the incoming user object.
// It takes in an id to identify the user in the DB, and updates as a user object.
// It returns the updated user as in the DB, and an error if one exists.
func (table *UserTable) Update(strId string, updates User) (updated User, err error) {
	data, err := table.connection.Update(strId, USER_TABLE, updates)
	if err != nil {
		return
	}
	err = utilities.FillStruct(data, &updated)
	return
}

// Delete permanently removes the user object from the table.
// It takes in an id for the user we wish to delete.
// It returns an error if one exists.
func (table *UserTable) Delete(strId string) error {
	return table.connection.Delete(strId, USER_TABLE)
}

// DeleteAll permanently removes all user objects from the table.
// It returns an error if one exists.
func (table *UserTable) DeleteAll() error {
	return table.connection.DeleteAll(USER_TABLE)
}
