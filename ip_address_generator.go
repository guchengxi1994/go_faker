package gofaker

import (
	"regexp"
	"strconv"

	ipaddress "github.com/guchengxi1994/go_faker/providers/ipAddress"
	"github.com/guchengxi1994/go_faker/utils"
)

func IsIpV4Valid(ip string, containDNSs bool) bool {
	if !containDNSs {
		if utils.StringArrayContains(ipaddress.DNSs, ip) {
			return false
		}
	}

	for _, v := range ipaddress.ReservedIps {
		regx, _ := regexp.Compile(v)
		res := regx.FindAllString(ip, -1)
		if len(res) > 0 {
			return false
		}
	}

	return true
}

func GenerateIpV4Address(containDNSs bool) string {
	var result = ""
	for i := 0; i < 4; i++ {
		result += strconv.Itoa(utils.Randn(255))
		if i != 3 {
			result += "."
		}
	}

	if IsIpV4Valid(result, containDNSs) {
		return result
	} else {
		return GenerateIpV4Address(containDNSs)
	}
}

func GeneratePort(max int) string {
	if max <= 0 {
		max = 65535
	}
	return strconv.Itoa(utils.Randn(max))
}
