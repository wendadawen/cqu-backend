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
	"net/http"
	"regexp"
	"strings"
)

var loginResponseMap = map[string]error{
	"登录成功": nil,
	"您提供的用户名或者密码有误":    object.CasAccountError,
	"当前存在其他用户使用同一帐号登录": object.CasContiuneError,
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
		Cookie := via[0].Header.Get("Cookie") // cardAuth2card
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

func checkLoginResult(res string, stringToError map[string]error) error {
	for key, err := range stringToError {
		if strings.Contains(res, key) {
			return err
		}
	}
	return object.UnknownError
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
