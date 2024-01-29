package service

import (
	"cqu-backend/src/bo"
	"cqu-backend/src/spider"
	"cqu-backend/src/spider/mis"
	"cqu-backend/src/spider/my"
	"cqu-backend/src/tool"
	"log"
)

type ClassService struct{}

func NewClassService() *ClassService {
	return &ClassService{}
}

func (this *ClassService) ClassSchedule(StuId string, CasPwd string) (*bo.ClassScheduleBo, error) {
	account := spider.SpiderAccount{
		Account:  StuId,
		Password: CasPwd,
	}
	classSchedule := &bo.ClassScheduleBo{}
	if tool.CheckGraduate(StuId, CasPwd) {
		Mis, err := mis.NewMisByCas(account)
		if err != nil {
			log.Printf("[ClassService ClassSchedule Error] Account=%s\n", StuId)
			return nil, err
		}
		classSchedule, err = Mis.ClassSchedule()
		if err != nil {
			log.Printf("[ClassService ClassSchedule Error] Account=%s\n", StuId)
			return nil, err
		}
	} else {
		My, err := my.NewMyByCas(account)
		if err != nil {
			log.Printf("[ClassService ClassSchedule Error] Account=%s\n", StuId)
			return nil, err
		}
		classSchedule, err = My.ClassSchedule()
		if err != nil {
			log.Printf("[ClassService ClassSchedule Error] Account=%s\n", StuId)
			return nil, err
		}
	}
	return classSchedule, nil
}
