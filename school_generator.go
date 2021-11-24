package gofaker

import (
	a "github.com/guchengxi1994/go_faker/providers/address"
	s "github.com/guchengxi1994/go_faker/providers/school"
	"github.com/guchengxi1994/go_faker/utils"
)

type School struct {
	Type   int // 0 primary school ,1 middle school, 2 university 3 职业技术学院
	Locale string
}

func (s *School) Generate() string {
	switch s.Locale {
	case "zh_CN":
		return getSchool_ZH(s.Type)
	default:
		return getSchool_ZH(s.Type)
	}
}

func getSchool_ZH(t int) string {
	var suffix string
	if t == 0 {
		suffix = utils.GetRandomItemFromStringList(s.Primary_school_suffix_zh)
	} else if t == 1 {
		suffix = utils.GetRandomItemFromStringList(s.Middle_school_suffix_zh)
	} else if t == 2 {
		suffix = utils.GetRandomItemFromStringList(s.University_suffix_zh)
	} else {
		suffix = "职业技术学院"
	}

	prefix := utils.GetRandomItemFromStringList(a.Zh_districts_simple)

	return prefix + suffix
}
