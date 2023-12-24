package spider

import (
	"net/http"
)

const (
	SpiderUserAgent  = "Mozilla/5.0 (compatible; Baiduspider/2.0;+http://www.baidu.com/search/spider.html）"
	FirefoxUserAgent = "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:34.0) Gecko/20100101 Firefox/34.0"
)

type WithLogin interface {
	Login() error
}

type WithClientDo interface {
	Do(method string, urlPath string, payload map[string]string) (string, error)
}

type SpiderAccount struct {
	Account  string
	Password string
}

type CookiesMap = map[string]*http.Cookie

func UpdateCookies(cookies CookiesMap, c []*http.Cookie) {
	if cookies == nil || c == nil {
		return
	}
	if len(c) > 0 {
		for _, cookie := range c {
			cookies[cookie.Name] = cookie
		}
	}
}
