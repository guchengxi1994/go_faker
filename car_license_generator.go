package gofaker

import (
	cl "github.com/guchengxi1994/go_faker/providers/car_license"
	"github.com/guchengxi1994/go_faker/utils"
)

type Carlicense struct {
	Locale string
}

func (c *Carlicense) Generate() string {
	switch c.Locale {
	case "zh_CN":
		return utils.GetRandomItemFromStringList(cl.Car_license_prefix) + "." + utils.GenerateRandomStrFromList(utils.Randn(2)+4, cl.Car_license_suffix)
	default:
		return utils.GetRandomItemFromStringList(cl.Car_license_prefix) + "." + utils.GenerateRandomStrFromList(utils.Randn(2)+4, cl.Car_license_suffix)
	}
}
