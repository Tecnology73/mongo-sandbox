package utilities

import (
	"reflect"
)

func TypeOf[T any]() (elemType reflect.Type) {
	elemType = reflect.TypeOf((*T)(nil)).Elem()

	if elemType.Kind() == reflect.Ptr {
		elemType = elemType.Elem()
	}

	return
}
