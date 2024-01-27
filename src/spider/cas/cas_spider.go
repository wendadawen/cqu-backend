package cas

import (
	"cqu-backend/src/object"
	"cqu-backend/src/tool"
	"github.com/go-resty/resty/v2"
	"log"
	"net"
	"net/http"
	"net/url"
)

func (this *auth) GetCookie(rawURL, name string) *http.Cookie {
	parse, _ := url.Parse(rawURL)
	cookies := this.client.GetClient().Jar.Cookies(parse)
	for _, cookie := range cookies {
		if cookie.Name == name {
			return cookie
		}
	}
	return nil
}
func (this *auth) GetJsSsoTicketId() string {
	return this.jSssoticketid
}
func (this *auth) GetToken() string {
	return this.token
}
func (this *auth) GetClient() *resty.Client {
	return this.client
}
func (this *auth) SetHost(host string) {
	this.client.SetHostURL(host)
}
func (this *auth) Login() error {
	if this.isLogin == true {
		return nil
	}
	res, err := this.Do(http.MethodGet, loginUrl, nil) // 请求表单数据
	if err != nil {
		log.Printf("[CasSpider Login Error] %+v\n", err)
		return err
	}
	payload, err := ParseAndBuildLoginPayload(res, this.casId, this.pwd)
	if err != nil {
		log.Printf("[CasSpider Login Error] %+v\n", err)
		return err
	}
	res, err = this.Do(http.MethodPost, loginUrl, payload) // 登录
	if err != nil {
		log.Printf("[CasSpider Login Error] %+v\n", err)
		return err
	}
	err = checkLoginResult(res, loginResponseMap)
	if err != nil {
		log.Printf("[CasSpider Login Error] %+v\n", err)
		return err
	}
	this.isLogin = true
	return nil
}
func (this *auth) Do(method string, urlPath string, payload map[string]string) (string, error) {

	if this.client.HostURL == "https://my.cqu.edu.cn" && tool.ShouldProxy(tool.MY) { // 根据 yaml 配置代理
		this.client.SetProxy(tool.ProxyUrl)
	}

	if urlPath == "/login" && method == http.MethodPost {
		this.client.GetClient().CheckRedirect = NoCheckRedirect
	} else {
		this.client.GetClient().CheckRedirect = this.DoCheckRedirect
	}

	r := this.client.
		SetRetryCount(2).
		AddRetryCondition(func(r *resty.Response, err error) bool { return r.IsError() || err != nil }).
		R()

	if len(this.token) != 0 {
		r.SetHeader("Authorization", "Bearer "+this.token)
	}

	jsonArray, ok := payload["jsonArray"]
	if ok {
		r = r.SetHeaders(map[string]string{
			"Content-Type":     "application/json",
			"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.81 Safari/537.36 Edg/104.0.1293.54",
			"X-Requested-With": "XMLHttpRequest",
		}).
			SetBody([]byte(jsonArray))
	} else {
		r = r.SetHeaders(map[string]string{
			"Content-Type":     "application/x-www-form-urlencoded",
			"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.81 Safari/537.36 Edg/104.0.1293.54",
			"X-Requested-With": "XMLHttpRequest",
		}).
			SetFormData(payload)
	}

	res, err := r.Execute(method, urlPath)

	if err != nil {
		if err, ok := err.(net.Error); ok && err.Timeout() {
			return "", object.HttpTimeout
		} else {
			return "", object.CasWebError
		}
	}

	if urlPath == "/login" && method == http.MethodPost {
		if _, ok = res.Header()["Location"]; !ok {
			return "", object.CasAccountError
		}
		return "登录成功", nil
	}

	return res.String(), nil
}
