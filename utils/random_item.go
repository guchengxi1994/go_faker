package utils

import (
	"fmt"
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
	// rand.Seed(time.Now().UnixNano())
	// index := rand.Intn(len(items))
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
