package object

import "golang.org/x/xerrors"

type Exception struct {
	Err  error
	Code int
}

func (this Exception) Error() string {
	return this.Err.Error()
}

var (
	//爬虫
	HttpTimeout     = Exception{xerrors.New("Http请求超时"), 100}
	RegNoMatchError = Exception{xerrors.New("正则表达式未能匹配"), 101}
	// 教务相关 Jwc
	JwcAccountError   = Exception{xerrors.New("教务网账号/密码错误"), 200}
	JwcSpiderError    = Exception{xerrors.New("教务网访问过频"), 201}
	JwcNotBound       = Exception{xerrors.New("教务信息未绑定"), 202}
	JwcWebError       = Exception{xerrors.New("教务网暂时无法访问"), 203}
	JwcVPNError       = Exception{xerrors.New("教务网限制内网访问"), 204}
	EmptyRoomError    = Exception{xerrors.New("暂时无法获取空教室"), 205}
	EmptyRoomNilError = Exception{xerrors.New("空教室信息暂未更新"), 206}
	NoRegisterError   = Exception{xerrors.New("您尚未报到注册成功，请到学院咨询并办理相关手续！"), 207}
	// 研究生管理系统相关 Mis
	MisMaintenanceError = Exception{xerrors.New("教务网正在维护⚙，请过会儿再试~"), 208}
	MisWebError         = Exception{xerrors.New("研究生管理系统暂时无法访问"), 209}
	MisVPNError         = Exception{xerrors.New("研究生管理系统限制内网访问"), 2010}
	// 统一认证相关 Cas
	CasNotBound      = Exception{xerrors.New("统一认证未绑定"), 2011}
	CasValidcode     = Exception{xerrors.New("统一认证账号/密码错误次数超过3次，请明日再试"), 2012}
	CasWebError      = Exception{xerrors.New("统一认证无法访问"), 300}
	CasAccountError  = Exception{xerrors.New("账号或密码错误"), 301}
	CasIdNoError     = Exception{xerrors.New("账号应为学号"), 302}
	CasContiuneError = Exception{xerrors.New("当前存在其他用户使用同一帐号登录"), 303}
	CasUpdateError   = Exception{xerrors.New("绑定失败"), 304}
	// 一卡通中心相关 Card
	CardVeriCodeError = Exception{xerrors.New("python识别验证码错误"), 2101}
	CardAccountError  = Exception{xerrors.New("一卡通账号/密码错误"), 210}
	CardRoomError     = Exception{xerrors.New("房间信息有误"), 211}
	CardFeeError      = Exception{xerrors.New("查询电费余额出错"), 212}
	CardBalanceError  = Exception{xerrors.New("获取一卡通余额出错"), 213}
	CardNotBoundError = Exception{xerrors.New("一卡通未绑定"), 214}
	RoomNotBoundError = Exception{xerrors.New("房间未绑定"), 214}
	CardError         = Exception{xerrors.New("一卡通网站正在维护中~"), 215}
	CardBusyError     = Exception{xerrors.New("😣一卡通太忙，请过会儿再试~"), 216}
	CardRsaError      = Exception{xerrors.New("一卡通RSA信息有误"), 217}
	CardEleError      = Exception{xerrors.New("电费查询网站正在维护中~"), 218}
	CardTicketError   = Exception{xerrors.New("费用查询未能成功跳转（ticket 缺失）"), 219}
	CardLocationError = Exception{xerrors.New("费用查询未能成功跳转（Location 缺失）"), 2110}
	CardTokenError    = Exception{xerrors.New("费用查询未能成功跳转（Token 缺失）"), 2111}
	CardCaptchaError  = Exception{xerrors.New("一卡通验证码错误"), 2112}
	CardCookieError   = Exception{xerrors.New("一卡通无法获取Cookie"), 2113}
	// 图书馆相关 Lib
	LibRequestError  = Exception{xerrors.New("图书馆请求错误"), 220}
	LibAccountError  = Exception{xerrors.New("图书馆账号或密码错误"), 221}
	LibNotBoundError = Exception{xerrors.New("图书馆没有绑定"), 222}
	// 其它
	RequestError     = Exception{xerrors.New("请求参数有误"), 900}
	WechatError      = Exception{xerrors.New("微信接口错误"), 901}
	MiniProgramError = Exception{xerrors.New("小程序接口错误"), 901}
	DBError          = Exception{xerrors.New("更新数据库错误"), 901}
	NewAccountError  = Exception{xerrors.New("❤️新生请等开学后再绑定哦❤️"), 902}
	UnknownError     = Exception{xerrors.New("其它错误"), 999}
	errReply         = "服务器错误，请稍后再试" //公众号回复
)
