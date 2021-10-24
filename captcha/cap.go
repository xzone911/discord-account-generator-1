package captcha

import (
	"fmt"
	a "generator/struct"
	"time"

	"github.com/imroc/req"
)

func PostCaptcha(sitekey, siteurl string) string {
	res_body := a.PostCaptchaStruct{}
	url := fmt.Sprintf("http://2captcha.com/in.php?method=hcaptcha&sitekey=%v&pageurl=%v&key=<2CAPTCHAAPIKEYHERE>&json=1", sitekey, siteurl)
	res, _ := req.Post(url)
	res.ToJSON(&res_body)
	return res_body.Request
}

func CheckCapRes(id string) string {
	res_body := a.PostCaptchaStruct{}
	url := fmt.Sprintf("http://2captcha.com/res.php?key=<2CAPTCHAAPIKEYHERE>&id=%v&action=get&json=1", id)
	res, _ := req.Get(url)
	res.ToJSON(&res_body)
	return res_body.Request
}

func RecursiveCaptchaCheck(id string) string {
	for {
		time.Sleep(5 * time.Second)
		i := CheckCapRes(id)
		fmt.Println("Waiting for captcha")
		if i != "CAPCHA_NOT_READY" {
			return i
		}
	}
}
