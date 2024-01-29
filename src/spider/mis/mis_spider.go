package mis

import (
	"cqu-backend/src/bo"
	"log"
	"net/http"
)

func (this *misTemplate) ClassSchedule() (*bo.ClassScheduleBo, error) {
	err := this.Login()
	if err != nil {
		log.Printf("[MisSpider ClassSchedule Error] %+v\n", err)
		return nil, err
	}
	info, err := this.StudentInfo()
	res, err := this.Do(http.MethodPost, misClass+info.Uid, nil)
	if err != nil {
		log.Printf("[MisSpider ClassSchedule Error] %+v\n", err)
		return nil, err
	}
	return extractClassSchedule(res), nil
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

func (this *misTemplate) Rank() (*bo.Rank, error) {
	// 通过查数据库实现
	return nil, nil
}

func (this *misTemplate) AllScore() (*bo.MyScoreResultBo, error) {
	err := this.Login()
	if err != nil {
		log.Printf("[MisSpider AllScore Error] %+v\n", err)
		return nil, err
	}
	res, err := this.Do(http.MethodGet, misScore, nil)
	if err != nil {
		log.Printf("[MisSpider AllScore Error] %+v\n", err)
		return nil, err
	}
	return extractAllScore(res), nil
}

func (this *misTemplate) CurrentScore() (*bo.MyScoreListBo, error) {
	err := this.Login()
	if err != nil {
		log.Printf("[MisSpider CurrentScore Error] %+v\n", err)
		return nil, err
	}
	res, err := this.Do(http.MethodGet, misScore, nil)
	if err != nil {
		log.Printf("[MisSpider CurrentScore Error] %+v\n", err)
		return nil, err
	}
	return extractCurrentScore(res), nil
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
