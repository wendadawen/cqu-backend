package spider

import "cqu-backend/src/bo"

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

// 本科生和研究生都需要的功能
type StudentUnionImplement interface {
	ExamSchedule() (*bo.ExamScheduleBo, error)
}
