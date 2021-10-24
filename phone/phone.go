package phone

import (
	"fmt"
	a "generator/struct"
	"strconv"
	"time"

	"github.com/imroc/req"
)

var r = req.New()

func GeneratePhone() (string, int) {
	res_body := a.GetPhoneStruct{}
	resp, _ := r.Get("https://onlinesim.ru/api/getNum.php?apikey=<ONLINESIMAPIKEYHERE>&service=discord&number=true&country=7")
	resp.ToJSON(&res_body)
	return res_body.Number, res_body.Tzid
}

func CheckSMSStatus(tzid int) string {
	res_body := a.SMSCodeStruct{}
	resp, _ := r.Get("https://onlinesim.ru/api/getState.php?apikey=<ONLINESIMAPIKEYHERE>&tzid=" + strconv.Itoa(tzid))
	resp.ToJSON(&res_body)
	return res_body[0].Response
}

func GetSMSToken(tzid int) string {
	res_body := a.SMSCodeStruct{}
	resp, _ := r.Get("https://onlinesim.ru/api/getState.php?apikey=<ONLINESIMAPIKEYHERE>&tzid=" + strconv.Itoa(tzid))
	resp.ToJSON(&res_body)
	return res_body[0].Msg
}

func RecursiveCheck(tzid int) string {
	for {
		time.Sleep(5 * time.Second)
		i := CheckSMSStatus(tzid)
		fmt.Println("Waiting for SMS")
		if i != "TZ_NUM_WAIT" {
			return GetSMSToken(tzid)
		}
	}
}
