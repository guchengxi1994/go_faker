package gofaker

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	p "github.com/guchengxi1994/go_faker/providers"
	utils "github.com/guchengxi1994/go_faker/utils"
)

// type BaseAddress interface {
// 	City(args ...float64) string
// 	BuildingNumber(int) string
// 	District(args ...float64) string
// 	Postcode(int) string
// 	StreetName(args ...float64) string
// 	StreetAddress(args ...float64) string
// 	Country() string
// 	Province(args ...float64) string
// }

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
			result = p.Format(p.Format_Zh_simple_city, true, args...)
		} else {
			result = p.Format(p.Format_Zh_simple_city, false)
		}
	default:
		if addr.UseWeighting && len(args) > 0 {
			result = p.Format(p.Format_Zh_simple_city, true, args...)
		} else {
			result = p.Format(p.Format_Zh_simple_city, false)
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
			result = p.Format(p.Format_Zh_simple_district, true, args...)
		} else {
			result = p.Format(p.Format_Zh_simple_district, false)
		}
	default:
		if addr.UseWeighting && len(args) > 0 {
			result = p.Format(p.Format_Zh_simple_district, true, args...)
		} else {
			result = p.Format(p.Format_Zh_simple_district, false)
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
			result = p.Format(p.Format_Zh_simple_street_name, true, args...)
		} else {
			result = p.Format(p.Format_Zh_simple_street_name, false)
		}
	default:
		if addr.UseWeighting && len(args) > 0 {
			result = p.Format(p.Format_Zh_simple_street_name, true, args...)
		} else {
			result = p.Format(p.Format_Zh_simple_street_name, false)
		}
	}

	return result
}

func (addr *Address) StreetAddress(length int, args ...float64) string {
	if len(args) == 0 && addr.Args != nil {
		args = addr.Args
	}
	var result string
	switch addr.Locole {
	case "zh_CN":
		if addr.UseWeighting && len(args) > 0 {
			result = p.Format(p.Format_Zh_simple_street_address, true, args...)
		} else {
			result = p.Format(p.Format_Zh_simple_street_address, false)
		}
	default:
		if addr.UseWeighting && len(args) > 0 {
			result = p.Format(p.Format_Zh_simple_street_address, true, args...)
		} else {
			result = p.Format(p.Format_Zh_simple_street_address, false)
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
			result = p.Format(p.Format_Zh_simple_province, true, args...)
		} else {
			result = p.Format(p.Format_Zh_simple_province, false)
		}
	default:
		if addr.UseWeighting && len(args) > 0 {
			result = p.Format(p.Format_Zh_simple_province, true, args...)
		} else {
			result = p.Format(p.Format_Zh_simple_province, false)
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
		if addr.UseWeighting && len(args) > 0 {
			result = p.Format(p.Format_Zh_simple_address, true, args...)
		} else {
			result = p.Format(p.Format_Zh_simple_address, false)
		}
	default:
		if addr.UseWeighting && len(args) > 0 {
			result = p.Format(p.Format_Zh_simple_address, true, args...)
		} else {
			result = p.Format(p.Format_Zh_simple_address, false)
		}
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
