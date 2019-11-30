package moocro

import (
	"reflect"
)

// ToType returns a string representation of the value
func ToType(value interface{}) string {
	return reflect.TypeOf(value).String()
}
