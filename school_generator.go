package gofaker

import (
	provider "github.com/guchengxi1994/go_faker/providers"
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
		suffix = provider.Format(provider.Format_Zh_simple_primary_school, false)
	} else if t == 1 {
		suffix = provider.Format(provider.Format_Zh_simple_middle_school, false)
	} else if t == 2 {
		suffix = provider.Format(provider.Format_Zh_simple_university, false)
	} else {
		suffix = "职业技术学院"
	}

	prefix := provider.Format(provider.Format_Zh_simple_district, false)

	return prefix + suffix
}
