package card

import (
	"cqu-backend/src/spider/cas"
	"log"
	"net/http"
)

type CasCard struct {
	auth    cas.Auth
	isLogin bool
}

// NewCardByCas 创建一个通过统一认证登录的实现类
func NewCardByCas(cardAccount CardAccount) (Card, error) {
	auth := cas.NewAuth(cardAccount.Account, cardAccount.Password)
	casCard := &CasCard{
		auth:    auth,
		isLogin: false,
	}
	cardTemplate := newCardTemplate(casCard)
	cardTemplate.account = cardAccount
	return cardTemplate, nil
}

// Login 使用统一认证方式登录
func (this *CasCard) Login() error {
	if this.isLogin == true {
		return nil
	}
	err := this.auth.Login()
	if err != nil {
		log.Printf("[CasCard Login Error] %+v\n", err)
		return err
	}
	card_url := cardAuth2card
	_, err = this.auth.Do(http.MethodGet, card_url, nil)
	if err != nil {
		log.Printf("[CasCard Login Error] %+v\n", err)
		return err
	}
	this.auth.SetHost(cardUrl)
	jsssoticketid := this.auth.GetJsSsoTicketId()
	_, err = this.auth.Do(http.MethodPost, "/cassyno/index", map[string]string{
		"errorcode":   "1",
		"continueurl": "http://card.cqu.edu.cn/cassyno/index",
		"ssoticketid": jsssoticketid,
	})
	return nil
}

func (this *CasCard) Do(method string, urlPath string, payload map[string]string) (string, error) {
	return this.auth.Do(method, urlPath, payload)
}
