package providers

import "github.com/guchengxi1994/go_faker/utils"

func generate(pattern string, list []string, useWeight bool, args ...float64) string {
	var result string
	if useWeight {
		m, _ := utils.GenerateWeights(list, args...)
		result = utils.GetRandomItemFromMapWithWeight(m)
	} else {
		result = utils.GetRandomItemFromStringList(list)
	}

	return result
}
