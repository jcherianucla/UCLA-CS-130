// db package houses the very lowest layer in the overall application
// architecture that communicates with the PostgreSQL database pool.
// This connection is used upstream by the model layer.
package db

import (
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

func (db *Db) Insert(table, op string, model, mQuery interface{}) (new map[string]interface{}, err error) {
	modelName := reflect.TypeOf(model).String()
	_, err = govalidator.ValidateStruct(model)
	if err != nil {
		err = errors.Wrapf(err, "%s model has invalid fields", modelName)
		return
	}
	objs, err := db.Get(mQuery, op)
	if err != nil {
		err = errors.Wrapf(err, "Failed to search for %s", modelName)
		return
	} else if users != nil {
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

	stmt, err := db.Pool.Prepare(query.String())
	if err != nil {
		err = errors.Wrapf(err, "Insertion query preparation failed")
		return
	}
	var id int
	err = stmt.QueryRow(values...).Scan(id)
	if err != nil {
		err = errors.Wrapf(err, "Insertion query failed to execute")
		return
	}
	newID := strconv.Itoa(id)
	// Retrieve user with the new id
	found, err := db.GetByID(newID)
	if err != nil {
		return
	}
	new = found
	return
}

func (db *Db) Get(mQuery interface{}, op, table string) (objects []map[string]interface{}, err error) {
	var query bytes.Buffer
	query.WriteString(fmt.Sprintf("SELECT * FROM %s WHERE", table))
	// Use reflection to analyze object fields
	fields := reflect.ValueOf(mQuery)
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
			m[colName] = *val
		}
		objects = append(objects, m)
	}
	return
}

func (orm *ORM) GetByID(strId string) (obj map[string]interface{}, err error) {
	id, err := strconv.Atoi(strId)
	if err != nil {
		err = errors.Wrapf(err, "Invalid ID")
		return
	}
	query := map[string]interface{}{
		"Id": id,
	}
	objs, err := orm.Get(query, "")
	if err != nil {
		err = errors.Wrapf(err, "Search error")
		return
	} else if users == nil {
		err = errors.New("Failed to find object")
		return
	} else if len(users) != 1 {
		err = errors.New("Found duplicate objects")
		return
	}
	obj = objs[0]
	return
}
