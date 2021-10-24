package gen

import (
	"fmt"
	a "generator/struct"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
)

var client *resty.Client = resty.New()

func GetCookies() {
	// client.R().
	// 	Get("http://localhost:3000")

	cookies := []*http.Cookie{
		&http.Cookie{
			Name:     "__sdcfduid",
			Value:    "sdc cookie here, im too lazy",
			HttpOnly: true,
		},
		&http.Cookie{
			Name:     "__dcfduid",
			Value:    "dcf cookie here, im too lazy",
			HttpOnly: true,
		},
	}
	client.SetCookies(cookies)
}

func GetFingerPrint() string {
	res := a.FingerPrint{}
	client.R().
		SetResult(&res).
		SetHeaders(map[string]string{
			"Host":                 "discord.com",
			"sec-ch-ua":            `"Chromium";v="94", "Google Chrome";v="94", ";Not A Brand";v="99"`,
			"x-super-properties":   GetSuperProperties(),
			"x-context-properties": "eyJsb2NhdGlvbiI6IkxvZ2luIn0=",
			"x-debug-options":      "bugReporterEnabled",
			"accept-language":      "en-GB",
			"sec-ch-ua-mobile":     "?0",
			"authorization":        "undefined",
			"user-agent":           "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.81 Safari/537.36",
			"sec-ch-ua-platform":   `"macOS"`,
			"accept":               "*/*",
			"sec-fetch-site":       "same-origin",
			"sec-fetch-mode":       "cors",
			"sec-fetch-dest":       "empty",
			"referer":              "https://discord.com/login",
		}).
		Get("https://discordapp.com/api/v9/experiments")

	return res.Fingerprint
}

func GetSuperProperties() string {
	return "eyJvcyI6Ik1hYyBPUyBYIiwiYnJvd3NlciI6IkNocm9tZSIsImRldmljZSI6IiIsInN5c3RlbV9sb2NhbGUiOiJlbi1HQiIsImJyb3dzZXJfdXNlcl9hZ2VudCI6Ik1vemlsbGEvNS4wIChNYWNpbnRvc2g7IEludGVsIE1hYyBPUyBYIDEwXzE0XzYpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS85NC4wLjQ2MDYuODEgU2FmYXJpLzUzNy4zNiIsImJyb3dzZXJfdmVyc2lvbiI6Ijk0LjAuNDYwNi44MSIsIm9zX3ZlcnNpb24iOiIxMC4xNC42IiwicmVmZXJyZXIiOiIiLCJyZWZlcnJpbmdfZG9tYWluIjoiIiwicmVmZXJyZXJfY3VycmVudCI6IiIsInJlZmVycmluZ19kb21haW5fY3VycmVudCI6IiIsInJlbGVhc2VfY2hhbm5lbCI6InN0YWJsZSIsImNsaWVudF9idWlsZF9udW1iZXIiOjEwMjExMywiY2xpZW50X2V2ZW50X3NvdXJjZSI6bnVsbH0="
}

