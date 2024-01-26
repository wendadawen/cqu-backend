package card

import (
	"cqu-backend/src/bo"
	"cqu-backend/src/spider"
)

/*
	Card抽象外层需要的接口函数
	CardImplement抽象校园卡的Login和Do方法
	CardTemplate实现Card和CardImplement
	若用不同方法登录校园卡，只需要实现CardImplement的两个方法Login和Do方法
	例如CasCard实现了Login和Do方法，使用CasCard创建一个CardTemplate，返回最外层一个Card用来调用抽象接口方法
*/

// Card 一卡通接口
type Card interface {
	Balance() (*bo.BalanceData, error)
	Record() (bo.ConsumptionRecord, error)
	RoomElectricCharge() (*bo.ElectricCharge, error)
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
	account       CardAccount
	acnt          string
}

// 传入一个实现了Login和Do接口的结构体，如CasCard实现了统一认证登录的Login和Do
func newCardTemplate(cardImplement cardImplement) *cardTemplate {
	return &cardTemplate{cardImplement: cardImplement}
}

func (this *cardTemplate) Login() error {
	return this.cardImplement.Login()
}
func (this *cardTemplate) Do(method string, urlPath string, payload map[string]string) (string, error) {
	return this.cardImplement.Do(method, urlPath, payload)
}
