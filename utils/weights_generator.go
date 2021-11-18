package utils

import "errors"

func GenerateWeights(list []string, args ...float64) (map[string]float64, error) {
	if len(list) < len(args) {
		return nil, errors.New("weights length too long")
	}

	if !OverZero(args) {
		return nil, errors.New("weights cannot below 0")
	}

	var _sumWeights = Sum(args)

	if _sumWeights > 1.0 {
		return nil, errors.New("weights sum cannot over 1.0")
	}

	var _extraLength int
	var _average float64

	if len(list) != len(args) {
		_extraLength = len(list) - len(args)
		_average = (1.0 - _sumWeights) / float64(_extraLength)
	} else {
		_extraLength = 0
		_average = 0
	}

	result := make(map[string]float64)

	for i := 0; i < len(list); i++ {
		if i < len(args) {
			result[list[i]] = args[i]
		} else {
			result[list[i]] = _average
		}
	}
	return result, nil
}
