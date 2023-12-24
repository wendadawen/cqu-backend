package cas

import (
	"bytes"
	"cqu-backend/src/object"
	"cqu-backend/src/tool"
	"crypto/des"
	"crypto/tls"
	"encoding/base64"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"log"
	"net"
	"net/http"
	"regexp"
	"strings"
)

var executionRegExp = regexp.MustCompile(`name="execution" value="(.*?)"`)

// ‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾ 统一认证错误处理 ‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
var timeoutHandler tool.Handler = func(err interface{}) (bool, interface{}) {
	if err, ok := err.(net.Error); ok && err.Timeout() {
		return true, object.HttpTimeout
	}
	return false, err
}

var defaultHandler tool.Handler = func(err interface{}) (bool, interface{}) {
	log.Printf("[Cas] %+v", err.(error))
	// 默认的错误是 CasWebError
	return true, object.CasWebError
}

func NoCheckRedirect(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}
func (this *authentication) DoCheckRedirect(req *http.Request, via []*http.Request) error {
	if strings.Contains(req.URL.String(), "/enroll/token-index?code=") {
		code := ""
		compile := regexp.MustCompile("=(.*?)&state")
		submatch := compile.FindStringSubmatch(req.URL.RawQuery)
		if len(submatch) > 0 {
			code = submatch[1]
		}
		_clinet := resty.New()
		if this.client.HostURL == "https://my.cqu.edu.cn" && tool.ShouldProxy(tool.MY) { // 根据 yaml 配置代理
			_clinet.SetProxy(tool.ProxyUrl)
			//log.Printf("my.cqu use prooxy, proxy adress: %s", tool.ProxyUrl)
			_clinet.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
		}
		tokenReq := _clinet.R()
		tokenReq.SetHeader("Cookie", via[0].Header.Get("Cookie"))
		tokenReq.SetHeader("accessToken", "[object Object]")
		tokenReq.SetHeader("Authorization", "Basic ZW5yb2xsLXByb2Q6YXBwLWEtMTIzNA==")
		tokenReq.SetHeader("Content-Type", "application/x-www-form-urlencoded")
		res, err := tokenReq.SetFormData(map[string]string{
			"client_id":     "enroll-prod",
			"client_secret": "app-a-1234",
			"code":          code,
			"redirect_uri":  "https://my.cqu.edu.cn/enroll/token-index",
			"grant_type":    "authorization_code"}).Execute(http.MethodPost, this.client.HostURL+"/authserver/oauth/token")
		if err != nil {
			log.Printf("[Cas] %+v", err)
			return err
		}
		if strings.Contains(res.String(), "access_token") { // 成功
			this.token = gjson.Get(res.String(), "access_token").Str
		}
	}

	if strings.Contains(req.URL.String(), "/sam/token-index?code=") {
		code := ""
		compile := regexp.MustCompile("=(.*?)&state")
		submatch := compile.FindStringSubmatch(req.URL.RawQuery)
		if len(submatch) > 0 {
			code = submatch[1]
		}
		_clinet := resty.New()
		if this.client.HostURL == "https://my.cqu.edu.cn" && tool.ShouldProxy(tool.MY) { // 根据 yaml 配置代理
			_clinet.SetProxy(tool.ProxyUrl)
			//log.Printf("my.cqu use prooxy, proxy adress: %s", tool.ProxyUrl)
			_clinet.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
		}
		tokenReq := _clinet.R()
		tokenReq.SetHeader("Cookie", via[0].Header.Get("Cookie"))
		tokenReq.SetHeader("Authorization", "Basic c2FtLXByZDphcHAtYS0xMjM0")
		tokenReq.SetHeader("Content-Type", "application/x-www-form-urlencoded")
		res, err := tokenReq.SetFormData(map[string]string{
			"client_id":     "sam-prd",
			"client_secret": "app-a-1234",
			"code":          code,
			"redirect_uri":  "https://my.cqu.edu.cn/sam/token-index",
			"grant_type":    "authorization_code"}).Execute(http.MethodPost, this.client.HostURL+"/authserver/oauth/token")
		if err != nil {
			log.Printf("[Cas] %+v", err)
			return err
		}
		if strings.Contains(res.String(), "access_token") { // 成功
			this.token = gjson.Get(res.String(), "access_token").Str
		}
	}
	if strings.Contains(req.URL.Path, "jsessionid=") && req.URL.Host == "card.cqu.edu.cn:7280" {
		Cookie := via[0].Header.Get("Cookie")
		client := resty.New()
		client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
		jsessionReq := client.R().SetHeader("Cookie", Cookie)
		res, err := jsessionReq.Execute(http.MethodGet, req.URL.String())
		if err != nil {
			return err
		}
		compile := regexp.MustCompile(`id="ssoticketid" value="(.*?)"`)
		this.jSssoticketid = compile.FindStringSubmatch(res.String())[1]
	}

	return nil
}

