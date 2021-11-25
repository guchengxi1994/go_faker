package gofaker

import (
	ni "github.com/guchengxi1994/go_faker/providers/nickname"
	"github.com/guchengxi1994/go_faker/utils"
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
		nickname.Prefix = utils.GetRandomItemFromStringList(ni.Prefix_nickname_male)
	} else {
		nickname.Prefix = utils.GetRandomItemFromStringList(ni.Prefix_nickname_female)
	}
	nickname.Suffix = utils.GetRandomItemFromStringList(ni.Common_nicknames)
}

func (nickname *Nickname) ToString() string {
	if !nickname.generated {
		nickname.Generate()
	}
	return nickname.Prefix + nickname.Suffix
}
