package gofaker

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	s "github.com/guchengxi1994/go_faker/providers/ssn"
	"github.com/guchengxi1994/go_faker/utils"
)

type SSN struct {
	Locale string
	MaxAge int
	MinAge int
	Gender bool
}

var checkNumber = []string{
	"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2",
}

var multiNumber = []int{
	7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 9, 9, 10, 5, 8, 4, 2,
}

func checksum(s string) string {
	if len(s) != 17 {
		return checkNumber[utils.Randn(len(checkNumber))]
	}
	var sum int
	for i := 0; i < len(s); i++ {
		val, _ := strconv.Atoi(string(s[i]))
		sum += val * multiNumber[i]
	}

	index := sum % 11
	return checkNumber[index]
}

func GenerateRandomDate(max, min int) string {
	if max <= min {
		max, min = min, max
	}

	year := time.Now().Year()
	maxYear := year - min
	minYear := year - max

	minDate := time.Date(minYear, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	maxDate := time.Date(maxYear, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := maxDate - minDate

	sec := int64(utils.Randn(int(delta))) + minDate
	_time := strings.Split(fmt.Sprintf("%v", time.Unix(sec, 0)), " ")[0]

	return strings.ReplaceAll(_time, "-", "")
}

func (ssn *SSN) GenerateSequenceCode() string {
	result := utils.GenerateRandomNumberNew(3)

	code, _ := strconv.Atoi(string(result[2]))
	if ssn.Gender {
		if code%2 == 0 {
			_result, _ := strconv.Atoi(result)
			_result += 1
			result = strconv.Itoa(_result)
		}
	}

	if !ssn.Gender {
		if code%2 != 0 {
			_result, _ := strconv.Atoi(result)
			_result += 1
			result = strconv.Itoa(_result)
		}
	}

	return result
}

func (ssn *SSN) Generate() string {
	var result string

	if ssn.MaxAge == 0 {
		ssn.MaxAge = 75
	}

	if ssn.MinAge == 0 {
		ssn.MinAge = 18
	}

	pre := utils.GetRandomItemFromStringList(s.Ssn_pre)

	date := GenerateRandomDate(ssn.MaxAge, ssn.MinAge)

	scode := ssn.GenerateSequenceCode()

	result = pre + date + scode

	checkcode := checksum(result)

	return result + checkcode
}
