// The models package houses the Model layer within the MVC architecture design.
// This is the lowest layer in the architecture, which directly communicates with
// the database layer. The models represent an abstraction to the DB object relations.
package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/jcherianucla/UCLA-CS-130/api/config/db"
	"github.com/jcherianucla/UCLA-CS-130/api/utilities"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// The table name as set in the Postgres DB creation.
const (
	CLASS_TABLE = "classes"
)

// Represents the connection to the DB instance.
type ClassTable struct {
	connection *db.Db
}

// Represents a Class row in the Class table within the DB
// Validator and json tags are used for convenient serialization and
// deserialization.
type Class struct {
	Id           int       `valid:"-" json:"-"`
	Name         string    `valid:"required" json:"name"`
	Description  string    `valid:"-" json:"description"`
	Quarter      string    `valid:"required" json:"quarter"`
	Year         string    `valid:"required" json:"year"`
	Creator_Id   int       `valid:"-" json:"creator_id"`
	Time_created time.Time `valid:"-" json:"-"`
}

// Represents all fields a class can be queried over.
type ClassQuery struct {
	Id         int
	Name       string
	Quarter    string
	Year       string
	Creator_Id int
}

// NewClass is used to create a new class object from an incoming HTTP request.
// It takes in the HTTP request in JSON format.
// It returns the class constructed and an error if one exists.
func NewClass(r *http.Request) (class Class, err error) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = errors.Wrapf(err, "Couldn't read request body")
	}
	// Converts JSON to class
	json.Unmarshal(b, &class)
	return
}

// Equals is a custom comparator for two class objects on non-auto parameter fields.
// It takes in a class object representing the other value.
// It returns a boolean indicating the equality
func (class *Class) Equals(other Class) bool {
	return class.Name == other.Name &&
		class.Quarter == other.Quarter &&
		class.Year == other.Year &&
		class.Creator_Id == other.Creator_Id
}

// NewClassTable creates a new table within the database for housing
// all class objects.
// It takes in a reference to an open database connection.
// It returns the constructed table, and an error if one exists.
func NewClassTable(db *db.Db) (classTable ClassTable, err error) {
	// Ensure connection is alive
	if db == nil {
		err = errors.New("Invalid database connection")
		return
	}
	classTable.connection = db
	// Create the actual table
	if err = classTable.createTable(); err != nil {
		err = errors.Wrapf(err, "Could not create table on initialization")
	}
	return
}

// createTable runs the actual PSQL query to create the table.
// It returns an error if one exists.
func (table *ClassTable) createTable() (err error) {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s(
			id SERIAL,
			name TEXT,
			description TEXT,
			quarter TEXT,
			year TEXT,
			creator_id int,
			time_created TIMESTAMP DEFAULT now()
		);`, CLASS_TABLE)

	utilities.Sugar.Infof("SQL Query: %s", query)

	// Execute the query
	if _, err = table.connection.Pool.Exec(query); err != nil {
		err = errors.Wrapf(err, "Table creation query failed")
	}
	return
}

// Insert will put a new class row within the table in the DB, verifying
// all fields are valid.
// It takes in the class object to insert and the user id to associate to the creator.
// It returns the new class as in the table and an error if one exists.
func (table *ClassTable) Insert(class Class, userId string) (new Class, err error) {
	_, err = govalidator.ValidateStruct(class)
	if err != nil {
		err = errors.Wrapf(err, "Class model has invalid fields")
		return
	}
	class.Creator_Id, err = strconv.Atoi(userId)
	if err != nil {
		err = errors.Wrapf(err, "Invalid user Id")
		return
	}
	classes, err := table.Get(ClassQuery{Name: class.Name, Quarter: class.Quarter, Year: class.Year, Creator_Id: class.Creator_Id}, "AND")
	if err != nil {
		err = errors.Wrapf(err, "Failed to search for class")
		return
	} else if classes != nil {
		err = errors.New("Class already exists")
		return
	}
	var query bytes.Buffer
	query.WriteString(fmt.Sprintf("INSERT INTO %s (", CLASS_TABLE))
	var values []interface{}
	var vStr, kStr bytes.Buffer
	vIdx := 1
	fields := reflect.ValueOf(class)
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
		values = append(values, v)
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
	// Retrieve class with the new id
	found, err := table.GetByID(newID)
	if err != nil {
		return
	}
	new = found
	return
}

// Get attempts to provide a generalized search through the class table based on the
// provided queries.
// It takes a class query for the queryable fields, and an operator such as "AND" or "OR" to
// define the context of the search.
// It returns all the found classes and an error if one exists.
func (table *ClassTable) Get(classQuery ClassQuery, op string) (classes []Class, err error) {
	var query bytes.Buffer
	query.WriteString(fmt.Sprintf("SELECT * FROM %s WHERE", CLASS_TABLE))
	// Use reflection to analyze object fields
	fields := reflect.ValueOf(classQuery)
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
	// Get all the class rows that matched the query
	for rows.Next() {
		var class Class
		err = rows.Scan(&class.Id, &class.Name, &class.Description, &class.Quarter, &class.Year, &class.Creator_Id, &class.Time_created)
		if err != nil {
			err = errors.Wrapf(err, "Row read failed")
			return
		}
		classes = append(classes, class)
	}
	return
}

// GetByID uses the internal get mechanism for the table to find a class given an id to search on.
// It takes an ID as a string to convert to an integer to then search on.
// It returns the found class and an error if one exists.
func (table *ClassTable) GetByID(strId string) (class Class, err error) {
	id, err := strconv.Atoi(strId)
	if err != nil {
		err = errors.Wrapf(err, "Invalid ID")
		return
	}
	classes, err := table.Get(ClassQuery{Id: id}, "")
	if err != nil {
		err = errors.Wrapf(err, "Search error")
		return
	} else if classes == nil {
		err = errors.New("Failed to find class")
		return
	} else if len(classes) != 1 {
		err = errors.New("Found duplicate classes")
		return
	}
	class = classes[0]
	return
}

// Update will update the class row in the table based on the incoming class object.
// It takes in an id to identify the class in the DB, and updates as a class object.
// It returns the updated class as in the DB, and an error if one exists.
func (table *ClassTable) Update(strId string, updates Class) (updated Class, err error) {
	class, err := table.GetByID(strId)
	if err != nil {
		return
	}
	var query bytes.Buffer
	query.WriteString(fmt.Sprintf("UPDATE %s SET", CLASS_TABLE))
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
		// Skip auto params or unset fields on the incoming Class
		if AUTO_PARAM[k] || utilities.IsUndeclared(fields.Field(i).Interface()) {
			continue
		}
		if first {
			query.WriteString(" ")
			first = false
		} else {
			query.WriteString(", ")
		}
		values = append(values, v)
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
	// Get updated class
	class, err = table.GetByID(strId)
	if err != nil {
		return
	}
	updated = class
	return
}

// Delete permanently removes the class object from the table.
// It takes in an id for the class we wish to delete.
// It returns an error if one exists.
func (table *ClassTable) Delete(strId string) (err error) {
	_, err = table.GetByID(strId)
	if err != nil {
		return
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1;", CLASS_TABLE)

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

// DeleteAll permanently removes all class objects from the table.
// It returns an error if one exists.
func (table *ClassTable) DeleteAll() (err error) {
	query := fmt.Sprintf("DELETE FROM %s;", CLASS_TABLE)
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
