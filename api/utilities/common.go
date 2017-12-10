package utilities

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"reflect"
)

// IsUndeclared uses reflection to see if the
// value of the field is set or not.
// It takes in an interface to reflect on.
// It returns the boolean if the field is set or not.
func IsUndeclared(v interface{}) bool {
	return reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface())
}

func GetVar(name string, _default string) string {
	env := os.Getenv(name)
	if env == "" {
		return _default
	}
	return env
}

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
	if key == "Password" {
		t, _ := value.([]byte)
		val = reflect.ValueOf(string(t))
	} else if key == "Lang" {
		t, _ := value.(Language)
		val = reflect.ValueOf(t)
	} else {
		val = reflect.ValueOf(value)
		if structFieldType != val.Type() && key != "Lang" {
			return errors.New("Provided value type didn't match obj field type")
		}
	}
	structFieldValue.Set(val)
	return nil
}

func FillStruct(data map[string]interface{}, result interface{}) error {
	for k, v := range data {
		err := SetField(result, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
