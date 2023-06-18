package utils

import (
	"net/url"
	"regexp"
)

func Match_Data(data string, pattern string) string {

	regex, err := regexp.Compile(pattern)
	if err != nil {
		return ""
	}

	match := regex.FindStringSubmatch(data)

	if len(match) < 1 {
		return ""
	}

	return match[1]
}

func Param_Extract(u string, param string) string {
	parsed, err := url.Parse(u)
	if err != nil {
		return ""
	}

	query := parsed.Query()
	paramValue := query.Get(param)

	return paramValue
}
