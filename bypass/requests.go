package bypass

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"strings"

	"github.com/sp4reeee/go-recaptcha-v3-bypass/constants"
	catch "github.com/sp4reeee/go-recaptcha-v3-bypass/errors"
	"github.com/sp4reeee/go-recaptcha-v3-bypass/utils"
)

var session = &http.Client{}

func get_recaptcha_token(u string, proxy string) string {

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return catch.ErrorToken()
	}

	req.Header.Set(constants.BASE_HEADERS_KEY, constants.BASE_HEADERS_VALUE)

	if proxy != "" {

		parsed, err := url.Parse(proxy)
		if err != nil {
			return catch.ErrorProxy()
		}

		transport := &http.Transport{
			Proxy: http.ProxyURL(parsed),
		}

		session.Transport = transport
	}

	resp, err := session.Do(req)
	if err != nil {
		return catch.ErrorProxy()
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return catch.ErrorToken()
	}

	token := utils.Match_Data(string(body), constants.REGEX_TOKEN_GET)

	return token
}

func get_recaptcha_response(u string, post_data string, proxy string) string {

	req, err := http.NewRequest("POST", u, strings.NewReader(post_data))
	if err != nil {
		return catch.ErrorResponse()
	}

	req.Header.Set(constants.BASE_HEADERS_KEY, constants.BASE_HEADERS_VALUE)

	if proxy != "" {

		parsed, err := url.Parse(proxy)
		if err != nil {
			return catch.ErrorProxy()
		}

		transport := &http.Transport{
			Proxy: http.ProxyURL(parsed),
		}

		session.Transport = transport
	}

	resp, err := session.Do(req)
	if err != nil {
		return catch.ErrorProxy()
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return catch.ErrorResponse()
	}

	token := utils.Match_Data(string(body), constants.REGEX_TOKEN_POST)

	return token
}
