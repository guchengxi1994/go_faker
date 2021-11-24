package gofaker

import (
	"fmt"

	"github.com/guchengxi1994/go_faker/utils"
)

type Experience struct {
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	Institution string `json:"institution"`
	Content     string `json:"content"`
}

func (e *Experience) ToString() string {
	return fmt.Sprintf("%s~%s,%s,%s", e.StartTime, e.EndTime, e.Institution, e.Content)
}

/*
	times is the number join work

	years is the number form when to now
*/
func GenerateWorkExperiences(times, years int, locale string) []Experience {
	var result []Experience
	exList := utils.SplitTime(times, years)
	job := Job{
		Locale: locale,
	}
	com := Company{
		Locale: locale,
	}
	for _, v := range exList {
		e := Experience{
			StartTime:   v.From,
			EndTime:     v.To,
			Institution: com.Generate(),
			Content:     job.Generate(),
		}
		result = append(result, e)
	}

	return result
}

/*
	grade > 0 primary school 6 years,

	grade = 1 middle school  3 years,

	grade = 2 high school 3 years,

	grade = 3 university bachelor  4 years,

	grade = 4 university master 3 years,

	grade = 5 university doctor 4 years,

	grade = 6 others 3 years 中专

	grade = 7 others 5 years 大专

	September to July

	studyStartAge = one of [5,6,7,8]
*/
func GenerateSchoolExperience(grades int, currentAge int, locale string) ([]Experience, int) {
	studyStartAge := utils.GetRandomItemFromIntList([]int{5, 6, 7, 8})

	if currentAge < studyStartAge {
		return make([]Experience, 0), -1
	}

	var studyEndAges = []int{
		// grade = 2 or grade = 6
		studyStartAge + 6 + 3 + 3, // studyStartAge + primary + middle + high
		// grade = 7
		studyStartAge + 6 + 3 + 5, // studyStartAge + primary + others (大专)
		// grade = 3
		studyStartAge + 6 + 3 + 3 + 4, //studyStartAge + primary + middle + high + university
		// grade = 4
		studyStartAge + 6 + 3 + 3 + 4 + 3, //studyStartAge + primary + middle + high + university
		// grade = 5
		studyStartAge + 6 + 3 + 3 + 4 + 3 + 4, //studyStartAge + primary + middle + high + university
	}

	var result []Experience

	var studyEndAge int

	var _school = School{
		Locale: locale,
	}

	if grades == 2 || grades == 6 {
		studyEndAge = studyEndAges[0]
		ts := utils.SplitEducationYears([]int{6, 3, 3}, currentAge, studyStartAge)
		var _grade int
		for i, v := range ts {
			_grade = i
			if i == 0 {
				_school.Type = 0
			} else if i == 1 {
				_school.Type = 1
			} else {
				if grades == 2 {
					_school.Type = 2
				} else {
					_school.Type = 3
				}
			}
			e := Experience{
				StartTime:   v.From,
				EndTime:     v.To,
				Institution: _school.Generate(),
				Content:     education(locale, _grade),
			}
			result = append(result, e)
		}

	} else if grades == 7 {
		studyEndAge = studyEndAges[1]
		var _grade int
		ts := utils.SplitEducationYears([]int{6, 3, 5}, currentAge, studyStartAge)
		for i, v := range ts {
			if i == 0 {
				_school.Type = 0
				_grade = 0
			} else if i == 1 {
				_school.Type = 1
				_grade = 0
			} else {
				_school.Type = 3
				_grade = 7
			}
			e := Experience{
				StartTime:   v.From,
				EndTime:     v.To,
				Institution: _school.Generate(),
				Content:     education(locale, _grade),
			}
			result = append(result, e)
		}

	} else if grades == 3 {
		studyEndAge = studyEndAges[2]
		ts := utils.SplitEducationYears([]int{6, 3, 3, 4}, currentAge, studyStartAge)
		var _grade int
		for i, v := range ts {
			if i == 0 {
				_school.Type = 0
			} else if i == 1 {
				_school.Type = 1
			} else if i == 2 {
				_school.Type = 1
			} else {
				_school.Type = 2
				_grade = 3
			}
			e := Experience{
				StartTime:   v.From,
				EndTime:     v.To,
				Institution: _school.Generate(),
				Content:     education(locale, _grade),
			}
			result = append(result, e)
		}

	} else if grades == 4 {
		studyEndAge = studyEndAges[3]
		ts := utils.SplitEducationYears([]int{6, 3, 3, 4, 3}, currentAge, studyStartAge)
		var _grade int
		for i, v := range ts {
			if i == 0 {
				_school.Type = 0
			} else if i == 1 {
				_school.Type = 1
			} else if i == 2 {
				_school.Type = 1
			} else if i == 3 {
				_school.Type = 2
				_grade = 3
			} else {
				_school.Type = 2
				_grade = 4
			}
			e := Experience{
				StartTime:   v.From,
				EndTime:     v.To,
				Institution: _school.Generate(),
				Content:     education(locale, _grade),
			}
			result = append(result, e)
		}

	} else {
		studyEndAge = studyEndAges[4]
		var _grade int
		ts := utils.SplitEducationYears([]int{6, 3, 3, 4, 3, 4}, currentAge, studyStartAge)
		for i, v := range ts {
			if i == 0 {
				_school.Type = 0
			} else if i == 1 {
				_school.Type = 1
			} else if i == 2 {
				_school.Type = 1
			} else if i == 3 {
				_school.Type = 2
				_grade = 3
			} else if i == 4 {
				_school.Type = 2
				_grade = 4
			} else {
				_school.Type = 2
				_grade = 5
			}
			e := Experience{
				StartTime:   v.From,
				EndTime:     v.To,
				Institution: _school.Generate(),
				Content:     education(locale, _grade),
			}
			result = append(result, e)
		}
	}

	return result, currentAge - studyEndAge
}

func education(locale string, grades int) string {
	switch locale {
	case "zh_CN":
		if grades <= 2 {
			return "学生"
		} else if grades == 7 {
			return "专科"
		} else if grades == 3 {
			return "本科"
		} else if grades == 4 {
			return "硕士"
		} else {
			return "博士"
		}
	default:
		if grades <= 2 {
			return "学生"
		} else if grades == 7 {
			return "专科"
		} else if grades == 3 {
			return "本科"
		} else if grades == 4 {
			return "硕士"
		} else {
			return "博士"
		}
	}
}
