package gofaker

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/guchengxi1994/go_faker/data"
	"github.com/guchengxi1994/go_faker/utils"
)

type Profile struct {
	Name            string `json:"name"`
	Gender          bool   `json:"gender"`
	IdNumber        string `json:"idNumber"`
	MobileNumber    string `json:"mobile"`
	Locale          string
	HomeAddress     string       `json:"homeAddress"`
	WorkExperiences []Experience `json:"workExperiences"`
	Description     string       `json:"description"`
	generated       bool
	Age             int `json:"age"`
}

func (p *Profile) Generate() {
	p.generated = true

	pname := PersonName{
		Locale: p.Locale,
		Gender: p.Gender,
	}

	pname.Generate()

	p.Name = pname.ToString(true)
	rand.Seed(time.Now().UnixNano())
	p.Gender = rand.Int31n(2) == 1
	ssn := SSN{
		Locale: p.Locale,
		MaxAge: 65,
		MinAge: 20,
		Gender: p.Gender,
	}
	p.IdNumber = ssn.Generate()
	mobile := Mobile{
		Locale: p.Locale,
	}
	p.MobileNumber = mobile.Generate("")
	homeAddress := Address{
		Locole:       p.Locale,
		UseWeighting: false,
	}
	p.HomeAddress = homeAddress.Address()

	p.Age = ssn.currentAge

	workAge := determineWorkAge(ssn.currentAge)

	var workTimes int

	workTimes = workAge / 2

	if workTimes > 10 {
		workTimes = 10
	}

	p.WorkExperiences = GenerateWorkExperiences(workTimes, p.Age-workAge, p.Locale)

	p.Description = utils.GetRandomItemFromStringList(data.ExperienceDescriptionZh)

}

func (p *Profile) ToString() string {
	var result string

	if !p.generated {
		p.Generate()
	}

	switch p.Locale {
	case "zh_CN":
		result =
			`
			姓名：%s,
			性别：%s,
			年龄：%v,
			手机号码：%s,
			身份证号码：%s,
			家庭住址：%s,
			工作经历：%v,
			简介：%s
			`
		_gender := func() string {
			if p.Gender {
				return "男"
			} else {
				return "女"
			}
		}()
		_ex := p.FormatExperice()
		return fmt.Sprintf(result, p.Name, _gender, p.Age, p.MobileNumber, p.IdNumber, p.HomeAddress, _ex, p.Description)
	default:
		result = `
			姓名：%s,
			性别：%s,
			年龄：%v,
			手机号码：%s,
			身份证号码：%s,
			家庭住址：%s,
			工作经历：%v
			简介：%s
		`
		_gender := func() string {
			if p.Gender {
				return "男"
			} else {
				return "女"
			}
		}()
		_ex := p.FormatExperice()
		return fmt.Sprintf(result, p.Name, _gender, p.Age, p.MobileNumber, p.IdNumber, p.HomeAddress, _ex, p.Description)
	}
}

func (p *Profile) FormatExperice() string {
	var result string
	for _, v := range p.WorkExperiences {
		result += v.ToString() + "\n                                  "
	}

	return result
}

func determineWorkAge(currentAge int) int {
	var ages = []int{
		18, 19, 20, 21, 22, 23,
	}

	if currentAge > ages[len(ages)-1] {
		return ages[utils.Randn(len(ages))]
	} else {
		return ages[0]
	}
}
