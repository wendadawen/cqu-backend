package controller

import (
	"cqu-backend/src/object"
	"cqu-backend/src/service"
	"cqu-backend/src/tool"
	"github.com/kataras/iris/v12"
	"golang.org/x/xerrors"
)

type CardController struct {
	Ctx         iris.Context
	StuService  *service.StudentService
	CardService *service.CardService
}

func NewCardController() *CardController {
	return &CardController{
		StuService:  service.NewStudentService(),
		CardService: service.NewCardService(),
	}
}

// PostBalance @summary 查询余额和最近消费记录
// @Success 200 {object} object.Result
// @tags 一卡通/card
// @Router /card/balance [post]
// @param UnionId formData string true "UnionId" default(666)
func (this *CardController) PostBalance() object.Result {
	// check unionId
	unionId := this.Ctx.FormValue("UnionId")
	if tool.CheckIsEmpty(unionId) {
		return object.CheckException(object.RequestError) // 请求参数有误
	}
	// 学生必须绑定统一认证
	student := this.StuService.GetStudent(unionId)
	if err := tool.CheckCasAccount(student); err != nil {
		return object.CheckException(err)
	}
	// Unsettle=card说明使用统一认证查不了
	if student.Unsettle != "card" {
		res, err := this.CardService.BalanceByCas(student.StuId, student.CasPwd)
		if err == nil {
			return object.DataResult(res)
		}
		if xerrors.Is(err, object.CardCookieError) { // 一卡通无法获取Cookie
			student.Unsettle = "card" // 只能通过一卡通查，数据库记录一下
			this.StuService.Update(student)
		}
	}
	// 使用一卡通查询消费记录
	if err := tool.CheckCardAccount(student); err != nil {
		return object.CheckException(err)
	}
	res, err := this.CardService.BalanceByAcnt(student.StuId, student.CardPwd)
	if err != nil {
		return object.CheckException(err)
	}
	return object.DataResult(res)
}

// PostFee @summary 查询电费
// @Success 200 {object} object.Result
// @tags 一卡通/card
// @Router /card/fee [post]
// @param UnionId formData string true "UnionId" default(666)
func (this *CardController) PostFee() object.Result {
	unionId := this.Ctx.FormValue("UnionId")
	if tool.CheckIsEmpty(unionId) { // 请求参数检查
		return object.CheckException(object.RequestError)
	}
	student := this.StuService.GetStudent(unionId)
	if err := tool.CheckCasAccount(student); err != nil { // 绑定统一认证
		return object.CheckException(err)
	}
	if err := tool.CheckRoomAccount(student); err != nil { // 绑定宿舍
		return object.CheckException(err)
	}
	if student.Unsettle != "card" {
		res, err := this.CardService.ElectricByCas(student.StuId, student.CasPwd)
		if err == nil {
			return object.DataResult(res)
		}
		if xerrors.Is(err, object.CardCookieError) { // 一卡通无法获取Cookie
			student.Unsettle = "card" // 只能通过一卡通查，数据库记录一下
			this.StuService.Update(student)
		}
	}
	// 使用一卡通查询消费记录
	if err := tool.CheckCardAccount(student); err != nil {
		return object.CheckException(err)
	}
	res, err := this.CardService.ElectricByAcnt(student.StuId, student.CardPwd)
	if err != nil {
		return object.CheckException(err)
	}
	return object.DataResult(res)
}
