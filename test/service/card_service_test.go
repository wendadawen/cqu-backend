package service

import (
	"cqu-backend/src/service"
	"cqu-backend/test"
	"cqu-backend/test/tool"
	"fmt"
	"testing"
)

func TestBalanceByCas(t *testing.T) {
	cardService := service.NewCardService()
	balanceData, err := cardService.BalanceByCas(test.StudId1, test.CasPwd1)
	if err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		println(tool.ToString(*balanceData))
	}
}

func TestElectricByCas(t *testing.T) {
	cardService := service.NewCardService()
	ElectricCharge, err := cardService.ElectricByCas(test.StudId1, test.CasPwd1, test.Room)
	if err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		println(tool.ToString(*ElectricCharge))
	}
}

//
//func TestDemo(t *testing.T) {
//	client := resty.New()
//	client.SetHostURL("http://sso.cqu.edu.cn")
//	res, _ := client.R().Get("http://www.baidu.com")
//	println(res.String())
//}
