package controller

import (
	"cqu-backend/src/object"
	"cqu-backend/src/service"
	"cqu-backend/src/tool"
	"github.com/kataras/iris/v12"
	"strings"
)

type StudentController struct {
	StudentService *service.StudentService
}

func NewStudentController() *StudentController {
	return &StudentController{
		StudentService: service.NewStudentService(),
	}
}

// PostTerm @summary 获取学年学期，开学放假日期，当前周数
// @Success 200 {object} object.Result
// @tags 绑定信息/student
// @Router /student/term [post]
func (this *StudentController) PostTerm(Ctx iris.Context) object.Result {
	return object.DataResult(this.StudentService.GetTerm())
}

// PostAccount @summary 绑定情况
// @Success 200 {object} object.Result
// @tags 绑定信息/student
// @Router /student/account [post]
// @param UnionId formData string true "UnionId" default(666)
func (this *StudentController) PostAccount(Ctx iris.Context) object.Result {
	UnionId := Ctx.FormValue("UnionId")
	if tool.CheckIsEmpty(UnionId) {
		return object.CheckException(object.RequestError)
	}
	account := this.StudentService.GetAccount(UnionId)
	if account == nil {
		return object.CheckException(object.CasNotBound)
	}
	return object.DataResult(account)
}

// PostBindCas @summary 绑定统一身份认证
// @Success 200 {object} object.Result
// @tags 绑定信息/student
// @Router /student/bind/cas [post]
// @param UnionId formData string true "UnionId" default(666)
// @param StuId formData string true "学号" default(202114131169)
// @param CasPwd formData string true "统一认证密码" default(er990715000614)
func (this *StudentController) PostBindCas(Ctx iris.Context) object.Result {
	UnionId := Ctx.FormValue("UnionId")
	StuId := Ctx.FormValue("StuId")
	CasPwd := Ctx.FormValue("CasPwd")
	if tool.CheckIsEmpty(UnionId, StuId, CasPwd) {
		return object.CheckException(object.RequestError)
	}
	student, err := this.StudentService.BindCas(UnionId, StuId, CasPwd)
	if err != nil {
		return object.CheckException(err)
	}
	service.Trans2Account(student)
	return object.DataResult(student)
}

// PostClearCas @summary 清除绑定统一身份认证
// @Success 200 {object} object.Result
// @tags 绑定信息/student
// @Router /student/clear/cas [post]
// @param UnionId formData string true "UnionId" default(666)
func (this *StudentController) PostClearCas(Ctx iris.Context) object.Result {
	UnionId := Ctx.FormValue("UnionId")
	if tool.CheckIsEmpty(UnionId) {
		return object.CheckException(object.RequestError)
	}
	this.StudentService.ClearCas(UnionId)
	return object.EmptyResult()
}

// PostBindRoom @summary 绑定房间
// @Success 200 {object} object.Result
// @tags 绑定信息/student
// @Router /student/bind/room [post]
// @param UnionId formData string true  "UnionId" Default(666)
// @param Room  formData string true   "房间编号" Default(a12s3910)
// @param Campus  formData string true   "校区" Default(A区)
// @param Dom  formData string true   "楼栋" Default(12舍)
// @param RoomNum  formData string true   "房间号" Default(3910)
func (this *StudentController) PostBindRoom(Ctx iris.Context) object.Result {
	UnionId := Ctx.FormValue("UnionId")
	Room := Ctx.FormValue("Room")
	Campus := Ctx.FormValue("Campus")
	Dom := Ctx.FormValue("Dom")
	RoomNum := Ctx.FormValue("RoomNum")
	if tool.CheckIsEmpty(UnionId, Room, Campus, Dom, RoomNum) {
		return object.CheckException(object.RequestError)
	}
	Room = strings.ToUpper(Room)
	res, err := this.StudentService.BindRoom(UnionId, Room, Campus, Dom, RoomNum)
	if err != nil {
		return object.CheckException(err)
	}
	return object.DataResult(res)
}

// PostClearRoom @summary 清除绑定房间
// @Success 200 {object} object.Result
// @tags 绑定信息/student
// @Router /student/clear/room [post]
// @param UnionId formData string true  "UnionId" Default(666)
func (this *StudentController) PostClearRoom(Ctx iris.Context) object.Result {
	UnionId := Ctx.FormValue("UnionId")
	if tool.CheckIsEmpty(UnionId) {
		return object.CheckException(object.RequestError)
	}
	this.StudentService.ClearRoom(UnionId)
	return object.EmptyResult()
}

// PostClearAll @summary 清除所有绑定
// @Success 200 {object} object.Result
// @tags 绑定信息/student
// @Router /student/clear/all [post]
// @param UnionId formData string true  "UnionId" Default(666)
func (this *StudentController) PostClearAll(Ctx iris.Context) object.Result {
	UnionId := Ctx.FormValue("UnionId")
	if tool.CheckIsEmpty(UnionId) {
		return object.CheckException(object.RequestError)
	}
	this.StudentService.ClearAll(UnionId)
	return object.EmptyResult()
}
