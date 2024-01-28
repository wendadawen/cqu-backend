package service

import (
	"cqu-backend/src/dao/model"
	"cqu-backend/src/spider"
	"cqu-backend/src/spider/mis"
	"cqu-backend/src/spider/my"
	"cqu-backend/src/tool"
	"log"
)

type RankService struct{}

func NewRankService() *RankService {
	return &RankService{}
}

// Rank 获取排名
func (this *RankService) Rank(StuId, CasPwd string) (*model.Rank, error) {
	Rank := &model.Rank{}
	account := spider.SpiderAccount{
		Account:  StuId,
		Password: CasPwd,
	}
	if tool.CheckGraduate(StuId, CasPwd) {
		Mis, err := mis.NewMisByCas(account)
		if err != nil {
			log.Printf("[RankService Rank Error] Account=%s\n", StuId)
			return nil, err
		}
		Rank, err = Mis.Rank()
	} else {
		My, err := my.NewMyByCas(account)
		if err != nil {
			log.Printf("[RankService Rank Error] Account=%s\n", StuId)
			return nil, err
		}
		Rank, err = My.Rank()
	}
	return Rank, nil
}
