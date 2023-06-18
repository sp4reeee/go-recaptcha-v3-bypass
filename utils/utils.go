package utils

import (
	"net/url"
	"regexp"

	catch "github.com/sp4reeee/go-recaptcha-v3-bypass/errors"
)

func Match_Data(data string, pattern string) string {

	regex, err := regexp.Compile(pattern)
	if err != nil {
		return "Error when compilating REGEX"
	}

	match := regex.FindStringSubmatch(data)

	if len(match) < 1 {
		return "No match found maybe use a proxy"
	}

	return match[1]
}

func Param_Extract(u string) (string, string, string) {
	parsed, err := url.Parse(u)
	if err != nil {
		return catch.ErrorParams(), "", ""
	}

	query := parsed.Query()
	v := query.Get("v")
	if v == "" {
		return "Parameters 'v' not found", "", ""
	}
	k := query.Get("k")
	if k == "" {
		return "Parameters 'k' not found", "", ""
	}
	co := query.Get("co")
	if co == "" {
		return "Parameters 'co' not found", "", ""
	}

	return v, k, co
}