func (this *authentication) Do(method string, urlPath string, payload map[string]string) (string, error) {

	//是否设置代理
	if this.client.HostURL == "https://my.cqu.edu.cn" && tool.ShouldProxy(tool.MY) { // 根据 yaml 配置代理
		this.client.SetProxy(tool.ProxyUrl)
		this.client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}

	// 中间重定向的内容
	if urlPath == "/login" && method == http.MethodPost {
		this.client.GetClient().CheckRedirect = NoCheckRedirect
	} else {
		this.client.GetClient().CheckRedirect = this.DoCheckRedirect
	}

	r := this.client.
		SetRetryCount(2).
		AddRetryCondition(func(r *resty.Response, err error) bool { return r.IsError() || err != nil }).
		R()

	//填充token
	if len(this.token) != 0 {
		r.SetHeader("Authorization", "Bearer "+this.token)
	}

	//填充payload form or body
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

	//请求
	res, err := r.Execute(method, urlPath)

	// 将获取到的cookie SESSION保存
	if urlPath == "/login" && method == http.MethodPost {
		Cookies := res.Cookies()
		for _, cookie := range Cookies {
			if strings.Contains(cookie.Name, "SOURCEID_TGC") {
				cookies := []*http.Cookie{
					{
						Name:  cookie.Name,
						Value: cookie.Value,
					},
				}
				this.client.SetCookies(cookies)
			}
			if strings.Contains(cookie.Name, "rg_objectid") {
				cookies := []*http.Cookie{
					&http.Cookie{
						Name:  cookie.Name,
						Value: cookie.Value,
					},
				}
				this.client.SetCookies(cookies)
			}
		}

		var location []string
		location, ok = res.Header()["Location"]
		if !ok {
			return "", object.CasAccountError
		}

		//如果是重定向
		resp, _ := this.client.R().Get(location[0])
		for _, cookie := range resp.Cookies() {
			if strings.Contains(cookie.Name, "SESSION") {
				cookies := []*http.Cookie{
					&http.Cookie{
						Name:  cookie.Name,
						Value: cookie.Value,
					},
				}
				this.client.SetCookies(cookies)
				break
			}
		}

		location, ok = res.Header()["Location"]
		if !ok {
			return "", object.CasAccountError
		}
		resp, _ = this.client.R().Get(location[0])

		for _, cookie := range resp.Cookies() {
			if strings.Contains(cookie.Name, "SESSION") {
				cookies := []*http.Cookie{
					{
						Name:  cookie.Name,
						Value: cookie.Value,
					},
				}

				this.client.SetCookies(cookies)
				break
			}
		}

		return "当前登录", nil
	}

	// TODO 错误处理，不会
	if err != nil {
		err = tool.NewChainHandler([]tool.Handler{timeoutHandler}).
			SetDefault(defaultHandler).Handle(err).(error)
		return "", err
	}

	return res.String(), nil
}

func ParseAndBuildLoginPayload(res string, id string, pwd string) (map[string]string, error) {
	payload := make(map[string]string)
	html, err := goquery.NewDocumentFromReader(strings.NewReader(res))
	if err != nil {
		return nil, err
	}
	payload["croypto"] = html.Find("#login-croypto").Text()
	payload["execution"] = html.Find("#login-page-flowkey").Text()
	payload["type"] = html.Find("#current-login-type").Text()
	payload["username"] = id
	payload["geolocation"] = ""
	payload["captcha_code"] = ""
	data, _ := base64.StdEncoding.DecodeString(payload["croypto"])
	da := DesECBEncrypt([]byte(pwd), data)
	payload["password"] = da
	//payload["lt"] = "11"
	payload["_eventId"] = "submit"
	return payload, nil
}

func DesECBEncrypt(data, key []byte) string {
	block, err := des.NewCipher(key)
	if err != nil {
		return ""
	}
	bs := block.BlockSize()
	data = Pkcs5Padding(data, bs)
	if len(data)%bs != 0 {
		return "need a multiple of the blocksize"
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return base64.StdEncoding.EncodeToString(out)
}

func Pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
