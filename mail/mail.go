package mail

import (
	"fmt"
	a "generator/struct"
	"strings"
	"time"

	"github.com/imroc/req"
)

var r = req.New()

func GenerateEmail() []string {
	res_body := a.GenMailStruct{}
	res, _ := r.Get("https://www.1secmail.com/api/v1/?action=genRandomMailbox&count=100")
	res.ToJSON(&res_body)
	index, _ := findIndex(res_body, "xojxe.com")
	email := strings.Split(res_body[index], "@")
	return email
}

func CheckEmail(email, domain string) int {
	res_body := a.CheckMailStruct{}
	param := req.QueryParam{
		"login":  email,
		"domain": domain,
	}
	res, _ := r.Get("https://www.1secmail.com/api/v1/?action=getMessages", param)
	res.ToJSON(&res_body)
	if len(res_body) == 0 {
		return -1
	} else {
		return res_body[0].ID
	}
}

func RecursiveCheck(email, domain string) int {
	for {
		time.Sleep(5 * time.Second)
		i := CheckEmail(email, domain)
		fmt.Println("Waiting for email")
		if i != -1 {
			return i
		}
	}
}

func GetTokenURLFromEmail(email, domain string, id int) string {
	res_body := a.RecEmail{}
	param := req.QueryParam{
		"login":  email,
		"domain": domain,
		"id":     id,
	}
	res, _ := r.Get("https://www.1secmail.com/api/v1/?action=readMessage", param)
	res.ToJSON(&res_body)
	return strings.TrimSpace(res_body.TextBody[strings.Index(res_body.TextBody, "https://click.discord"):]) //get token at index, remove whitespace
}

func findIndex(array []string, key string) (int, error) {
	for i, ele := range array {
		if strings.Contains(ele, key) {
			return i, nil
		}
	}
	return -1, fmt.Errorf("nothing found")
}
