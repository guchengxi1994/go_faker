package gofaker

import (
	p "github.com/guchengxi1994/go_faker/providers"
)

type Nickname struct {
	Prefix    string
	Suffix    string
	Gender    bool
	generated bool
}

func (nickname *Nickname) Generate() {
	nickname.generated = true
	if nickname.Gender {
		nickname.Prefix = p.Format(p.Format_nickname_male_pre, false)
	} else {
		nickname.Prefix = p.Format(p.Format_nickname_female_pre, false)
	}
	nickname.Suffix = p.Format(p.Format_nickname_common, false)
}

func (nickname *Nickname) ToString() string {
	if !nickname.generated {
		nickname.Generate()
	}
	return nickname.Prefix + nickname.Suffix
}
