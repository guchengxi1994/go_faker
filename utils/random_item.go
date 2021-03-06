package utils

import (
	"fmt"
	"strings"

	locales "github.com/guchengxi1994/go_faker/providers/locales"
)

func GetRandomItemFromMapWithWeight(items map[string]float64) string {
	_values := GetValues(items)
	_minVal := Minimum(true, _values)
	_factor := 1 / _minVal
	_tmp := turn_float_to_int(items, _factor)

	var choices []Choice

	for k, v := range _tmp {
		choices = append(choices, Choice{
			Item:   k,
			Weight: v,
		})

	}
	chooser, err := NewChooser(choices...)
	if err != nil {
		fmt.Println(err)
		keys := GetKeys(items)
		return GetRandomItemFromStringList(keys)
	}
	return chooser.Pick().(string)
}

func GetRandomItemFromStringList(items []string) string {
	index := Randn(len(items))
	return items[index]
}

func GetRandomItemFromIntList(items []int) int {
	index := Randn(len(items))
	return items[index]
}

func turn_float_to_int(items map[string]float64, factor float64) map[string]uint {
	result := make(map[string]uint)
	for k, v := range items {
		result[k] = uint(int(v * factor))
	}
	return result
}

func GetRandomRuneFromString(s string) string {
	index := Randn(len([]rune(s)))
	return string([]rune(s)[index])
}

func RandomLocale() string {
	return strings.ReplaceAll(GetRandomItemFromStringList(locales.Locales), "_", "-")
}
