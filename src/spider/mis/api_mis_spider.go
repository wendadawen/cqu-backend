package mis

import "cqu-backend/src/spider"

type Mis interface {
	spider.StudentUnionImplement
}

type misImplement interface {
	spider.WithLogin
	spider.WithClientDo
}

type misTemplate struct {
	misImplement misImplement
}

func newMisTemplate(misImplement misImplement) *misTemplate {
	return &misTemplate{misImplement: misImplement}
}

func (this *misTemplate) Login() error {
	return this.misImplement.Login()
}
func (this *misTemplate) Do(method string, urlPath string, payload map[string]string) (string, error) {
	return this.misImplement.Do(method, urlPath, payload)
}
