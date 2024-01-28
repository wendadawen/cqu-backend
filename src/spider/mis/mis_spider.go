package mis

import (
	"cqu-backend/src/bo"
	"cqu-backend/src/dao/model"
	"log"
	"net/http"
)

func (this *misTemplate) ClassSchedule() (*bo.ClassScheduleBo, error) {
	return nil, nil
}

func (this *misTemplate) ExamSchedule() (*bo.ExamScheduleBo, error) {
	err := this.Login()
	if err != nil {
		log.Printf("[MisSpider ExamSchedule Error] %+v\n", err)
		return nil, err
	}
	res, err := this.Do(http.MethodGet, misExam, nil)
	if err != nil {
		log.Printf("[MisSpider ExamSchedule Error] %+v\n", err)
		return nil, err
	}
	exams := extractExam(res)
	return exams, nil
}

func (this *misTemplate) Rank() (*model.Rank, error) {
	// 通过查数据库实现
	return nil, nil
}

func (this *misTemplate) AllScore() (*bo.MyScoreResultBo, error) {

	return nil, nil
}

func (this *misTemplate) CurrentScore() (*bo.MyScoreListBo, error) {

	return nil, nil
}

func (this *misTemplate) StudentInfo() (*bo.StudentInfoBo, error) {
	err := this.Login()
	if err != nil {
		log.Printf("[MisSpider StudentInfo Error] %+v\n", err)
		return nil, err
	}
	res, err := this.Do(http.MethodGet, misInfo, nil)
	if err != nil {
		log.Printf("[MisSpider StudentInfo Error] %+v\n", err)
		return nil, err
	}
	return extractStudentInfo(res), nil
}
