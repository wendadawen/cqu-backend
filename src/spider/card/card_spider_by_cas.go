package card

import (
	"cqu-backend/src/object"
	"cqu-backend/src/spider"
	"cqu-backend/src/spider/cas"
	"log"
	"net/http"
)

type CasCard struct {
	auth    cas.Auth
	cookies spider.CookiesMap
	isLogin bool
}

// NewCardByCas 创建一个通过统一认证登录一卡通的实现类
func NewCardByCas(cardAccount CardAccount) (Card, error) {
	auth := cas.NewAuth(cardAccount.Account, cardAccount.Password)
	casCard := &CasCard{
		auth:    auth,
		isLogin: false,
		cookies: spider.CookiesMap{},
	}
	cardTemplate := newCardTemplate(casCard)
	cardTemplate.cardAccount = cardAccount
	return cardTemplate, nil
}

// Login 使用一卡通方式登录校园卡
func (this *CasCard) Login() error {
	err := this.auth.Login()
	if err != nil {
		log.Printf("[CasCard Login authLogin Error] %+v\n", err)
		return err
	}
	card_url := "https://sso.cqu.edu.cn/login?service=http:%2F%2Fcard.cqu.edu.cn:7280%2Fias%2Fprelogin%3Fsysid%3DFWDT%26continueurl%3Dhttp%253A%252F%252Fcard.cqu.edu.cn%252Fcassyno%252Findex"
	_, err = this.auth.Do(http.MethodGet, card_url, nil)
	if err != nil {
		log.Printf("[CasCard Login authDo1 Error] %+v\n", err)
		return err
	}
	this.auth.SetHost(cardUrl)
	jsssoticketid := this.auth.GetJsSsoTicketId()
	_, err = this.auth.Do(http.MethodPost, "/cassyno/index", map[string]string{
		"errorcode":   "1",
		"continueurl": "http://card.cqu.edu.cn/cassyno/index",
		"ssoticketid": jsssoticketid,
	})
	if err != nil {
		log.Printf("[CasCard Login autoDo2 Error] %+v\n", err)
		return err
	}
	cookie := this.auth.GetCookie(cardUrl, "hallticket")
	if cookie == nil {
		log.Printf("[CasCard Login GetCookie Error] %+v\n", object.CardCookieError)
		return object.CardCookieError
	}
	spider.UpdateCookies(this.cookies, []*http.Cookie{cookie})
	return nil
}

func (this *CasCard) Do(method string, urlPath string, payload map[string]string) (string, error) {
	return this.auth.Do(method, urlPath, payload)
}