func CreateToken(captcha_key, fingerprint, email string) string {
	test, _ := client.R().
		SetHeaders(map[string]string{
			"Host":               "discord.com",
			"sec-ch-ua":          `"Chromium";v="94", "Google Chrome";v="94", ";Not A Brand";v="99"`,
			"x-super-properties": GetSuperProperties(),
			"x-fingerprint":      fingerprint,
			"x-debug-options":    "bugReporterEnabled",
			"accept-language":    "en-GB",
			"sec-ch-ua-mobile":   "?0",
			"authorization":      "undefined",
			"user-agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.81 Safari/537.36",
			"content-type":       "application/json",
			"sec-ch-ua-platform": `"macOS"`,
			"accept":             "*/*",
			"origin":             "https://discord.com",
			"sec-fetch-site":     "same-origin",
			"sec-fetch-mode":     "cors",
			"sec-fetch-dest":     "empty",
			"referer":            "https://discord.com/register",
		}).
		SetBody(map[string]interface{}{"fingerprint": fingerprint, "email": email, "username": "budfgjahdikuwad", "password": "Hokejista2003_", "invite": nil, "consent": true, "date_of_birth": "1999-11-01", "gift_code_sku_id": nil, "captcha_key": nil, "promotional_email_opt_in": false}).
		Post("https://discord.com/api/v9/auth/register")

	fmt.Println(test)

	res := a.TokenStruct{}
	client.R().
		SetHeaders(map[string]string{
			"Host":               "discord.com",
			"sec-ch-ua":          `"Chromium";v="94", "Google Chrome";v="94", ";Not A Brand";v="99"`,
			"x-super-properties": GetSuperProperties(),
			"x-fingerprint":      fingerprint,
			"x-debug-options":    "bugReporterEnabled",
			"accept-language":    "en-GB",
			"sec-ch-ua-mobile":   "?0",
			"authorization":      "undefined",
			"user-agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.81 Safari/537.36",
			"content-type":       "application/json",
			"sec-ch-ua-platform": `"macOS"`,
			"accept":             "*/*",
			"origin":             "https://discord.com",
			"sec-fetch-site":     "same-origin",
			"sec-fetch-mode":     "cors",
			"sec-fetch-dest":     "empty",
			"referer":            "https://discord.com/register",
		}).
		SetResult(&res).
		SetBody(map[string]interface{}{"fingerprint": fingerprint, "email": email, "username": "budfgjahdikuwad", "password": "Hokejista2003_", "invite": nil, "consent": true, "date_of_birth": "1999-11-01", "gift_code_sku_id": nil, "captcha_key": captcha_key, "promotional_email_opt_in": false}).
		Post("https://discord.com/api/v9/auth/register")

	fmt.Println(res)

	return res.Token
}

func VerifyEmail(captcha_key, email_key, fingerprint, token string) {
	resp, _ := client.R().
		SetHeaders(map[string]string{
			"Host":               "discord.com",
			"sec-ch-ua":          `"Chromium";v="94", "Google Chrome";v="94", ";Not A Brand";v="99"`,
			"x-super-properties": GetSuperProperties(),
			"x-debug-options":    "bugReporterEnabled",
			"accept-language":    "en-US",
			"sec-ch-ua-mobile":   "?0",
			"authorization":      token,
			"content-type":       "application/json",
			"user-agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.81 Safari/537.36",
			"sec-ch-ua-platform": `"macOS"`,
			"accept":             "*/*",
			"origin":             "https://discord.com",
			"sec-fetch-site":     "same-origin",
			"sec-fetch-mode":     "cors",
			"sec-fetch-dest":     "empty",
			"referer":            "https://discord.com/verify",
		}).
		SetBody(map[string]interface{}{"captcha_key": captcha_key, "token": email_key}).
		Post("https://discord.com/api/v9/auth/verify")

	fmt.Println(resp)
}

func GetEmailToken(url string) string {
	client.SetRedirectPolicy(resty.FlexibleRedirectPolicy(15))
	res, _ := client.R().
		Get(url)

	return strings.Split(res.RawResponse.Request.URL.String(), "=")[1]

}

func AddPhone(phone_number, fingerprint, token string) {
	resp, _ := client.R().
		SetHeaders(map[string]string{
			"Host":               "discord.com",
			"sec-ch-ua":          `"Chromium";v="94", "Google Chrome";v="94", ";Not A Brand";v="99"`,
			"x-super-properties": GetSuperProperties(),
			"x-debug-options":    "bugReporterEnabled",
			"accept-language":    "en-GB",
			"sec-ch-ua-mobile":   "?0",
			"authorization":      token,
			"content-type":       "application/json",
			"user-agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.81 Safari/537.36",
			"sec-ch-ua-platform": `"macOS"`,
			"accept":             "*/*",
			"origin":             "https://discord.com",
			"sec-fetch-site":     "same-origin",
			"sec-fetch-mode":     "cors",
			"sec-fetch-dest":     "empty",
			"referer":            "https://discord.com/channels/@me",
		}).
		SetBody(map[string]interface{}{"phone": phone_number}).
		Post("https://discord.com/api/v9/users/@me/phone")
	fmt.Println("resp")
	fmt.Println(resp)
}

