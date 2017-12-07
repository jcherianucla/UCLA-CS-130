package utilities

import "reflect"

// IsUndeclared uses reflection to see if the
// value of the field is set or not.
// It takes in an interface to reflect on.
// It returns the boolean if the field is set or not.
func IsUndeclared(v interface{}) bool {
	return reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface())
}

func FillStruct(data map[string]interface{}, result interface{}) {
	t := reflect.ValueOf(result).Elem()
	for k, v := range data {
		val := t.FieldByName(k)
		val.Set(reflect.ValueOf(v))
	}
}
