package gofaker

import (
	fp "github.com/guchengxi1994/go_faker/providers"
	"github.com/guchengxi1994/go_faker/utils"
)

type UserAgent struct {
	PlatForm string
	Browser  string
}

func (u *UserAgent) Generate(allRandom bool) string {
	if allRandom {
		u.Browser = utils.GetRandomItemFromStringList([]string{
			"chrome", "firefox", "safari", "opera", "ie",
		})
		u.PlatForm = utils.GetRandomItemFromStringList([]string{
			"linux", "windows", "mac", "android", "ios",
		})
	}

	switch u.Browser {
	case "chrome":
		switch u.PlatForm {
		case "windows":
			return fp.Format(fp.Format_chrome_user_agent_windows, false)
		case "linux":
			return fp.Format(fp.Format_chrome_user_agent_linux, false)
		case "android":
			return fp.Format(fp.Format_chrome_user_agent_android, false)
		case "mac":
			return fp.Format(fp.Format_chrome_user_agent_mac, false)
		case "ios":
			return fp.Format(fp.Format_chrome_user_agent_ios, false)
		default:
			return fp.Format(fp.Format_chrome_user_agent_windows, false)
		}
	case "ie":
		return fp.Format(fp.Format_ie_user_agent, false)
	default:
		return ""
	}
}
