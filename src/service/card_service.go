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
	Card, err := card.NewCardByCas(cardAccount) // 通过一卡通的方式登录校园卡官网查询
	if err != nil {
		log.Printf("[CardService BalanceByCas Error] Account=%s\n", StuId)
		return nil, err
	}
	balance, err := Card.Balance()
	if err != nil {
		log.Printf("[CardService BalanceByCas Error] Account=%s\n", StuId)
		return nil, err
	}
	balance.Record, err = Card.Record()
	if err != nil {
		log.Printf("[CardService BalanceByCas Error] Account=%s\n", StuId)
		return nil, err
	}
	return balance, nil
}

// ElectricByCas 通过统一认证获取宿舍电费
func (this *CardService) ElectricByCas(StuId string, CasPwd string, Room string) (*bo.ElectricCharge, error) {
	cardAccount := card.CardAccount{
		SpiderAccount: spider.SpiderAccount{
			Account:  StuId,
			Password: CasPwd,
		},
		Room: Room,
	}
	Card, err := card.NewCardByCas(cardAccount)
	if err != nil {
		log.Printf("[CardService ElectricByCas Error] Account=%s\n", StuId)
		return nil, err
	}
	fee, err := Card.RoomElectricCharge()
	if err == nil {
		log.Printf("[CardService ElectricByCas Error] Account=%s\n", StuId)
		return nil, err
	}
	return fee, nil
}

// BalanceByAcnt 通过一卡通获取最近的消费记录
func (this *CardService) BalanceByAcnt(StuId string, CardPwd string) (*bo.BalanceData, error) {

	return &bo.BalanceData{}, nil
}

// ElectricByAcnt 通过一卡通获取宿舍电费
func (this *CardService) ElectricByAcnt(StuId string, CardPwd string) (*bo.ElectricCharge, error) {

	return &bo.ElectricCharge{}, nil
}
