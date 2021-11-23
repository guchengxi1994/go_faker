package gofaker

import (
	j "github.com/guchengxi1994/go_faker/providers/jobs"
	"github.com/guchengxi1994/go_faker/utils"
)

type Job struct {
	Locale string
}

func (job *Job) Generate() string {
	switch job.Locale {
	case "zh_CN":
		return utils.GetRandomItemFromStringList(j.Job_zh)
	default:
		return utils.GetRandomItemFromStringList(j.Job_zh)
	}
}
