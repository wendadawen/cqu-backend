package spider

import (
	"cqu-backend/src/bo"
	"cqu-backend/src/model"
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

// 本科生和研究生都需要的功能
type StudentUnionImplement interface {
	ExamSchedule() (*bo.ExamScheduleBo, error)   // 考表
	Rank() (*model.Rank, error)                  // 绩点排名
	AllScore() (*bo.MyScoreResultBo, error)      // 全部成绩
	CurrentScore() (*bo.MyScoreListBo, error)    // 最近学期成绩
	ClassSchedule() (*bo.ClassScheduleBo, error) // 课表
	StudentInfo() (*bo.StudentInfoBo, error)     // 学生信息
}