func GetPhoneToken(phone_number, sms_code, fingerprint, token string) string {
	res := a.PhoneTokenStruct{}
	client.R().
		SetHeaders(map[string]string{
			"Host":               "discord.com",
			"sec-ch-ua":          `"Chromium";v="94", "Google Chrome";v="94", ";Not A Brand";v="99"`,
			"x-super-properties": GetSuperProperties(),
			"x-debug-options":    "bugReporterEnabled",
			"accept-language":    "en-GB",
			"sec-ch-ua-mobile":   "?0",
			"authorization":      token,
			"content-type":       "application/json",
			"user-agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.81 Safari/537.36",
			"sec-ch-ua-platform": `"macOS"`,
			"accept":             "*/*",
			"origin":             "https://discord.com",
			"sec-fetch-site":     "same-origin",
			"sec-fetch-mode":     "cors",
			"sec-fetch-dest":     "empty",
			"referer":            "https://discord.com/channels/@me",
		}).
		SetResult(&res).
		SetBody(map[string]interface{}{"phone": phone_number, "code": sms_code}).
		Post("https://discord.com/api/v9/phone-verifications/verify")

	fmt.Println(res)
	return res.Token
}

func VerifyPhone(sms_token, fingerprint, token string) {
	// res := a.PhoneTokenStruct{}
	resp, _ := client.R().
		SetHeaders(map[string]string{
			"Host":               "discord.com",
			"sec-ch-ua":          `"Chromium";v="94", "Google Chrome";v="94", ";Not A Brand";v="99"`,
			"x-super-properties": GetSuperProperties(),
			"x-debug-options":    "bugReporterEnabled",
			"accept-language":    "en-GB",
			"sec-ch-ua-mobile":   "?0",
			"authorization":      token,
			"content-type":       "application/json",
			"user-agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.81 Safari/537.36",
			"sec-ch-ua-platform": `"macOS"`,
			"accept":             "*/*",
			"origin":             "https://discord.com",
			"sec-fetch-site":     "same-origin",
			"sec-fetch-mode":     "cors",
			"sec-fetch-dest":     "empty",
			"referer":            "https://discord.com/channels/@me",
		}).
		// SetResult(&res).
		SetBody(map[string]interface{}{"phone_token": sms_token, "password": "Hokejista2003_"}).
		Post("https://discord.com/api/v9/users/@me/phone")

	fmt.Println(resp)
	// return res.Token
}

/*

	payload := a.PosScienceStruct{
		Token: token,
		Events: []a.InsideEvent{
			{
				Type: "network_action_user_register",
				Properties: a.InsideProp{
					ClientTrackTimestamp:        time - 20000,
					StatusCode:                  201,
					URL:                         "/auth/register",
					RequestMethod:               "post",
					InviteCode:                  nil,
					PromotionalEmailOptIn:       false,
					PromotionalEmailPreChecked:  false,
					LocationSection:             "impression_user_registration",
					AccessibilitySupportEnabled: false,
					AccessibilityFeatures:       256,
					ClientUUID:                  uuid,
					ClientSendTimestamp:         time,
				},
			},
			{
				Type: "nuo_transition",
				Properties: a.InsideProp{
					ClientTrackTimestamp:        time,
					FlowType:                    "organic_registration",
					FromStep:                    nil,
					ToStep:                      "nuf_started",
					SecsOnFromStep:              0,
					AccessibilitySupportEnabled: false,
					AccessibilityFeatures:       256,
					ClientUUID:                  uuid,
					ClientSendTimestamp:         time,
				},
			},
		},
	}
*/
