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
	spider.WithClientDo
	spider.WithLogin
	SetHost(host string)
	GetCookie(rawRUL string, name string) *http.Cookie
	GetJsSsoTicketId() string
	GetClient() *resty.Client
}

type authentication struct {
	casId   string
	IdNo    string
	pwd     string
	client  *resty.Client
	cookies *cookiejar.Jar
	isLogin bool

	token         string
	jSssoticketid string
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
	return &authentication{
		client:  client,
		casId:   stuId,
		pwd:     pwd,
		isLogin: false,
		cookies: cookieJar,
	}
}
