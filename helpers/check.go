package helpers

import "reflect"

func Typeof(v interface{}) reflect.Kind {
	return reflect.ValueOf(v).Kind()
}
