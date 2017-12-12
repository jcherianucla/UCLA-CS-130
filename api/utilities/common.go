package utilities

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"reflect"
	"time"
)

var (
	// Skip these fields from an old version of the database table
	SKIP_PARAM = map[string]bool{
		"Post_results": true,
		"Pre_results":  true,
	}
)

// IsUndeclared uses reflection to see if the
// value of the field is set or not.
// It takes in an interface to reflect on.
// It returns the boolean if the field is set or not.
func IsUndeclared(v interface{}) bool {
	return reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface())
}

// GetVar tries to get an environment variable if it exists, else
// returns some default value.
// It takes in the name of the environment variable to search for, and the
// default value to return instead.
// It returns the string variable.
func GetVar(name string, _default string) string {
	env := os.Getenv(name)
	if env == "" {
		return _default
	}
	return env
}

// SetField sets the value of the generic of the name key. That means
// for a given struct, the field will be set to the incoming value.
// It takes in the object to modify, the name of the field and the value the
// field should be set to.
// It returns an error if one exists.
func SetField(obj interface{}, key string, value interface{}) error {
	structFieldValue := reflect.ValueOf(obj).Elem().FieldByName(key)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", key)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", key)
	}

	structFieldType := structFieldValue.Type()
	var val reflect.Value
	// Type assertion required for a few types
	if key == "Password" {
		t, _ := value.([]byte)
		val = reflect.ValueOf(string(t))
	} else if key == "Lang" {
		v, _ := value.(int64)
		// Comply with Language type alias
		t := GetLanguageFromInt(v)
		val = reflect.ValueOf(t)
	} else {
		val = reflect.ValueOf(value)
		if structFieldType != val.Type() && !SKIP_PARAM[key] {
			return errors.New("Provided value type didn't match obj field type")
		}
	}
	structFieldValue.Set(val)
	return nil
}

// FillStruct creates a generic object through mapped data.
// It takes in data as a map of strings to generics and reference to the object to construct.
// It returns an error if one exists.
func FillStruct(data map[string]interface{}, result interface{}) error {
	for k, v := range data {
		err := SetField(result, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

// Determines whether the current time is before the intended deadline, converting
// down to Milliseconds to ensure timezone issues don't come into effect.
// It takes in the deadline to compare against.
// It returns the boolean value indicating the result.
func BeforeDeadline(deadline time.Time) bool {
	t := (time.Now().UnixNano() / int64(time.Millisecond))
	d := deadline.UnixNano()/int64(time.Millisecond) + (8 * 3600 * 1000)
	return (t < d)
}
