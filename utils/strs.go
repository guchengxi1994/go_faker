package utils

import (
	"reflect"
	"strconv"
)

var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

/*
	Following https://github.com/joke2k/faker/blob/d7c9b118ecb6f6e741f07cfa7fe62e965da14ebc/faker/providers/__init__.py#L612
*/
func Numerify(text string) string {
	var result string
	for _, s := range text {
		switch string(s) {
		case "#":
			result += strconv.Itoa(Randn(9))
		case "%":
			result += strconv.Itoa(Randn(8) + 1)
		case "!":
			v := Randn(10)
			if v == 10 {
				result += ""
			} else {
				result += strconv.Itoa(v)
			}
		case "@":
			v := Randn(9)
			if v == 9 {
				result += ""
			} else {
				result += strconv.Itoa(v + 1)
			}

		default:
			result += string(s)
		}
	}

	return result
}

// following https://github.com/joke2k/faker/blob/d7c9b118ecb6f6e741f07cfa7fe62e965da14ebc/faker/providers/__init__.py#L636
func Lexify(text string) string {
	var result string
	for _, s := range text {
		switch string(s) {
		case "?":
			result += GetRandomRuneFromString(letters)
		default:
			result += string(s)
		}
	}

	return result
}

// following https://github.com/joke2k/faker/blob/d7c9b118ecb6f6e741f07cfa7fe62e965da14ebc/faker/providers/__init__.py#L647
func Bothify(text string) string {
	var result string
	for _, s := range text {
		switch string(s) {
		case "?":
			result += GetRandomRuneFromString(letters)
		case "#":
			result += strconv.Itoa(Randn(10))
		default:
			result += string(s)
		}
	}
	return result
}

// following https://github.com/joke2k/faker/blob/d7c9b118ecb6f6e741f07cfa7fe62e965da14ebc/faker/providers/user_agent/__init__.py#L142
func Saf() string {
	return strconv.Itoa(Randn(6)+531) + "." + strconv.Itoa(Randn(2))
}

func Chrome_version(from, to interface{}) string {
	switch reflect.TypeOf(from).Name() {
	case "string":
		min, _ := strconv.Atoi(from.(string))
		max, _ := strconv.Atoi(to.(string))
		return strconv.Itoa(Randn(max-min) + min)

	case "int":
		min := from.(int)
		max := to.(int)
		return strconv.Itoa(Randn(max-min) + min)
	default:
		return "60"
	}
}

func Chrome_build_version(from, to interface{}) string {
	switch reflect.TypeOf(from).Name() {
	case "string":
		min, _ := strconv.Atoi(from.(string))
		max, _ := strconv.Atoi(to.(string))
		return strconv.Itoa(Randn(max-min) + min)

	case "int":
		min := from.(int)
		max := to.(int)
		return strconv.Itoa(Randn(max-min) + min)
	default:
		return "849"
	}
}
