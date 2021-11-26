package gofaker

import (
	"fmt"

	p "github.com/guchengxi1994/go_faker/providers"
)

type Mobile struct {
	Locale   string
	City     string
	Supplier string
}

func (mobile *Mobile) Generate() string {
	switch mobile.Locale {
	case "zh_CN":
		var _supplier_param_name string
		if mobile.Supplier == "yidong" {
			_supplier_param_name = "Yidong_prefix"
		} else if mobile.Supplier == "liantong" {
			_supplier_param_name = "Liantong_prefix"
		} else if mobile.Supplier == "dianxin" {
			_supplier_param_name = "Dianxin_prefix"
		} else {
			_supplier_param_name = "Zh_mobile_all_prefix"
		}
		runStr := fmt.Sprintf(p.Function_Zh_mobile, _supplier_param_name)
		return p.Format(runStr, false)
	default:
		var _supplier_param_name string
		if mobile.Supplier == "yidong" {
			_supplier_param_name = "Yidong_prefix"
		} else if mobile.Supplier == "liantong" {
			_supplier_param_name = "Liantong_prefix"
		} else if mobile.Supplier == "dianxin" {
			_supplier_param_name = "Dianxin_prefix"
		} else {
			_supplier_param_name = "Zh_mobile_all_prefix"
		}
		runStr := fmt.Sprintf(p.Function_Zh_mobile, _supplier_param_name)
		return p.Format(runStr, false)
	}
}
