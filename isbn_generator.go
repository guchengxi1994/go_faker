package gofaker

import (
	"fmt"
	"strconv"
	"strings"

	is "github.com/guchengxi1994/go_faker/providers/isbn"
	"github.com/guchengxi1994/go_faker/utils"
)

type ISBN struct {
	ISBN13 string
	ISBN10 string
}

var weights10 = []int{
	10, 9, 8, 7, 6, 5, 4, 3, 2,
}

var weight13 = []int{
	1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
}

func (isbn *ISBN) Generate() {
	pre := utils.GetRandomItemFromStringList(is.ISBN13_prefix)
	preList := strings.Split(pre, "-")
	i13pre := preList[0]
	code := preList[1]
	restLength := 10 - len(code) - 1 // 1 is the checknum

	pressLength := utils.Randn(restLength-2) + 2
	bookLength := restLength - pressLength

	pressCode := utils.GenerateRandomNumberNew(pressLength)
	bookCode := utils.GenerateRandomNumberNew(bookLength)

	isbn10checknum := checkNumber10(code + pressCode + bookCode)
	isbn13checknum := checkNumber13(i13pre + code + pressCode + bookCode)

	isbn.ISBN10 = code + "-" + pressCode + "-" + bookCode + "-" + isbn10checknum

	isbn.ISBN13 = i13pre + "-" + code + "-" + pressCode + "-" + bookCode + "-" + isbn13checknum

}

func checkNumber10(str string) string {
	var sum int
	for i := 0; i < len(str); i++ {
		val, _ := strconv.Atoi(string(str[i]))
		sum += val * weights10[i]
	}
	index := 11 - sum%11

	if index == 10 {
		return "X"
	} else if index == 11 {
		return "0"
	} else {
		return strconv.Itoa(index)
	}
}

func checkNumber13(str string) string {
	var sum int
	for i := 0; i < len(str); i++ {
		val, _ := strconv.Atoi(string(str[i]))
		sum += val * weight13[i]
	}
	index := 10 - sum%10

	if sum%10 == 0 {
		return "0"
	} else {
		return strconv.Itoa(index)
	}
}

func (isbn *ISBN) ToString() string {
	return fmt.Sprintf("isbn10:%s,isbn13:%s", isbn.ISBN10, isbn.ISBN13)
}
