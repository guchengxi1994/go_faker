package gofaker

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	p "github.com/guchengxi1994/go_faker/providers"
	a "github.com/guchengxi1994/go_faker/providers/address"
	utils "github.com/guchengxi1994/go_faker/utils"
)

type BaseAddress interface {
	City(args ...float64) string
	BuildingNumber(int) string
	District(args ...float64) string
	Postcode(int) string
	StreetName(args ...float64) string
	StreetAddress(args ...float64) string
	Country() string
	Province(args ...float64) string
}

type Address struct {
	PostcodeLength int
	Args           []float64 // weight
	UseWeighting   bool
	Locole         string
}

type Location struct {
	Longitude string // 经度  0-180
	Latitude  string // 维度  0-90
	WE        string
	NS        string
}

func (addr *Address) BuildingNumber(length int) string {
	var result string
	buildCode := utils.GenerateRandomNumberNew(length)
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
	if len(args) == 0 && addr.Args != nil {
		args = addr.Args
	}

	switch addr.Locole {
	case "zh_CN":
		if addr.UseWeighting && len(args) > 0 {
			m, _ := utils.GenerateWeights(a.Zh_cities_simple, args...)
			result = fmt.Sprintf(p.Zh_city_format_simple, utils.GetRandomItemFromMapWithWeight(m))
		} else {
			result = fmt.Sprintf(p.Zh_city_format_simple, utils.GetRandomItemFromStringList(a.Zh_cities_simple))
		}
	default:
		if addr.UseWeighting && len(args) > 0 {
			m, _ := utils.GenerateWeights(a.Zh_cities_simple, args...)
			result = fmt.Sprintf(p.Zh_city_format_simple, utils.GetRandomItemFromMapWithWeight(m))
		} else {
			result = fmt.Sprintf(p.Zh_city_format_simple, utils.GetRandomItemFromStringList(a.Zh_cities_simple))
		}
	}

	return result
}

func (addr *Address) District(args ...float64) string {
	if len(args) == 0 && addr.Args != nil {
		args = addr.Args
	}
	var result string
	switch addr.Locole {
	case "zh_CN":
		if addr.UseWeighting && len(args) > 0 {
			m, _ := utils.GenerateWeights(a.Zh_districts_simple, args...)
			result = fmt.Sprintf(p.Zh_district_format_simple, utils.GetRandomItemFromMapWithWeight(m))
		} else {
			result = fmt.Sprintf(p.Zh_district_format_simple, utils.GetRandomItemFromStringList(a.Zh_districts_simple))
		}
	default:
		if addr.UseWeighting && len(args) > 0 {
			m, _ := utils.GenerateWeights(a.Zh_districts_simple, args...)
			result = fmt.Sprintf(p.Zh_district_format_simple, utils.GetRandomItemFromMapWithWeight(m))
		} else {
			result = fmt.Sprintf(p.Zh_district_format_simple, utils.GetRandomItemFromStringList(a.Zh_districts_simple))
		}
	}

	return result
}

func (addr *Address) StreetName(args ...float64) string {
	if len(args) == 0 && addr.Args != nil {
		args = addr.Args
	}
	var result string
	switch addr.Locole {
	case "zh_CN":
		if addr.UseWeighting && len(args) > 0 {
			m, _ := utils.GenerateWeights(a.Zh_cities_simple, args...)
			result = fmt.Sprintf(p.Zh_street_name_format_simple, utils.GetRandomItemFromMapWithWeight(m), utils.GetRandomItemFromStringList(a.Zh_street_suffixes_simple))
		} else {
			result = fmt.Sprintf(p.Zh_street_name_format_simple, utils.GetRandomItemFromStringList(a.Zh_cities_simple), utils.GetRandomItemFromStringList(a.Zh_street_suffixes_simple))
		}
	default:
		if addr.UseWeighting && len(args) > 0 {
			m, _ := utils.GenerateWeights(a.Zh_cities_simple, args...)
			result = fmt.Sprintf(p.Zh_street_name_format_simple, utils.GetRandomItemFromMapWithWeight(m), utils.GetRandomItemFromStringList(a.Zh_street_suffixes_simple))
		} else {
			result = fmt.Sprintf(p.Zh_street_name_format_simple, utils.GetRandomItemFromStringList(a.Zh_cities_simple), utils.GetRandomItemFromStringList(a.Zh_street_suffixes_simple))
		}
	}

	return result
}

