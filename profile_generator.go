package gofaker

import (
	"fmt"

	"github.com/guchengxi1994/go_faker/data"
	"github.com/guchengxi1994/go_faker/utils"
)

type Profile struct {
	Name                 string `json:"name"`
	Gender               bool   `json:"gender"`
	IdNumber             string `json:"idNumber"`
	MobileNumber         string `json:"mobile"`
	Locale               string
	HomeAddress          string       `json:"homeAddress"`
	WorkExperiences      []Experience `json:"workExperiences"`
	Description          string       `json:"description"`
	generated            bool
	Age                  int          `json:"age"`
	EducationExperiences []Experience `json:"educationExperiences"` // 至少初中毕业，高中毕业，大学毕业，硕士，博士，大专
}

func (p *Profile) Generate(generators *Generators) {
	p.generated = true

	var pname PersonName

	if generators != nil && generators.Gpname != nil {
		pname = *generators.Gpname
		p.Gender = generators.Gpname.Gender
	} else {
		pname = PersonName{
			Locale: p.Locale,
			Gender: p.Gender,
		}
	}

	pname.Generate()

	p.Name = pname.ToString(true)

	var ssn SSN

	if generators != nil && generators.Gssn != nil {
		ssn = *generators.Gssn
	} else {
		ssn = SSN{
			Locale: p.Locale,
			MaxAge: 65,
			MinAge: 20,
			Gender: p.Gender,
		}
	}

	p.IdNumber = ssn.Generate()

	var mobile Mobile

	if generators != nil && generators.Gmobile != nil {
		mobile = *generators.Gmobile
	} else {
		mobile = Mobile{
			Locale: p.Locale,
		}
	}

	p.MobileNumber = mobile.Generate()

	var homeAddress Address

	if generators != nil && generators.Gaddress != nil {
		homeAddress = *generators.Gaddress
	} else {
		homeAddress = Address{
			Locole:       p.Locale,
			UseWeighting: false,
		}
	}

	p.HomeAddress = homeAddress.Address()

	p.Age = ssn.currentAge

	grades := utils.GetRandomItemFromIntList([]int{2, 3, 4, 5, 6, 7})

	eduExperience, rest := GenerateSchoolExperience(grades, p.Age, p.Locale)
	p.EducationExperiences = eduExperience

	if rest > 0 {
		var workTimes int
		workTimes = rest / 2
		if workTimes > 10 {
			workTimes = 10
		}
		p.WorkExperiences = GenerateWorkExperiences(workTimes, rest, p.Locale)
	}
	p.Description = utils.GetRandomItemFromStringList(data.ExperienceDescriptionZh)

}

func (p Profile) ToString() string {
	var result string

	if !p.generated {
		p.Generate(nil)
	}

	switch p.Locale {
	case "zh_CN":
		result =
			`
			姓名：%s
			性别：%s
			年龄：%v
			手机号码：%s
			身份证号码：%s
			家庭住址：%s
			学习经历：%v
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
		_ex_1 := p.FormatExperice(1)
		_ex_2 := p.FormatExperice(0)
		return fmt.Sprintf(result, p.Name, _gender, p.Age, p.MobileNumber, p.IdNumber, p.HomeAddress, _ex_1, _ex_2, p.Description)
	default:
		result = `
			姓名：%s
			性别：%s
			年龄：%v
			手机号码：%s
			身份证号码：%s
			家庭住址：%s
			学习经历：%v
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
		_ex_1 := p.FormatExperice(1)
		_ex_2 := p.FormatExperice(0)
		return fmt.Sprintf(result, p.Name, _gender, p.Age, p.MobileNumber, p.IdNumber, p.HomeAddress, _ex_1, _ex_2, p.Description)
	}
}

func (p *Profile) FormatExperice(t int) string {
	var result string
	if t == 0 {
		if len(p.WorkExperiences) == 0 {
			return "--"
		}

		for _, v := range p.WorkExperiences {
			result += v.ToString() + "\n                                  "
		}

		return result
	} else {
		if len(p.EducationExperiences) == 0 {
			return "--"
		}
		for _, v := range p.EducationExperiences {
			result += v.ToString() + "\n                                  "
		}

		return result
	}

}
