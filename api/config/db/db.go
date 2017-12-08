// db package houses the very lowest layer in the overall application
// architecture that communicates with the PostgreSQL database pool.
// This connection is used upstream by the model layer. The package also provides a minimal ORM for basic CRUD operations on the database.
package db

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/jcherianucla/UCLA-CS-130/api/utilities"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"reflect"
	"strconv"
	"strings"
)

// Params that should be automatically set by the Postgres DB,
// that is it should never be explicitly set in any code.
var (
	AUTO_PARAM = map[string]bool{
		"id":           true,
		"time_created": true,
	}
)

// Represents the database pool with the DB specific
// configurations.
type Db struct {
	Pool *sql.DB
	cfg  Config
}

type idQuery struct {
	Id int
}

// Represents the configuration parameters needed to
// open up a connection to the database.
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

// Close ends the connection to the specific database.
// It returns the error if one exists.
func (db *Db) Close() (err error) {
	if db.Pool == nil {
		return
	}
	if err = db.Pool.Close(); err != nil {
		err = errors.Wrapf(err, "Could not close postgres db")
	}
	return
}

// New creates a new connection to the existing database on the
// host system.
// It takes in a configuration to open the specific database.
// It returns the open database and an error if one exists.
func New(cfg Config) (db Db, err error) {
	if cfg.Host == "" || cfg.Port == "" || cfg.User == "" ||
		cfg.Password == "" || cfg.Database == "" {
		err = errors.New("Provide all fields for config")
		return
	}
	db.cfg = cfg

	pqDb, err := sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.User, cfg.Password, cfg.Database, cfg.Host, cfg.Port))
	if err != nil {
		err = errors.Wrapf(err, "Failed to open a connection to the database")
		return
	}
	// Ping the connection to ensure its open
	if err = pqDb.Ping(); err != nil {
		err = errors.Wrapf(err, "Unable to ping database")
		return
	}
	db.Pool = pqDb
	return
}

// CreateTable runs the actual PSQL query to create the table.
// It takes the specific creation PSQL query to run, as this varies with each object.
// It returns an error if one exists.
func (db *Db) CreateTable(query string) (err error) {
	utilities.Sugar.Infof("SQL Query: %s", query)

	// Execute the query
	if _, err = db.Pool.Exec(query); err != nil {
		err = errors.Wrapf(err, "Classes Table creation query failed")
	}
	return
}

// Insert will put a new model row within the specified table in the DB, verifying all fields are valid.
// It takes in the object to insert, the operator to perform the query with, the table name and the query itself.
// It returns the new object as in the table and an error if one exists.
func (db *Db) Insert(table, op string, model, mQuery interface{}) (new map[string]interface{}, err error) {
	modelName := reflect.TypeOf(model).String()
	_, err = govalidator.ValidateStruct(model)
	if err != nil {
		err = errors.Wrapf(err, "%s model has invalid fields", modelName)
		return
	}
	objs, err := db.Get(mQuery, op, table)
	if err != nil {
		err = errors.Wrapf(err, "Failed to search for %s", modelName)
		return
	} else if objs != nil {
		err = errors.New(fmt.Sprintf("%s already exists", modelName))
		return
	}
	var query bytes.Buffer
	query.WriteString(fmt.Sprintf("INSERT INTO %s (", table))
	var values []interface{}
	var vStr, kStr bytes.Buffer
	vIdx := 1
	fields := reflect.ValueOf(model)
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
		if k == "password" {
			pass, _ := v.(string)
			var hash []byte
			hash, err = bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
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

	stmt, err := db.Pool.Prepare(query.String())
	if err != nil {
		err = errors.Wrapf(err, "Insertion query preparation failed")
		return
	}
	id := 0
	err = stmt.QueryRow(values...).Scan(&id)
	if err != nil {
		err = errors.Wrapf(err, "Insertion query failed to execute")
		return
	}
	newID := strconv.Itoa(id)
	// Retrieve object with the new id
	found, err := db.GetByID(newID, table)
	if err != nil {
		return
	}
	new = found
	return
}

// Get attempts to provide a generalized search through the specified table based on the provided queries.
// It takes a query for the queryable fields, and an operator such as "AND" or "OR" to define the context of the search. It takes in a table name to act on.
// It returns all the data for all found objects and an error if one exists.
func (db *Db) Get(mQuery interface{}, op, table string) (objects []map[string]interface{}, err error) {
	var query bytes.Buffer
	query.WriteString(fmt.Sprintf("SELECT * FROM %s WHERE", table))
	// Use reflection to analyze object fields
	fields := reflect.ValueOf(mQuery)
	first := true
	var values []interface{}
	vIdx := 1
	for i := 0; i < fields.NumField(); i++ {
		v := fields.Field(i).Interface()
		// Skip fields that are not set to query on
		if !utilities.IsUndeclared(v) {
			if first {
				query.WriteString(" ")
				first = false
			} else {
				if op != "" {
					query.WriteString(fmt.Sprintf(" %s ", op))
				}
			}
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

	stmt, err := db.Pool.Prepare(query.String())
	if err != nil {
		err = errors.Wrapf(err, "Get query preparation failed")
		return
	}
	rows, err := stmt.Query(values...)
	if err != nil {
		err = errors.Wrapf(err, "Get query failed to execute")
		return
	}
	// Get all the rows that matched the query
	cols, _ := rows.Columns()
	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPtrs := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPtrs[i] = &columns[i]
		}
		if err = rows.Scan(columnPtrs...); err != nil {
			err = errors.Wrapf(err, "Get query failed to execute")
			return
		}
		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPtrs[i].(*interface{})
			m[strings.Title(colName)] = *val
		}
		objects = append(objects, m)
	}
	return
}

