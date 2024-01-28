package my

import (
	"cqu-backend/src/bo"
	"cqu-backend/src/dao/model"
	"log"
	"net/http"
)

func (this *myTemplate) ExamSchedule() (*bo.ExamScheduleBo, error) {
	err := this.Login()
	if err != nil {
		log.Printf("[MySpider ExamSchedule Error] %+v\n", err)
		return nil, err
	}
	res, err := this.Do(http.MethodGet, myExam, nil)
	if err != nil {
		log.Printf("[MySpider ExamSchedule Error] %+v\n", err)
		return nil, err
	}
	exams := exactExams(res)
	return exams, nil
}

func (this *myTemplate) Rank() (*model.Rank, error) {
	err := this.Login()
	if err != nil {
		log.Printf("[MySpider Rank Error] %+v\n", err)
		return nil, err
	}
	res, err := this.Do(http.MethodGet, myRank, nil)
	if err != nil {
		log.Printf("[MySpider Rank Error] %+v\n", err)
		return nil, err
	}
	return exacRank(res), nil
}
