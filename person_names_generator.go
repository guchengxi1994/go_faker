package gofaker

import (
	provider "github.com/guchengxi1994/go_faker/providers"
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
		p.Lastname = provider.Format(provider.Format_Zh_simple_lastname, true, args...)
		if p.Gender {
			p.Firstname = provider.Format(provider.Format_Zh_simple_male_firstname, false)
		} else {
			p.Firstname = provider.Format(provider.Format_Zh_simple_female_firstname, false)
		}

	default:
		p.Lastname = provider.Format(provider.Format_Zh_simple_lastname, true, args...)
		if p.Gender {
			p.Firstname = provider.Format(provider.Format_Zh_simple_male_firstname, false)
		} else {
			p.Firstname = provider.Format(provider.Format_Zh_simple_female_firstname, false)
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
