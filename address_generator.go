package gofaker

import (
	"fmt"
	p "guchengxi1994/go_faker/providers"
	a "guchengxi1994/go_faker/providers/address"
	utils "guchengxi1994/go_faker/utils"
)

type BaseAddress interface {
	City(args ...float64) string
	BuildingNumber(int) string
	District() string
	Postcode(int) string
	StreetName() string
	Country() string
	Province() string
}

type Address struct {
	Unique       bool
	UseWeighting bool
	Locole       string
}

func (addr *Address) BuildingNumber(length int) string {
	var result string
	buildCode := utils.GenerateRandomNumber(length)
	switch addr.Locole {
	case "zh_CN":
		result = fmt.Sprintf(p.Zh_building_number_format_simple, buildCode)
	default:
		result = fmt.Sprintf(p.Zh_building_number_format_simple, buildCode)
	}

	return result
}

func (addr *Address) City(args ...float64) string {
	var result string
	switch addr.Locole {
	case "zh_CN":
		if addr.UseWeighting && len(args) > 0 {
			m, _ := utils.GenerateWeights(a.Zh_cities_simple, args...)
			result = utils.GetRandomItemFromMapWithWeight(m)
		} else {
			result = utils.GetRandomItemFromStringList(a.Zh_cities_simple)
		}
	default:
		if addr.UseWeighting && len(args) > 0 {
			m, _ := utils.GenerateWeights(a.Zh_cities_simple, args...)
			result = utils.GetRandomItemFromMapWithWeight(m)
		} else {
			result = utils.GetRandomItemFromStringList(a.Zh_cities_simple)
		}
	}

	return result
}
