package my

import (
	"cqu-backend/src/spider"
	"cqu-backend/src/spider/cas"
	"log"
	"net/http"
)

type CasMy struct {
	*myTemplate
	auth    cas.Auth
	isLogin bool
}

func NewMyByCas(account spider.SpiderAccount) (My, error) {
	auth := cas.NewAuth(account.Account, account.Password)
	casMy := &CasMy{
		auth:    auth,
		isLogin: false,
	}
	myTemplate := newMyTemplate(casMy)
	casMy.myTemplate = myTemplate
	return myTemplate, nil
}

func (this *CasMy) Login() error {
	if this.isLogin == true {
		return nil
	}
	err := this.auth.Login()
	if err != nil {
		log.Printf("[CasMy Login Error] %+v\n", err)
		return err
	}
	this.auth.SetHost(myHost)
	_, err = this.auth.Do(http.MethodGet, myAuthserverCas, nil)       // 获取必要的Cookies
	_, err = this.auth.Do(http.MethodGet, myAuthserverAuthorize, nil) // 获取token
	if err != nil {
		log.Printf("[CasMy Login Error] %+v\n", err)
		return err
	}
	this.isLogin = true
	return nil
}

func (this *CasMy) Do(method string, urlPath string, payload map[string]string) (string, error) {
	return this.auth.Do(method, urlPath, payload)
}
