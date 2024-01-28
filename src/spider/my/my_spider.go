package my

import (
	"cqu-backend/src/bo"
	"cqu-backend/src/model"
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
	exams := extractExam(res)
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
	return extractRank(res), nil
}

func (this *myTemplate) AllScore() (*bo.MyScoreResultBo, error) {
	err := this.Login()
	if err != nil {
		log.Printf("[MySpider AllScore Error] %+v\n", err)
		return nil, err
	}
	res, err := this.Do(http.MethodGet, myScore, nil)
	if err != nil {
		log.Printf("[MySpider AllScore Error] %+v\n", err)
		return nil, err
	}
	return extractAllScore(res), nil
}

func (this *myTemplate) CurrentScore() (*bo.MyScoreListBo, error) {
	err := this.Login()
	if err != nil {
		log.Printf("[MySpider CurrentScore Error] %+v\n", err)
		return nil, err
	}
	res, err := this.Do(http.MethodGet, myScore, nil)
	if err != nil {
		log.Printf("[MySpider CurrentScore Error] %+v\n", err)
		return nil, err
	}
	return extractCurrentScore(res), nil
}

func (this *myTemplate) ClassSchedule() (*bo.ClassScheduleBo, error) {
	err := this.Login()
	if err != nil {
		log.Printf("[MySpider ClassSchedule Error] %+v\n", err)
		return nil, err
	}
	res, err := this.Do(http.MethodPost, myClass, nil)
	if err != nil {
		log.Printf("[MySpider ClassSchedule Error] %+v\n", err)
		return nil, err
	}
	return extractClassSchedule(res), nil
}

func (this *myTemplate) StudentInfo() (*bo.StudentInfoBo, error) {
	err := this.Login()
	if err != nil {
		log.Printf("[MySpider StudentInfo Error] %+v\n", err)
		return nil, err
	}
	res, err := this.Do(http.MethodGet, myInfo, nil)
	if err != nil {
		log.Printf("[MySpider StudentInfo Error] %+v\n", err)
		return nil, err
	}
	return extractStudentInfo(res), nil
}
