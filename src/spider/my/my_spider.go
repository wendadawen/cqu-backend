package my

import (
	"cqu-backend/src/bo"
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
