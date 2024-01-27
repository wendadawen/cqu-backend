package cas

import (
	"cqu-backend/src/spider"
	"cqu-backend/src/tool"
	"crypto/tls"
	"github.com/go-resty/resty/v2"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type Auth interface {
	spider.WithClientDo // 统一认证的Do函数可能给以下方面调用：card除了电费不用，my
	spider.WithLogin
	SetHost(host string)
	GetCookie(rawURL, name string) *http.Cookie
	GetJsSsoTicketId() string
	GetToken() string
	GetClient() *resty.Client
}

type auth struct {
	casId   string
	pwd     string
	client  *resty.Client
	isLogin bool

	token         string // my系统用
	jSssoticketid string // cara系统用
}

func NewAuth(stuId string, pwd string) Auth {
	client := resty.New()
	client.SetHostURL(authUrl)
	client.SetTimeout(5 * time.Second)
	if tool.ShouldProxy(tool.CAS) {
		client.SetProxy(tool.ProxyUrl)
	}
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	cookieJar, _ := cookiejar.New(nil)
	client.SetCookieJar(cookieJar)
	return &auth{
		client:  client,
		casId:   stuId,
		pwd:     pwd,
		isLogin: false,
	}
}
