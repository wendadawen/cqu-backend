package card

const (
	cardUrl = "http://card.cqu.edu.cn"
	// 统一认证跳转一卡通
	cardAuth2card    = "https://sso.cqu.edu.cn/login?service=http:%2F%2Fcard.cqu.edu.cn:7280%2Fias%2Fprelogin%3Fsysid%3DFWDT%26continueurl%3Dhttp%253A%252F%252Fcard.cqu.edu.cn%252Fcassyno%252Findex"
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
