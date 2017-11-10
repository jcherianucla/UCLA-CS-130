package models

import (
	"bytes"
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/jcherianucla/UCLA-CS-130/api/config"
	"github.com/jcherianucla/UCLA-CS-130/api/utilities"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/oleiade/reflections.v1"
	"time"
)

const (
	TABLE = "users"
)

var (
	// These are auto maintained by the DB
	AUTO_PARAM = map[string]bool{
		"Id":         true,
		"Created_at": true,
	}
)

// Holds the connection to the db instance
type UserTable struct {
	connection *Db
}

// Basic user model
type User struct {
	Id           int       `valid:"-" json:"-"`
	Is_professor bool      `valid:"-" json:"is_professor"`
	Email        string    `valid:"email,required" json:"email"`
	First_name   string    `valid:"first_name,required" json:"first_name"`
	Last_name    string    `valid:"last_name,required" json:"last_name"`
	Password     []byte    `valid:"required" json:"password"`
	Created_at   time.Time `valid:"-" json:"-"`
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
	user.Password = bcrypt.GenerateFromPassword(user.Password, bcrypt.DefaultCost())
	return
}

// Creates a new user table
func NewUserTable(db *Db) (userTable UserTable, err error) {
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
			password BYTEA
			time_created TIMESTAMP DEFAULT now()
		);`, TABLE)
	if _, err = table.connection.Exec(query); err != nil {
		err = errors.Wrapf(err, "Table creation query failed")
	}
	return
}

// Get by unique email
func (table *UserTable) GetUserByEmail(email string) (user User, err error) {
	userQuery := UserQuery{email: email}
	users, err := table.GetUser(UserQuery, "")
	if len(users) > 0 {
		user = users[0]
	}
	return
}

// Can get all users based on any queryable fields and the varying operation to combine fields
func (table *UserTable) GetUser(query UserQuery, op string) (users []User, err error) {
	var query bytes.Buffer
	query.WriteString(fmt.Sprintf("SELECT * FROM %s WHERE", TABLE))
	fields := reflect.ValueOf(query)
	if fields.NumField() <= 0 {
		err = errors.New("Invalid number of query fields")
		return
	}
	first := true
	var values []string
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
		if !utilities.IsUndeclared(reflect.Field(i)) {
			k := reflect.Field(i).Name()
			v := fmt.Sprintf("%v", reflect.Field(i).Interface())
			values = append(values, v)
			query.WriteString(fmt.Sprintf("%s=$%d", k, vIdx))
			vIdx++
		}
	}
	query.WriteString(";")

	stmt, err := table.connection.Prepare(query.String())
	if err != nil {
		err = errors.Wrapf(err, "Could not prepare db query for user")
		return
	}
	rows, err := table.connection.Query(values...)
	if err != nil {
		err = errors.Wrapf(err, "Could not retrieve rows for query")
		return
	}
	// Get all the user rows that matched the query
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Is_professor, &user.Email, &user.First_name, &user.Last_name, &user.Password, &user.Created_at)
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
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		err = errors.Wrapf(err, "Invalid user model")
		return
	}
	// Make sure user doesn't already exist
	exist, err := table.GetUserByEmail(user.Email)
	if err != nil {
		err = errors.Wrapf(err, "Failed to check for existing user")
		return
	} else if exist != nil {
		err = errors.New(fmt.Sprintf("User with email: %s already exists", user.Email))
		return
	}
	var query bytes.Buffer
	query.WriteString(fmt.Sprintf("INSERT INTO %s (", TABLE))
	var values []string
	var vStr, kStr bytes.Buffer
	vIdx := 1
	fields := reflect.ValueOf(user)
	first := true
	for i := 0; i < fields.NumField(); i++ {
		k := reflect.Field(i).Name()
		v := fmt.Sprintf("%v", reflect.Field(i).Interface())
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
	query.WriteString(fmt.Sprintf("%s) VALUES (%s) RETURNING id;", kStr, vStr))
	stmt, err := table.connection.Prepare(query)
	if err != nil {
		err = errors.Wrapf(err, "Could not prepare insertion query")
		return
	}
	err := table.connection.QueryRow(values...).Scan(&new.Id)
	if err != nil {
		err = errors.Wrapf(err, "Could not insert user")
		return
	}
	// Retrieve user with the new id
	new = table.GetUser(UserQuery{id: &new.Id})
	return
}
