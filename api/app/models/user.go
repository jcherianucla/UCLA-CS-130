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

const (
	TABLE = "users"
)

var (
	// These are auto maintained by the DB
	AUTO_PARAM = map[string]bool{
		"id":           true,
		"time_created": true,
	}
)

// Holds the connection to the db instance
type UserTable struct {
	connection *db.Db
}

// Basic user model
type User struct {
	Id           int       `valid:"-" json:"-"`
	Is_professor bool      `valid:"-" json:"is_professor"`
	Email        string    `valid:"email,required" json:"email"`
	First_name   string    `valid:"required" json:"first_name"`
	Last_name    string    `valid:"required" json:"last_name"`
	Password     []byte    `valid:"required" json:"password"`
	Time_created time.Time `valid:"-" json:"-"`
}

// Queryable over the following
type UserQuery struct {
	Id           int
	Is_professor bool
	Email        string
	First_name   string
	Last_name    string
}

// Creates a new user from a request
func NewUser(r *http.Request) (user User, err error) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = errors.Wrapf(err, "Couldn't read request body")
	}
	// Converts JSON to user
	json.Unmarshal(b, &user)
	// Hashes password
	user.Password, err = bcrypt.GenerateFromPassword(user.Password, bcrypt.DefaultCost)
	if err != nil {
		err = errors.Wrapf(err, "Error hashing password")
	}
	return
}

func (user *User) GenerateJWT() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, _ := token.SignedString(utilities.GP_TOKEN_SECRET)
	return tokenString
}

