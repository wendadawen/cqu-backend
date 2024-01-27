package card

import (
	"cqu-backend/src/bo"
	"cqu-backend/src/object"
	"fmt"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
	"log"
	"net/http"
	"strings"
	"time"
)

func (this *cardTemplate) Balance() (*bo.BalanceData, error) {
	err := this.Login()
	if err != nil {
		log.Printf("[CardSpider Balance Error] %+v\n", err)
		return nil, err
	}
	res, err := this.Do(http.MethodPost, cardUrl+cardNcBalanceUrl, map[string]string{"json": "true"})
	if err != nil || !strings.Contains(res, `"respInfo\":\"处理成功`) {
		log.Printf("[CardSpider Balance Error] %+v\n", err)
		return nil, object.CardBalanceError
	}
	res = strings.ReplaceAll(res, `\`, "")
	res = res[1 : len(res)-1]
	objs := gjson.Get(res, "objs").Array()
	if len(objs) == 0 {
		log.Printf("[CardSpider Balance Error] %+v\n", object.CardBalanceError)
		return nil, object.CardBalanceError
	}
	balance := objs[0].Get("acctAmt").Int()
	unsettle := objs[0].Get("tmpBal").Int()
	return &bo.BalanceData{
		Balance:  fmt.Sprintf("%.2f", cast.ToFloat32(balance)/100.0),
		Unsettle: fmt.Sprintf("%.2f", cast.ToFloat32(unsettle)/100.0),
	}, nil
}
func (this *cardTemplate) Record() (bo.ConsumptionRecord, error) {
	err := this.Login()
	if err != nil {
		log.Printf("[CardSpider Record Error] %+v\n", err)
		return nil, err
	}
	now := time.Now()
	start := now.AddDate(0, -3, 0)
	sdate := start.Format("2006-01-02")
	edate := now.Format("2006-01-02")
	res, err := this.Do(http.MethodPost, cardUrl+cardNcDetailUrl, map[string]string{
		"sdate": sdate,
		"edate": edate,
		//"account":  this.acnt, // 统一认证不需要这个参数，一卡通登录需要
		"trancode": "01,02,03,13,14,15,17,18,21,23,25,39,40,41,42,43,44,45,49,A0,A1,A4,5A,5B",
		"page":     "1",
		"rows":     "1000",
	})
	if err != nil {
		log.Printf("[CardSpider Record Error] %+v\n", err)
		return nil, err
	}
	record := ParseRecord(res)
	return record, nil
}

// 这个网站： http://card.cqu.edu.cn/Phone/Index
func (this *cardTemplate) RoomElectricCharge() (*bo.ElectricCharge, error) {
	err := this.Login()
	if err != nil {
		log.Printf("[CardSpider RoomElectricCharge Error] %+v\n", err)
		return nil, err
	}
	room := this.account.Room
	var campus campusFee
	if strings.Contains(room, "S") { // 需要在调用前进行大写转换，以及合法性判定
		campus = oldCampusFee
	} else if strings.Contains(room, "XHC") { // A区新华村学生宿舍
		campus = oldCampusFee
	} else {
		campus = newCampusFee
	}
	charge, err := this.eleChargeAt(campus)
	if err != nil {
		log.Printf("[CardSpider RoomElectricCharge Error] %+v\n", err)
		return nil, err
	}
	return charge, nil
}
