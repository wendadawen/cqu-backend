package service

import (
	"cqu-backend/src/bo"
	"cqu-backend/src/spider"
	"cqu-backend/src/spider/mis"
	"cqu-backend/src/spider/my"
	"cqu-backend/src/tool"
	"log"
)

type ScoreService struct{}

func NewScoreService() *ScoreService {
	return &ScoreService{}
}

func (this *ScoreService) AllScore(StuId string, CasPwd string) (*bo.MyScoreResultBo, error) {
	account := spider.SpiderAccount{
		Account:  StuId,
		Password: CasPwd,
	}
	score := &bo.MyScoreResultBo{}
	if tool.CheckGraduate(StuId, CasPwd) {
		Mis, err := mis.NewMisByCas(account)
		if err != nil {
			log.Printf("[ScoreService AllScore Error] Account=%s\n", StuId)
			return nil, err
		}
		score, err = Mis.AllScore()
		if err != nil {
			log.Printf("[ScoreService AllScore Error] Account=%s\n", StuId)
			return nil, err
		}
	} else {
		My, err := my.NewMyByCas(account)
		if err != nil {
			log.Printf("[ScoreService AllScore Error] Account=%s\n", StuId)
			return nil, err
		}
		score, err = My.AllScore()
		if err != nil {
			log.Printf("[ScoreService AllScore Error] Account=%s\n", StuId)
			return nil, err
		}
	}
	return score, nil
}
