package providers

import (
	"reflect"
	"regexp"
	"strings"

	address "github.com/guchengxi1994/go_faker/providers/address"
	carlicense "github.com/guchengxi1994/go_faker/providers/car_license"
	company "github.com/guchengxi1994/go_faker/providers/company"
	isbn "github.com/guchengxi1994/go_faker/providers/isbn"
	jobs "github.com/guchengxi1994/go_faker/providers/jobs"
	lorem "github.com/guchengxi1994/go_faker/providers/lorem"
	mobile "github.com/guchengxi1994/go_faker/providers/mobile"
	nickname "github.com/guchengxi1994/go_faker/providers/nickname"
	names "github.com/guchengxi1994/go_faker/providers/person_names"
	school "github.com/guchengxi1994/go_faker/providers/school"
	ssn "github.com/guchengxi1994/go_faker/providers/ssn"
	"github.com/guchengxi1994/go_faker/utils"
	plt "github.com/guchengxi1994/pyLikeType"
)

var global_variants = map[string]*[]string{
	"Zh_provinces_simple":       &address.Zh_provinces_simple,
	"Zh_districts_simple":       &address.Zh_districts_simple,
	"Zh_cities_simple":          &address.Zh_cities_simple,
	"Zh_countries_simple":       &address.Zh_countries_simple,
	"Zh_city_suffixes_simple":   &address.Zh_city_suffixes_simple,
	"Zh_street_suffixes_simple": &address.Zh_street_suffixes_simple,
	"Car_license_prefix":        &carlicense.Car_license_prefix,
	"Car_license_suffix":        &carlicense.Car_license_suffix,
	"Company_prefixs_zh":        &company.Company_prefixs_zh,
	"Company_suffix_zh":         &company.Company_suffix_zh,
	"ISBN13_prefix":             &isbn.ISBN13_prefix,
	"Job_zh":                    &jobs.Job_zh,
	"Lorem_words_zh":            &lorem.Lorem_words_zh,
	"Yidong_prefix":             &mobile.Yidong_prefix,
	"Liantong_prefix":           &mobile.Liantong_prefix,
	"Dianxin_prefix":            &mobile.Dianxin_prefix,
	"Prefix_nickname_male":      &nickname.Prefix_nickname_male,
	"Prefix_nickname_female":    &nickname.Prefix_nickname_female,
	"Common_nicknames":          &nickname.Common_nicknames,
	"First_names_male_zh":       &names.First_names_male_zh,
	"First_fenames_male_zh":     &names.First_fenames_male_zh,
	"Last_names_zh":             &names.Last_names_zh,
	"Primary_school_suffix_zh":  &school.Primary_school_suffix_zh,
	"Middle_school_suffix_zh":   &school.Middle_school_suffix_zh,
	"University_suffix_zh":      &school.University_suffix_zh,
	"Ssn_pre_zh":                &ssn.Ssn_pre_zh,
	"Zh_mobile_all_prefix":      &mobile.Zh_mobile_all_prefix,
}

var global_function = map[string]interface{}{
	"GenerateMobileFromList": utils.GenerateMobileFromList,
}

func Format(pattern string, useWeight bool, args ...float64) string {
	var result string
	result = pattern

	var patternFormat = `{[^}]+}`
	regx, _ := regexp.Compile(patternFormat)
	res := regx.FindAllString(pattern, -1)
	for _, v := range res {
		_v := strings.Replace(v, "{", "", 1)
		_v = strings.Replace(_v, "}", "", 1)
		var _randomStr string

		exv := plt.ExtendString{
			Value: _v,
		}

		if _v == "RANDOM" {
			_randomStr = utils.GenerateRandomNumberNew(4)
		} else if exv.StartsWith("func") {
			var funcName, params string
			_fp := strings.Split(_v, ".")
			if len(_fp) == 1 || len(_fp) > 3 {
				_randomStr = ""
			} else if len(_fp) == 2 {
				funcName = _fp[1]
				_randomStr = applyFunction(funcName)
			} else {
				funcName = _fp[1]
				params = _fp[2]
				params = strings.Replace(params, "[", "", 1)
				params = strings.Replace(params, "]", "", 1)
				_params := strings.Split(params, ",")
				_randomStr = applyFunction(funcName, _params...)
			}

		} else {
			value, ok := global_variants[_v]
			if ok {
				if useWeight {
					m, _ := utils.GenerateWeights(*value, args...)
					_randomStr = utils.GetRandomItemFromMapWithWeight(m)
				} else {
					_randomStr = utils.GetRandomItemFromStringList(*value)
				}
			} else {
				_randomStr = ""
			}
		}
		result = strings.Replace(result, v, _randomStr, 1)
	}

	return result
}

func applyFunction(functionname string, argname ...string) string {
	function, ok := global_function[functionname]
	if !ok {
		return ""
	}
	// without argname
	if len(argname) == 0 {
		values, err := Call(function)
		if err != nil {
			return ""
		} else {
			return values[0].String()
		}
	} else {
		var params []interface{}
		for _, v := range argname {
			_param, ok := global_variants[v]
			if !ok {
				return ""
			}
			if reflect.ValueOf(_param).Kind() == reflect.Ptr {
				_tmp := *_param
				params = append(params, _tmp)
			} else {
				params = append(params, _param)
			}

		}
		values, err := Call(function, params...)
		if err != nil {
			return ""
		} else {
			return values[0].String()
		}
	}

}
