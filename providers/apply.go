package providers

import (
	"errors"
	"reflect"
)

func Call(name interface{}, params ...interface{}) ([]reflect.Value, error) {
	f := reflect.ValueOf(name)
	if len(params) != f.Type().NumIn() {
		return nil, errors.New("the number of input params not match")
	}
	in := make([]reflect.Value, len(params))
	for k, v := range params {
		in[k] = reflect.ValueOf(v)
	}
	return f.Call(in), nil
}
