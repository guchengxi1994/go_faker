package gofaker

import (
	p "github.com/guchengxi1994/go_faker/providers"
)

type Job struct {
	Locale string
}

func (job *Job) Generate() string {
	switch job.Locale {
	case "zh_CN":
		return p.Format(p.Format_Zh_simple_job, false)
	default:
		return p.Format(p.Format_Zh_simple_job, false)
	}
}
