package extra_data

type GenMailStruct []string

type CheckMailStruct []struct {
	ID      int    `json:"id"`
	From    string `json:"from"`
	Subject string `json:"subject"`
	Date    string `json:"date"`
}

type RecEmail struct {
	ID          int           `json:"id"`
	From        string        `json:"from"`
	Subject     string        `json:"subject"`
	Date        string        `json:"date"`
	Attachments []interface{} `json:"attachments"`
	Body        string        `json:"body"`
	TextBody    string        `json:"textBody"`
	HTMLBody    string        `json:"htmlBody"`
}

type PostCaptchaStruct struct {
	Status  int    `json:"status"`
	Request string `json:"request"`
}

type PostVerifyAccStruct struct {
	CaptchaKey string `json:"captcha_key"`
	Token      string `json:"token"`
}

type GetBypassCap struct {
	C struct {
		Type string `json:"type"`
		Req  string `json:"req"`
	}
}

type CapReqThing struct {
	Header struct {
		Typ string `json:"typ"`
		Alg string `json:"alg"`
	} `json:"header"`
	Payload struct {
		S int    `json:"s"`
		T string `json:"t"`
		D string `json:"d"`
		L string `json:"l"`
		E int    `json:"e"`
	} `json:"payload"`
	Raw struct {
		Header    string `json:"header"`
		Payload   string `json:"payload"`
		Signature string `json:"signature"`
	} `json:"raw"`
}

type FingerPrint struct {
	Fingerprint string          `json:"fingerprint"`
	Assignments [][]interface{} `json:"assignments"`
}

type TokenStruct struct {
	Token string `json:"token"`
}

type GetPhoneStruct struct {
	Response int    `json:"response"`
	Tzid     int    `json:"tzid"`
	Number   string `json:"number"`
}

type SMSCodeStruct []struct {
	Country  int    `json:"country"`
	Sum      int    `json:"sum"`
	Service  string `json:"service"`
	Number   string `json:"number"`
	Response string `json:"response"`
	Tzid     int    `json:"tzid"`
	Time     int    `json:"time"`
	Msg      string `json:"msg"`
	Form     string `json:"form"`
}

type PhoneTokenStruct struct {
	Token string `json:"token"`
}
