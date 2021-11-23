package gofaker

import (
	c "github.com/guchengxi1994/go_faker/providers/company"
	"github.com/guchengxi1994/go_faker/utils"
)

type Company struct {
	Locale string
	Args   []float64 // weight
}

func (company *Company) Generate() string {
	w, err := utils.GenerateWeights(c.Company_prefixs_zh, company.Args...)
	switch company.Locale {
	case "zh_CN":
		if len(company.Args) != 0 && err == nil {
			return utils.GetRandomItemFromMapWithWeight(w) + utils.GetRandomItemFromStringList(c.Company_suffix_zh)
		}
		return utils.GetRandomItemFromStringList(c.Company_prefixs_zh) + utils.GetRandomItemFromStringList(c.Company_suffix_zh)
	default:
		if len(company.Args) != 0 && err == nil {
			return utils.GetRandomItemFromMapWithWeight(w) + utils.GetRandomItemFromStringList(c.Company_suffix_zh)
		}
		return utils.GetRandomItemFromStringList(c.Company_prefixs_zh) + utils.GetRandomItemFromStringList(c.Company_suffix_zh)
	}
}
