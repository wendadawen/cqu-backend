package service

import (
	"cqu-backend/src/bo"
	"cqu-backend/src/spider"
	"cqu-backend/src/spider/mis"
	"cqu-backend/src/spider/my"
	"cqu-backend/src/tool"
	"log"
)

type ExamService struct{}

func NewExamService() *ExamService {
	return &ExamService{}
}

func (this *ExamService) ExamSchedule(StuId, CasPwd string) (*bo.ExamScheduleBo, error) {
	examSchedule := &bo.ExamScheduleBo{}
	account := spider.SpiderAccount{
		Account:  StuId,
		Password: CasPwd,
	}
	// 研究生就用Mis系统，本科生就用My系统
	if tool.CheckGraduate(StuId, CasPwd) {
		Mis, err := mis.NewMisByCas(account)
		if err != nil {
			log.Printf("[ExamService ExamSchedule Error] Account=%s\n", StuId)
			return nil, err
		}
		examSchedule, err = Mis.ExamSchedule()
	} else {
		My, err := my.NewMyByCas(account)
		if err != nil {
			log.Printf("[ExamService ExamSchedule Error] Account=%s\n", StuId)
			return nil, err
		}
		examSchedule, err = My.ExamSchedule()
	}
	return examSchedule, nil
}
