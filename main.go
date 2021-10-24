package main

import (
	"fmt"
	"generator/captcha"
	"generator/gen"
	"generator/mail"
	"generator/phone"
	"time"
)

func main() {
	email := mail.GenerateEmail()
	gen.GetCookies()
	finger := gen.GetFingerPrint()
	capkey := captcha.PostCaptcha("f5561ba9-8f1e-40ca-9b5b-a0b3f719ef34", "https://discord.com/")
	captchares := captcha.RecursiveCaptchaCheck(capkey)

	auth_token := gen.CreateToken(captchares, finger, email[0]+"@"+email[1]) //discord auth token

	time.Sleep(10 * time.Second)

	if len(auth_token) == 0 {
		return
	}
	phone_number, tzid := phone.GeneratePhone()
	fmt.Println(phone_number, tzid)
	gen.AddPhone(phone_number, finger, auth_token)
	time.Sleep(10 * time.Second)
	sms_code := phone.RecursiveCheck(tzid)
	sms_token := gen.GetPhoneToken(phone_number, sms_code, finger, auth_token)
	time.Sleep(10 * time.Second)

	gen.VerifyPhone(sms_token, finger, auth_token)
	time.Sleep(1 * time.Second)
	mail_id := mail.RecursiveCheck(email[0], email[1])
	verify_acc_token := gen.GetEmailToken(mail.GetTokenURLFromEmail(email[0], email[1], mail_id))

	capkey2 := captcha.PostCaptcha("f5561ba9-8f1e-40ca-9b5b-a0b3f719ef34", "https://discord.com/")
	cap_key2 := captcha.RecursiveCaptchaCheck(capkey2)

	gen.VerifyEmail(cap_key2, verify_acc_token, finger, auth_token)
	fmt.Println(email)
	fmt.Println(auth_token)
}
