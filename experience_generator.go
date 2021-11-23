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