// GetByID uses the internal get mechanism for the table to find an object given an id to search on.
// It takes an ID as a string to convert to an integer to then search on, and a string representing the table name.
// It returns the data for the found object, or an error if one exists.
func (db *Db) GetByID(strId, table string) (data map[string]interface{}, err error) {
	id, err := strconv.Atoi(strId)
	if err != nil {
		err = errors.Wrapf(err, "Invalid ID")
		return
	}
	query := idQuery{Id: id}
	objs, err := db.Get(query, "", table)
	if err != nil {
		err = errors.Wrapf(err, "Search error")
		return
	} else if objs == nil {
		err = errors.New("Failed to find object")
		return
	} else if len(objs) != 1 {
		err = errors.New("Found duplicate objects")
		return
	}
	data = objs[0]
	return
}

// Update will update the model row in the table based on the incoming object.
// It takes in an id to identify the object in the DB, a string representing the table name and the fileds to update the object on.
// It returns the data representing an updated model.
func (db *Db) Update(strId, table string, updates interface{}) (data map[string]interface{}, err error) {
	found, err := db.GetByID(strId, table)
	if err != nil {
		return
	}
	var query bytes.Buffer
	query.WriteString(fmt.Sprintf("UPDATE %s SET", table))
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

	stmt, err := db.Pool.Prepare(query.String())
	if err != nil {
		err = errors.Wrapf(err, "Update query preparation failed")
		return
	}
	if _, err = stmt.Exec(values...); err != nil {
		err = errors.Wrapf(err, "Update query failed to execute")
		return
	}
	// Get updated model
	found, err = db.GetByID(strId, table)
	if err != nil {
		return
	}
	data = found
	return
}

// Delete permanently removes the object from the table.
// It takes in an id for the user we wish to delete, and a table name to act on.
// It returns an error if one exists.
func (db *Db) Delete(strId, table string) (err error) {
	_, err = db.GetByID(strId, table)
	if err != nil {
		return
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1;", table)

	utilities.Sugar.Infof("SQL Query: %s", query)

	stmt, err := db.Pool.Prepare(query)
	if err != nil {
		err = errors.Wrapf(err, "Delete query preparation failed")
		return
	}

	if _, err = stmt.Exec(strId); err != nil {
		err = errors.Wrapf(err, "Delete query failed to execute")
	}
	return
}

// DeleteAll permanently removes all objects from the table.
// It takes in a string representing the table name.
// It returns an error if one exists.
func (db *Db) DeleteAll(table string) (err error) {
	query := fmt.Sprintf("DELETE FROM %s;", table)
	utilities.Sugar.Infof("SQL Query: %s", query)

	stmt, err := db.Pool.Prepare(query)
	if err != nil {
		err = errors.Wrapf(err, "Delete all query preparation failed")
		return
	}

	if _, err = stmt.Exec(); err != nil {
		err = errors.Wrapf(err, "Delete all query failed to execute")
	}
	return
}
