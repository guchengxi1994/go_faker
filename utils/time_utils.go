package utils

import (
	"math/rand"
	"strconv"
	"time"
)

/*
	times is the number join work

	years is the number form when to now
*/
func SplitTime(times, years int) []TimeStruct {
	var result []TimeStruct
	now := time.Now()
	month, _ := strconv.Atoi(time.Now().Format("01"))
	startMonth := Randn(12)
	months := (years-1)*12 + month + (12 - startMonth)
	splitMonths := SplitNumber(times, months)
	for i := 0; i < times; i++ {
		t := TimeStruct{}
		if i == 0 {
			t.To = "Now"
		} else {
			t.To = now.Format("2006-01")
		}
		t.Duration = splitMonths[i]
		_sub_year := -splitMonths[i] / 12
		_sub_month := -splitMonths[i] % 12
		now = now.AddDate(_sub_year, _sub_month, 0)
		t.From = now.Format("2006-01")
		result = append(result, t)
	}

	return result
}

/*
	zhihu.com/question/22625187
	据说是微信抢红包算法
*/
func SplitNumber(times, sum int) []int {
	var result []int
	_remainTimes := times
	_remainSum := sum
	for i := 0; i < times; i++ {
		number := func() int {
			if _remainTimes == 1 {
				_remainTimes--
				return _remainSum
			}
			_min := 1
			_max := _remainSum / _remainTimes * 2
			rand.Seed(time.Now().UnixNano())
			_res := rand.Intn(_max)
			if _res == 0 {
				_res = _min
			}
			_remainTimes--
			_remainSum -= _res

			return _res
		}()
		result = append(result, number)
	}

	return result
}

func SplitEducationYears(years []int, currentAge int, studyAge int) []TimeStruct {
	var _age int = studyAge
	var ts []TimeStruct
	birthYear := time.Now().Year() - currentAge

	for i, _v := range years {
		var t TimeStruct
		if i == 0 {
			t.From = strconv.Itoa(birthYear+studyAge) + "-9"
		} else {
			t.From = strconv.Itoa(_age+birthYear) + "-9"
		}
		_age += _v
		if _age < currentAge {
			t.To = strconv.Itoa(_age+birthYear) + "-7"
			ts = append(ts, t)
		} else {
			t.To = "now"
			ts = append(ts, t)
			break
		}
	}
	return ts

}

func DateTimeBetween() string {
	now := time.Now()
	month, _ := strconv.Atoi(now.Format("01"))
	monthMax := month + (now.Year()-1970)*12
	months := Randn(monthMax)
	_sub_year := -months / 12
	_sub_month := -months % 12
	_dateTime := now.AddDate(_sub_year, _sub_month, 0)
	return _dateTime.Format("2006-01-02")
}
