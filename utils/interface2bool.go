package utils

import "reflect"

func InterfaceToBoolean(v interface{}) bool {
	vType := reflect.TypeOf(v)
	switch vType.String() {
	case "bool":
		return v.(bool)
	case "int":
		return v.(int) > 0
	case "float":
		return v.(float64) > 0
	case "string":
		return v.(string) == "true"
	default:
		return false
	}
}
