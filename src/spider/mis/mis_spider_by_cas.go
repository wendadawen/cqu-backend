package mis

import (
	"cqu-backend/src/spider"
	"cqu-backend/src/spider/cas"
	"log"
)

type CasMis struct {
	*misTemplate
	auth    cas.Auth
	isLogin bool
}

// NewMisByCas 创建一个通过统一认证登录的实现类
func NewMisByCas(account spider.SpiderAccount) (Mis, error) {
	auth := cas.NewAuth(account.Account, account.Password)
	casMis := &CasMis{
		auth:    auth,
		isLogin: false,
	}
	misTemplate := newMisTemplate(casMis)
	casMis.misTemplate = misTemplate
	return misTemplate, nil
}

func (this *CasMis) Login() error {
	if this.isLogin == true {
		return nil
	}
	err := this.auth.Login()
	if err != nil {
		log.Printf("[CasMis Login Error] %+v\n", err)
		return err
	}
	// TODO
	this.isLogin = true
	return nil
}

func (this *CasMis) Do(method string, urlPath string, payload map[string]string) (string, error) {
	return this.auth.Do(method, urlPath, payload)
}
