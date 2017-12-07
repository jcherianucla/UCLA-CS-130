package utilities

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
)

// IsUndeclared uses reflection to see if the
// value of the field is set or not.
// It takes in an interface to reflect on.
// It returns the boolean if the field is set or not.
func IsUndeclared(v interface{}) bool {
	return reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface())
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
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type")
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