// Creates a new user table
func NewUserTable(db *db.Db) (userTable UserTable, err error) {
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

// Runs the SQL to create the table within the DB
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
		);`, TABLE)
	if _, err = table.connection.Pool.Exec(query); err != nil {
		err = errors.Wrapf(err, "Table creation query failed")
	}
	return
}

// Get by unique email
func (table *UserTable) GetUserByEmail(email string) (user User, err error) {
	users, err := table.GetUser(UserQuery{Email: email}, "")
	if len(users) > 0 {
		user = users[0]
	}
	return
}

// Can get all users based on any queryable fields and the varying operation to combine fields
func (table *UserTable) GetUser(userQuery UserQuery, op string) (users []User, err error) {
	var query bytes.Buffer
	query.WriteString(fmt.Sprintf("SELECT * FROM %s WHERE", TABLE))
	fields := reflect.ValueOf(userQuery)
	if fields.NumField() <= 0 {
		err = errors.New("Invalid number of query fields")
		return
	}
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
		// Skip fields that are not set to query on
		if !utilities.IsUndeclared(fields.Field(i)) {
			k := strings.ToLower(fields.Type().Field(i).Name)
			v := fmt.Sprintf("%v", fields.Field(i).Interface())
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
		err = errors.Wrapf(err, "Could not prepare db query for user")
		return
	}
	rows, err := stmt.Query(values...)
	if err != nil {
		err = errors.Wrapf(err, "Could not retrieve rows for query")
		return
	}
	// Get all the user rows that matched the query
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Is_professor, &user.Email, &user.First_name, &user.Last_name, &user.Password, &user.Time_created)
		if err != nil {
			err = errors.Wrapf(err, "Failed to scan row")
			return
		}
		users = append(users, user)
	}
	return
}

// Inserts a new user into the DB
func (table *UserTable) InsertUser(user User) (new User, err error) {
	_, err = govalidator.ValidateStruct(user)
	if err != nil {
		err = errors.Wrapf(err, "Invalid user model")
		return
	}
	users, err := table.GetUser(UserQuery{Email: user.Email}, "AND")
	if err != nil {
		err = errors.Wrapf(err, "Search error")
		return
	} else if users != nil {
		err = errors.New("Found duplicate user")
		return
	}
	var query bytes.Buffer
	query.WriteString(fmt.Sprintf("INSERT INTO %s (", TABLE))
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
		v := fmt.Sprintf("%v", fields.Field(i).Interface())
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
		values = append(values, v)
		vIdx++
	}
	query.WriteString(fmt.Sprintf("%s) VALUES (%s) RETURNING id;", kStr.String(), vStr.String()))

	utilities.Sugar.Infof("SQL Query: %s", query.String())
	utilities.Sugar.Infof("Values: %v", values)

	stmt, err := table.connection.Pool.Prepare(query.String())
	if err != nil {
		err = errors.Wrapf(err, "Could not prepare insertion query")
		return
	}
	err = stmt.QueryRow(values...).Scan(&new.Id)
	if err != nil {
		err = errors.Wrapf(err, "Could not insert user")
		return
	}
	// Retrieve user with the new id
	users, err = table.GetUser(UserQuery{Id: new.Id}, "")
	if err != nil {
		err = errors.Wrapf(err, "Failed to get user")
	} else if len(users) != 1 {
		err = errors.New("Found duplicate users while inserting")
	} else {
		new = users[0]
	}
	return
}

func (table *UserTable) userExists(userQuery UserQuery) (err error) {
	users, err := table.GetUser(userQuery, "AND")
	// Check if user exists
	if err != nil {
		err = errors.Wrapf(err, "Error while searching for user")
		return
	} else if users != nil {
		return
	}
	return
}

func (table *UserTable) UpdateUser(id int, updates User) (updated User, err error) {
	users, err := table.GetUser(UserQuery{Id: id}, "AND")
	if err != nil {
		err = errors.Wrapf(err, "Search error")
		return
	} else if users == nil {
		err = errors.New("Couldn't find user")
		return
	}
	var query bytes.Buffer
	query.WriteString(fmt.Sprintf("UPDATE %s SET", TABLE))
	var values []interface{}
	vIdx := 1
	fields := reflect.ValueOf(updates)
	if fields.NumField() < 1 {
		err = errors.New("Invalid number of fields given")
		return
	}
	first := true
	for i := 0; i < fields.NumField(); i++ {
		k := strings.ToLower(fields.Type().Field(i).Name)
		v := fmt.Sprintf("%v", fields.Field(i).Interface())
		// Skip auto params or unset fields
		if AUTO_PARAM[k] || utilities.IsUndeclared(fields.Field(i)) {
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
			hash, cryptErr := bcrypt.GenerateFromPassword([]byte(v), bcrypt.DefaultCost)
			if cryptErr != nil {
				err = errors.Wrapf(cryptErr, "Problem creating password hash")
				return
			}
			values = append(values, hash)
		} else {
			values = append(values, v)
		}
		query.WriteString(fmt.Sprintf("%v=$%d", k, vIdx))
		vIdx += 1
	}
	// End of query
	query.WriteString(fmt.Sprintf(" WHERE id=%d;", id))

	stmt, err := table.connection.Pool.Prepare(query.String())
	if err != nil {
		err = errors.Wrapf(err, "Couldn't prepare update query")
		return
	}
	if _, err = stmt.Exec(values...); err != nil {
		err = errors.Wrapf(err, "Couldn't execute update query")
		return
	}
	// Get updated user
	users, err = table.GetUser(UserQuery{Id: id}, "")
	if err != nil {
		err = errors.Wrapf(err, "Error getting updated user")
		return
	} else if len(users) != 1 {
		err = errors.New("Found duplicate users")
		return
	}
	updated = users[0]
	return
}

func (table *UserTable) DeleteUser(id int) (err error) {
	users, err := table.GetUser(UserQuery{Id: id}, "AND")
	if err != nil {
		err = errors.Wrapf(err, "Search error")
		return
	} else if users == nil {
		err = errors.New("Couldn't find user")
		return
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1;", TABLE)

	stmt, err := table.connection.Pool.Prepare(query)
	if err != nil {
		err = errors.Wrapf(err, "Couldn't prepare delete query")
		return
	}

	if _, err = stmt.Exec(strconv.Itoa(id)); err != nil {
		err = errors.Wrapf(err, "Couldn't execute delete query")
	}
	return
}
