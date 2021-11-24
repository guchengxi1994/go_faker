package utils

import (
	"math"
	"sort"
	"strconv"
)

func AddToList(list []string, item string) []string {
	if !IsExist(list, item) {
		list = append(list, item)
	}
	return list
}

func IsExist(list []string, item string) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}
	return false
}

func Sum(list []float64) float64 {
	var result float64
	for _, v := range list {
		result += v
	}
	return result
}

func OverZero(list []float64) bool {
	for _, v := range list {
		if v < 0 {
			return false
		}
	}
	return true
}

func Minimum(ignoreZero bool, list []float64) float64 {
	_list := FloatList{}
	_list = list
	sort.Sort(_list)

	for _, i := range _list {
		if i != 0 {
			return i
		}
	}
	return _list[_list.Len()-1]
}

func GetValues(amap map[string]float64) []float64 {
	vals := make([]float64, len(amap))
	for _, v := range amap {
		vals = append(vals, v)
	}
	return vals
}

func GetKeys(amap map[string]float64) []string {
	keys := make([]string, len(amap))
	for k := range amap {
		keys = append(keys, k)
	}
	return keys
}

func GenerateRandomNumber(length int) string {
	var result string
	for i := 0; i < length; i++ {
		result += strconv.Itoa(Randn(10))
	}
	return result
}

func GenerateRandomNumberNew(length int) string {
	if length == 0 {
		length = 1
	}
	max := math.Pow(10, float64(length))

	s := strconv.Itoa(Randn(int(max)))

	if len(s) < length {
		return multiple("0", length-len(s)) + s
	}

	return s
}

func StringArrayContains(strs []string, pattern string) bool {
	for _, v := range strs {
		if v == pattern {
			return true
		}
	}

	return false
}

func multiple(s string, times int) string {
	var result string
	for i := 0; i < times; i++ {
		result += s
	}
	return result
}

func StrFirstToUpper(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArry := []rune(str)
	if strArry[0] >= 97 && strArry[0] <= 122 {
		strArry[0] = strArry[0] - 32
	}
	return string(strArry)
}
