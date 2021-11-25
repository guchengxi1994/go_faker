package gofaker

import (
	lo "github.com/guchengxi1994/go_faker/providers/lorem"
	"github.com/guchengxi1994/go_faker/utils"
)

type Lorem struct {
	LikeUC    bool // only supports chinese
	Locale    string
	Type      int // 0 for paragraph,1 for short sentence(titles)
	MaxLength int // length for each sentence
	generated bool
	Content   string //generated string
}

func (lorem *Lorem) Generate() {
	lorem.generated = true
	switch lorem.Locale {
	case "zh_CN":
		if lorem.Type == 0 {
			for i := 0; i < 3; i++ {
				lorem.Content += "  " + generateSentence(lorem.MaxLength, lo.Lorem_words_zh)
				if i != 2 {
					lorem.Content += ","
				} else {
					lorem.Content += "。"
				}
			}

		} else {
			lorem.Content = generateSentence(lorem.MaxLength, lo.Lorem_words_zh)
		}
	default:
		if lorem.Type == 0 {
			for i := 0; i < 3; i++ {
				lorem.Content += "  " + generateSentence(lorem.MaxLength, lo.Lorem_words_zh)
				if i != 2 {
					lorem.Content += ","
				} else {
					lorem.Content += "。"
				}
			}
		} else {
			lorem.Content = generateSentence(lorem.MaxLength, lo.Lorem_words_zh)
		}
	}
}

func (lorem *Lorem) ToString() string {
	if lorem.LikeUC && lorem.Locale == "zh_CN" && lorem.Type == 1 {
		if !utils.StringEndsWith(lorem.Content, []rune("!！")) {
			var _tmp = []rune(lorem.Content)
			lorem.Content = string(_tmp[0:len(_tmp)-2]) + "!"
		}
		return "震惊！" + lorem.Content + "真相竟然是..."
	}
	return lorem.Content
}

func generateSentence(maxLength int, words []string) string {
	var result string
	for i := 0; i < maxLength; i++ {
		if maxLength >= 15 && i == maxLength/2 {
			result += ","
		}
		result += utils.GetRandomItemFromStringList(words)
	}

	return result
}
