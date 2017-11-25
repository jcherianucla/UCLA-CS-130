// The models package houses the Model layer within the MVC architecture design.
// This is the lowest layer in the architecture, which directly communicates with
// the database layer. The models represent an abstraction to the DB object relations.
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
	"reflect"
	"strconv"
	"strings"
	"time"
)

// The table name as set in the Postgres DB creation.
const (
	USER_TABLE = "users"
)

// Params that should be automatically set by the Postgres DB,
// that is it should never be explicitly set in any code.
var (
	AUTO_PARAM = map[string]bool{
		"id":           true,
		"time_created": true,
	}
)

// Represents the connection to the DB instance.
type UserTable struct {
	connection *db.Db
}

// Represents a User row in the User table within the DB
// Validator and json tags are used for convenient serialization and
// deserialization.
type User struct {
	Id int `valid:"-" json:"-"`
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
	tokenString, err := token.SignedString([]byte(utilities.GP_TOKEN_SECRET))
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
	// Create the actual table
	if err = userTable.createTable(); err != nil {
		err = errors.Wrapf(err, "Could not create table on initialization")
	}
	return
}

// createTable runs the actual PSQL query to create the table.
// It returns an error if one exists.
func (table *UserTable) createTable() (err error) {
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

	utilities.Sugar.Infof("SQL Query: %s", query)

	// Execute the query
	if _, err = table.connection.Pool.Exec(query); err != nil {
		err = errors.Wrapf(err, "Table creation query failed")
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
	_, err = govalidator.ValidateStruct(user)
	if err != nil {
		err = errors.Wrapf(err, "User model has invalid fields")
		return
	}
	users, err := table.Get(UserQuery{Email: user.Email}, "")
	if err != nil {
		err = errors.Wrapf(err, "Failed to search for user")
		return
	} else if users != nil {
		err = errors.New("User already exists")
		return
	}
	var query bytes.Buffer
	query.WriteString(fmt.Sprintf("INSERT INTO %s (", USER_TABLE))
	var values []interface{}
	var vStr, kStr bytes.Buffer
	vIdx := 1
	fields := reflect.ValueOf(user)
	if fields.NumField() < 1 {
		err = errors.New("Invalid number of fields given")
		return
	}
	first := true
	for i := 0; i < fields.NumField(); i++ {
		k := strings.ToLower(fields.Type().Field(i).Name)
		v := fields.Field(i).Interface()
		// Skip auto params
		if AUTO_PARAM[k] {
			continue
		}
		if first {
			first = false
		} else {
			vStr.WriteString(", ")
			kStr.WriteString(", ")
		}
		kStr.WriteString(k)
		vStr.WriteString(fmt.Sprintf("$%d", vIdx))
		// Hash the password
		if k == "Password" {
			vStr, _ := v.(string)
			var hash []byte
			hash, err = bcrypt.GenerateFromPassword([]byte(vStr), bcrypt.DefaultCost)
			if err != nil {
				err = errors.Wrapf(err, "Password hash failed")
				return
			}
			values = append(values, hash)
		} else {
			values = append(values, v)
		}
		vIdx++
	}
	query.WriteString(fmt.Sprintf("%s) VALUES (%s) RETURNING id;", kStr.String(), vStr.String()))

	utilities.Sugar.Infof("SQL Query: %s", query.String())
	utilities.Sugar.Infof("Values: %v", values)

	stmt, err := table.connection.Pool.Prepare(query.String())
	if err != nil {
		err = errors.Wrapf(err, "Insertion query preparation failed")
		return
	}
	err = stmt.QueryRow(values...).Scan(&new.Id)
	if err != nil {
		err = errors.Wrapf(err, "Insertion query failed to execute")
		return
	}
	newID := strconv.Itoa(new.Id)
	// Retrieve user with the new id
	found, err := table.GetByID(newID)
	if err != nil {
		return
	}
	new = found
	return
}

// Get attempts to provide a generalized search through the user table based on the
// provided queries.
// It takes a user query for the queryable fields, and an operator such as "AND" or "OR" to
// define the context of the search.
// It returns all the found users and an error if one exists.
func (table *UserTable) Get(userQuery UserQuery, op string) (users []User, err error) {
	var query bytes.Buffer
	query.WriteString(fmt.Sprintf("SELECT * FROM %s WHERE", USER_TABLE))
	// Use reflection to analyze object fields
	fields := reflect.ValueOf(userQuery)
	first := true
	var values []interface{}
	vIdx := 1
	for i := 0; i < fields.NumField(); i++ {
		if first {
			query.WriteString(" ")
			first = false
		} else {
			if op != "" {
				query.WriteString(fmt.Sprintf(" %s ", op))
			}
		}
		v := fields.Field(i).Interface()
		// Skip fields that are not set to query on
		if !utilities.IsUndeclared(v) {
			k := strings.ToLower(fields.Type().Field(i).Name)
			v = fmt.Sprintf("%v", v)
			values = append(values, v)
			query.WriteString(fmt.Sprintf("%s=$%d", k, vIdx))
			vIdx++
		}
	}
	query.WriteString(";")

	utilities.Sugar.Infof("SQL Query: %s", query.String())
	utilities.Sugar.Infof("Values: %v", values)

	stmt, err := table.connection.Pool.Prepare(query.String())
	if err != nil {
		err = errors.Wrapf(err, "Get query preparation failed")
		return
	}
	rows, err := stmt.Query(values...)
	if err != nil {
		err = errors.Wrapf(err, "Get query failed to execute")
		return
	}
	// Get all the user rows that matched the query
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Is_professor, &user.Email, &user.First_name, &user.Last_name, &user.Password, &user.Time_created)
		if err != nil {
			err = errors.Wrapf(err, "Row read failed")
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
	id, err := strconv.Atoi(strId)
	if err != nil {
		err = errors.Wrapf(err, "Invalid ID")
		return
	}
	users, err := table.Get(UserQuery{Id: id}, "")
	if err != nil {
		err = errors.Wrapf(err, "Search error")
		return
	} else if users == nil {
		err = errors.New("Failed to find user")
		return
	} else if len(users) != 1 {
		err = errors.New("Found duplicate users")
		return
	}
	user = users[0]
	return
}

// Update will update the user row in the table based on the incoming user object.
// It takes in an id to identify the user in the DB, and updates as a user object.
// It returns the updated user as in the DB, and an error if one exists.
func (table *UserTable) Update(strId string, updates User) (updated User, err error) {
	user, err := table.GetByID(strId)
	if err != nil {
		return
	}
	var query bytes.Buffer
	query.WriteString(fmt.Sprintf("UPDATE %s SET", USER_TABLE))
	var values []interface{}
	vIdx := 1
	fields := reflect.ValueOf(updates)
	if fields.NumField() < 1 {
		err = errors.New("Invalid number of query fields")
		return
	}
	first := true
	for i := 0; i < fields.NumField(); i++ {
		k := strings.ToLower(fields.Type().Field(i).Name)
		v := fields.Field(i).Interface()
		// Skip auto params or unset fields on the incoming User
		if AUTO_PARAM[k] || utilities.IsUndeclared(fields.Field(i).Interface()) {
			continue
		}
		if first {
			query.WriteString(" ")
			first = false
		} else {
			query.WriteString(", ")
		}
		// Hash new password
		if k == "Password" {
			vStr := v.(string)
			hash, cryptErr := bcrypt.GenerateFromPassword([]byte(vStr), bcrypt.DefaultCost)
			if cryptErr != nil {
				err = errors.Wrapf(cryptErr, "Password hash failed")
				return
			}
			values = append(values, hash)
		} else {
			values = append(values, v)
		}
		query.WriteString(fmt.Sprintf("%v=$%d", k, vIdx))
		vIdx += 1
	}
	query.WriteString(fmt.Sprintf(" WHERE id=%s;", strId))

	utilities.Sugar.Infof("SQL Query: %s", query.String())
	utilities.Sugar.Infof("Values: %v", values)

	stmt, err := table.connection.Pool.Prepare(query.String())
	if err != nil {
		err = errors.Wrapf(err, "Update query preparation failed")
		return
	}
	if _, err = stmt.Exec(values...); err != nil {
		err = errors.Wrapf(err, "Update query failed to execute")
		return
	}
	// Get updated user
	user, err = table.GetByID(strId)
	if err != nil {
		return
	}
	updated = user
	return
}

// Delete permanently removes the user object from the table.
// It takes in an id for the user we wish to delete.
// It returns an error if one exists.
func (table *UserTable) Delete(strId string) (err error) {
	_, err = table.GetByID(strId)
	if err != nil {
		return
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1;", USER_TABLE)

	utilities.Sugar.Infof("SQL Query: %s", query)

	stmt, err := table.connection.Pool.Prepare(query)
	if err != nil {
		err = errors.Wrapf(err, "Delete query preparation failed")
		return
	}

	if _, err = stmt.Exec(strId); err != nil {
		err = errors.Wrapf(err, "Delete query failed to execute")
	}
	return
}

// DeleteAll permanently removes all user objects from the table.
// It returns an error if one exists.
func (table *UserTable) DeleteAll() (err error) {
	query := fmt.Sprintf("DELETE FROM %s;", USER_TABLE)
	utilities.Sugar.Infof("SQL Query: %s", query)

	stmt, err := table.connection.Pool.Prepare(query)
	if err != nil {
		err = errors.Wrapf(err, "Delete all query preparation failed")
		return
	}

	if _, err = stmt.Exec(); err != nil {
		err = errors.Wrapf(err, "Delete all query failed to execute")
	}
	return
}
