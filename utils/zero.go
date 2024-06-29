package utils

import "reflect"

func IsZeroValue(i any) bool {
	if i == nil {
		return true
	}
	return reflect.Zero(reflect.TypeOf(i)).Interface() == i
}
