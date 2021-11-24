package gofaker

import (
	pn "github.com/guchengxi1994/go_faker/providers/person_names"
	utils "github.com/guchengxi1994/go_faker/utils"
	"github.com/mozillazg/go-pinyin"
)

type PersonName struct {
	Locale    string
	Firstname string
	Lastname  string
	Gender    bool
	Pinyin    string //only for chinese names
	generated bool
}

func (p *PersonName) Generate(args ...float64) {
	p.generated = true
	switch p.Locale {
	case "zh_CN":
		if len(args) > 0 {
			m, err := utils.GenerateWeights(pn.Last_names_zh, args...)
			if err != nil {
				p.Lastname = utils.GetRandomItemFromStringList(pn.Last_names_zh)
			}
			p.Lastname = utils.GetRandomItemFromMapWithWeight(m)
		} else {
			p.Lastname = utils.GetRandomItemFromStringList(pn.Last_names_zh)
		}
		if p.Gender {
			p.Firstname = utils.GetRandomItemFromStringList(pn.First_names_male_zh)
		} else {
			p.Firstname = utils.GetRandomItemFromStringList(pn.First_fenames_male_zh)
		}

	default:
		if len(args) > 0 {
			m, err := utils.GenerateWeights(pn.Last_names_zh, args...)
			if err != nil {
				p.Lastname = utils.GetRandomItemFromStringList(pn.Last_names_zh)
			}
			p.Lastname = utils.GetRandomItemFromMapWithWeight(m)
		} else {
			p.Lastname = utils.GetRandomItemFromStringList(pn.Last_names_zh)
		}
		if p.Gender {
			p.Firstname = utils.GetRandomItemFromStringList(pn.First_names_male_zh)
		} else {
			p.Firstname = utils.GetRandomItemFromStringList(pn.First_fenames_male_zh)
		}
	}
}

func (p *PersonName) ToString(withPinyin bool) string {
	var _pinyin string = ", "
	if !p.generated {
		p.Generate()
	}
	switch p.Locale {
	case "zh_CN":
		if withPinyin {
			_firstnamePinyinList := pinyin.LazyConvert(p.Firstname, nil)
			for i, v := range _firstnamePinyinList {
				if i == 0 {
					v = utils.StrFirstToUpper(v)
				}
				_pinyin += v
			}
			_pinyin += " "
			_lastnamePinyinList := pinyin.LazyConvert(p.Lastname, nil)
			for i, v := range _lastnamePinyinList {
				if i == 0 {
					v = utils.StrFirstToUpper(v)
				}
				_pinyin += v
			}
			return p.Lastname + p.Firstname + _pinyin
		}
		return p.Lastname + p.Firstname
	default:
		return p.Lastname + p.Firstname
	}
}
