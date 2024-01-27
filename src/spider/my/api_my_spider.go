package my

import "cqu-backend/src/spider"

type My interface {
	spider.StudentUnionImplement
}

type myImplement interface {
	spider.WithLogin
	spider.WithClientDo
}

type myTemplate struct {
	myImplement myImplement
}

func newMyTemplate(myImplement myImplement) *myTemplate {
	return &myTemplate{myImplement: myImplement}
}

func (this *myTemplate) Login() error {
	return this.myImplement.Login()
}
func (this *myTemplate) Do(method string, urlPath string, payload map[string]string) (string, error) {
	return this.myImplement.Do(method, urlPath, payload)
}
