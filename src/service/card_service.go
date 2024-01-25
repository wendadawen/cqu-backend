package service

import (
	"cqu-backend/src/bo"
	"cqu-backend/src/spider"
	"cqu-backend/src/spider/card"
	"log"
)

type CardService struct{}

func NewCardService() *CardService {
	return &CardService{}
}

// BalanceByCas 通过统一认证获取最近的消费记录
func (this *CardService) BalanceByCas(StuId string, CasPwd string) (*bo.BalanceData, error) {
	cardAccount := card.CardAccount{
		SpiderAccount: spider.SpiderAccount{
			Account:  StuId,
			Password: CasPwd,
		},
	}
	Card, err := card.NewCardByCas(cardAccount)
	if err != nil {
		log.Printf("[CardService BalanceByCas NewCardByCas Error] Account=%s\n", StuId)
		return nil, err
	}
	balance, err := Card.Balance() // TODO 到这里
	if err != nil {
		log.Printf("[CardService BalanceByCas Balance Error] Account=%s\n", StuId)
		return nil, err
	}
	balance.Record, err = Card.Record()
	if err != nil {
		log.Printf("[CardService BalanceByCas Record Error] Account=%s\n", StuId)
		return nil, err
	}
	return balance, nil
}

// BalanceByAcnt 通过一卡通获取最近的消费记录
func (this *CardService) BalanceByAcnt(StuId string, CardPwd string) (*bo.BalanceData, error) {

	return &bo.BalanceData{}, nil
}

// ElectricByCas 通过统一认证获取宿舍电费
func (this *CardService) ElectricByCas(StuId string, CardPwd string) (*bo.ElectricCharge, error) {

	return &bo.ElectricCharge{}, nil
}

// ElectricByAcnt 通过一卡通获取宿舍电费
func (this *CardService) ElectricByAcnt(StuId string, CardPwd string) (*bo.ElectricCharge, error) {

	return &bo.ElectricCharge{}, nil
}
