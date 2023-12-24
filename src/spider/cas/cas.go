package cas

import (
	"cqu-backend/src/object"
	"github.com/go-resty/resty/v2"
	"net/http"
	"net/url"
	"strings"
)

func (this *authentication) GetCookie(rawURL string, name string) *http.Cookie {
	parse, _ := url.Parse(rawURL)
	cookies := this.cookies.Cookies(parse)
	for _, cookie := range cookies {
		if cookie.Name == name {
			return cookie
		}
	}
	return nil
}
func (this *authentication) GetJsSsoTicketId() string {
	return this.jSssoticketid
}
func (this *authentication) GetClient() *resty.Client {
	return this.client
}
func (this *authentication) SetHost(host string) {
	this.client.SetHostURL(host)
}

func (this *authentication) Login() error {
	if this.isLogin == true {
		return nil
	}
	res, _ := this.Do(http.MethodGet, validCodeUrl+this.casId, nil) // 验证码
	if strings.Contains(res, "true") {
		return object.CasValidcode
	}
	res, err := this.Do(http.MethodGet, loginUrl, nil) // 请求表单数据
	if err != nil {
		return err
	}
	payload, err := ParseAndBuildLoginPayload(res, this.casId, this.pwd)
	if err != nil {
		return err
	}
	res, err = this.Do(http.MethodPost, loginUrl, payload) // 登录
	if err != nil {
		//fmt.Printf("%+v\n", err)
		return err
	}
	err = checkLoginResult(res, loginResponseMap)
	/*	if xerrors.Is(err, object.CasContiuneError) {
		execution := tool.FindFirstSubMatch(executionRegExp, res)
		res, err = this.Do(http.MethodPost, loginUrl, map[string]string{
			"execution": execution,
			"_eventId":  "continue",
		})
	}*/
	if err != nil {
		return err
	}
	this.isLogin = true
	return nil
}

// ‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾ 统一认证登录返回信息处理 ‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
var loginResponseMap = map[string]error{
	"当前登录": nil,
	"您提供的用户名或者密码有误":    object.CasAccountError,
	"当前存在其他用户使用同一帐号登录": object.CasContiuneError,
}

func checkLoginResult(res string, stringToError map[string]error) error {
	for key, err := range stringToError {
		if strings.Contains(res, key) {
			return err
		}
	}
	//log.Printf("[Cas Login] Unkown Login result \n%s", res)
	return object.UnknownError
}
