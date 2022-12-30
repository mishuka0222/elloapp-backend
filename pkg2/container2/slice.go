package container2

import (
	"fmt"
	"reflect"
)

func Contains(obj interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	}
	return false, fmt.Errorf("not in")
}

func AppendIgnoreNil(slice []interface{}, elems ...interface{}) []interface{} {
	for _, e := range elems {
		if e != nil {
			slice = append(slice, e)
		}
	}
	return slice
}
