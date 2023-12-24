package card

import (
	"cqu-backend/src/bo"
	"cqu-backend/src/object"
	"cqu-backend/src/spider"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
	"net/http"
	"strings"
)

const (
	cardUrl          = "http://card.cqu.edu.cn"
	cardLoginUrl     = "/Login/LoginBySnoQuery"
	cardValidCodeUrl = "/Login/GetValidateCode"
	cardRsaUrl       = "/Common/GetRsaKey"
	cardNcLoginUrl   = "/Login/NcLogin"
	cardBalanceUrl   = "/User/GetCardInfoByAccountNoParm"
	cardNcBalanceUrl = "/NcAccType/GetCurrentAccountList"
	cardElectricUrl  = "/Tsm/TsmCommon"
	cardDetailUrl    = "/Report/GetMyBill"
	cardNcDetailUrl  = "/NcReport/GetPersonTrjn"

	feeReferUrl      = "http://card.cqu.edu.cn:8080/blade-auth/token/thirdToToken/fwdt?referer=app&ticket=%s&from=ehall&cometype="
	feeEleUrl        = "http://card.cqu.edu.cn:8080/charge/feeitem/getThirdData"
	ValidCodeMaxTime = 4
	cardLoginSucceed = `"IsSucceed":true`
	cardLoginFailed  = "用户名或密码错误"
	newFeeItemId     = "182"
	oldFeeItemId     = "181"
)

// Card 一卡通接口
type Card interface {
	Balance() (*bo.BalanceData, error)
	Record() (bo.ConsumptionRecord, error)
}

// CardAccount 一卡通账户，包括账号密码和房间号
type CardAccount struct {
	spider.SpiderAccount        // 学生账号
	Room                 string // 学生房间号
}

// cardImplement 一卡通信息的查询接口，是选择通过统一认证查询还是一卡通查询
type cardImplement interface {
	spider.WithLogin
	spider.WithClientDo
}

// cardTemplate Card和cardImplement接口的实现类
type cardTemplate struct {
	cardImplement cardImplement // 继承一个实现了Login和Do接口的结构体, 如CasCard实现了统一认证登录的Login和Do
	client        *resty.Client
	cardAccount   CardAccount // 一卡通账号

	acnt string
}

// 传入一个实现了Login和Do接口的结构体，如CasCard实现了统一认证登录的Login和Do
func newCardTemplate(cardImplement cardImplement) *cardTemplate {
	return &cardTemplate{cardImplement: cardImplement}
}

// ====================实现cardImplement接口=====================
func (this *cardTemplate) Login() error {
	return this.cardImplement.Login()
}
func (this *cardTemplate) Do(method string, urlPath string, payload map[string]string) (string, error) {
	return this.cardImplement.Do(method, urlPath, payload)
}

// ====================实现Card接口==============================
func (this *cardTemplate) Balance() (*bo.BalanceData, error) {
	err := this.Login()
	if err != nil {
		return nil, err
	}
	res, err := this.Do(http.MethodPost, cardUrl+cardNcBalanceUrl, map[string]string{"json": "true"})
	if err != nil || !strings.Contains(res, `"respInfo\":\"处理成功`) {
		return nil, object.CardBalanceError
	}
	res = strings.ReplaceAll(res, `\`, "")
	res = res[1 : len(res)-1]
	objs := gjson.Get(res, "objs").Array()
	if len(objs) == 0 {
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

	return nil, nil
}
