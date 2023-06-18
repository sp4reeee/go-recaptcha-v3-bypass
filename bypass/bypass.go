package bypass

import (
	"fmt"

	"github.com/sp4reeee/go-recaptcha-v3-bypass/constants"
	"github.com/sp4reeee/go-recaptcha-v3-bypass/utils"
)

func Bypass(url string, proxy string) string {

	v, k, co := utils.Param_Extract(url)

	token := get_recaptcha_token(url, proxy)

	version := utils.Match_Data(url, constants.REGEX_CAPTCHA_VERSION)

	post_url := fmt.Sprintf(constants.BASE_URL_POST, version, k)

	post_data := fmt.Sprintf(constants.POST_DATA, v, token, k, co)

	resp_token := get_recaptcha_response(post_url, post_data, proxy)

	return resp_token
}
