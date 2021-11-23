package gofaker

import (
	pm "github.com/guchengxi1994/go_faker/providers/mobile"
	"github.com/guchengxi1994/go_faker/utils"
)

type Mobile struct {
	Locale string
	City   string
}

func (mobile *Mobile) Generate(supplier string) string {
	switch mobile.Locale {
	case "zh_CN":
		return generateFromSupplier_Zh(supplier)
	default:
		return generateFromSupplier_Zh(supplier)
	}
}

func generateFromSupplier_Zh(supplier string) string {
	switch supplier {
	case "yidong":
		return generateMobile(pm.Yidong_prefix, 11)
	case "liantong":
		return generateMobile(pm.Liantong_prefix, 11)
	case "dianxin":
		return generateMobile(pm.Dianxin_prefix, 11)
	default:
		var pres []string
		pres = append(pres, pm.Dianxin_prefix...)
		pres = append(pres, pm.Liantong_prefix...)
		pres = append(pres, pm.Yidong_prefix...)
		return generateMobile(pres, 11)
	}
}

func generateMobile(prefix []string, sufLength int) string {
	pre := utils.GetRandomItemFromStringList(prefix)
	suf := utils.GenerateRandomNumberNew(sufLength - len(pre))
	return pre + suf
}
