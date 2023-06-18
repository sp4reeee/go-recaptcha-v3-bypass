package bypass

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sp4reeee/go-recaptcha-v3-bypass/constants"
	"github.com/sp4reeee/go-recaptcha-v3-bypass/utils"
)

var session = &http.Client{}

func get_recaptcha_token(url string) string {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ""
	}

	req.Header.Set(constants.BASE_HEADERS_KEY, constants.BASE_HEADERS_VALUE)

	resp, err := session.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	token := utils.Match_Data(string(body), constants.REGEX_TOKEN_GET)

	return token
}

func get_recaptcha_response(url string, post_data string) string {

	req, err := http.NewRequest("POST", url, strings.NewReader(post_data))
	if err != nil {
		return ""
	}

	req.Header.Set(constants.BASE_HEADERS_KEY, constants.BASE_HEADERS_VALUE)

	resp, err := session.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	token := utils.Match_Data(string(body), constants.REGEX_TOKEN_POST)

	return token
}