func (addr *Address) StreetAddress(length int, args ...float64) string {
	if len(args) == 0 && addr.Args != nil {
		args = addr.Args
	}
	var result string
	streetName := addr.StreetName(args...)
	switch addr.Locole {
	case "zh_CN":
		if addr.UseWeighting && len(args) > 0 {
			result = fmt.Sprintf(p.Zh_street_address_format_simple, streetName, addr.BuildingNumber(length))
		} else {
			result = fmt.Sprintf(p.Zh_street_address_format_simple, streetName, addr.BuildingNumber(length))
		}
	default:
		if addr.UseWeighting && len(args) > 0 {
			result = fmt.Sprintf(p.Zh_street_address_format_simple, streetName, addr.BuildingNumber(length))
		} else {
			result = fmt.Sprintf(p.Zh_street_address_format_simple, streetName, addr.BuildingNumber(length))
		}
	}

	return result
}

func (addr *Address) Province(args ...float64) string {
	if len(args) == 0 && addr.Args != nil {
		args = addr.Args
	}
	var result string
	switch addr.Locole {
	case "zh_CN":
		if addr.UseWeighting && len(args) > 0 {
			m, _ := utils.GenerateWeights(a.Zh_provinces_simple, args...)
			result = utils.GetRandomItemFromMapWithWeight(m)
		} else {
			result = utils.GetRandomItemFromStringList(a.Zh_provinces_simple)
		}
	default:
		if addr.UseWeighting && len(args) > 0 {
			m, _ := utils.GenerateWeights(a.Zh_provinces_simple, args...)
			result = utils.GetRandomItemFromMapWithWeight(m)
		} else {
			result = utils.GetRandomItemFromStringList(a.Zh_provinces_simple)
		}
	}

	return result
}

// generate
func (addr *Address) Address(args ...float64) string {
	if len(args) == 0 && addr.Args != nil {
		args = addr.Args
	}
	if addr.PostcodeLength == 0 {
		addr.PostcodeLength = 3
	}
	var result string
	switch addr.Locole {
	case "zh_CN":
		result = fmt.Sprintf(p.Zh_address_format_simple, addr.Province(args...), addr.City(args...), addr.District(args...), addr.StreetAddress(addr.PostcodeLength, args...))
	default:
		result = fmt.Sprintf(p.Zh_address_format_simple, addr.Province(args...), addr.City(args...), addr.District(args...), addr.StreetAddress(addr.PostcodeLength, args...))
	}
	return result
}

func (location *Location) GetPositon() string {
	randLongInt := utils.Randn(65535)
	randLongFloat := float64(randLongInt) / 65535 * 180
	randLatInt := utils.Randn(65535)
	randLatFloat := float64(randLatInt) / 65535 * 90

	location.Latitude = strconv.FormatFloat(randLatFloat, 'f', 5, 64)
	location.Longitude = strconv.FormatFloat(randLongFloat, 'f', 5, 64)
	we := func() string {
		rand.Seed(time.Now().UnixNano())
		if rand.Int31n(2) == 1 {
			return "E"
		}
		return "W"
	}

	ns := func() string {
		rand.Seed(time.Now().UnixNano())
		if rand.Int31n(2) == 1 {
			return "N"
		}
		return "S"
	}

	location.WE = we()
	location.NS = ns()

	return fmt.Sprintf("%s %s,%s %s", location.NS, location.Latitude, location.WE, location.Longitude)

}
