package controller

import (
	"cqu-backend/src/object"
	"cqu-backend/src/service"
	"cqu-backend/src/tool"
	"github.com/kataras/iris/v12"
)

type StudentController struct {
	//Ctx            iris.Context
	StudentService *service.StudentService
}

func NewStudentController() *StudentController {
	return &StudentController{
		StudentService: service.NewStudentService(),
	}
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
// @param StuId formData string true "学号" default(666)
// @param CasPwd formData string true "统一认证密码" default(666)
func (this *StudentController) PostBindCas(Ctx iris.Context) object.Result {
	UnionId := Ctx.FormValue("UnionId")
	if tool.CheckIsEmpty(UnionId) {
		return object.CheckException(object.RequestError)
	}
	StuId := Ctx.FormValue("StuId")
	if tool.CheckIsEmpty(StuId) {
		return object.CheckException(object.RequestError)
	}
	CasPwd := Ctx.FormValue("CasPwd")
	if tool.CheckIsEmpty(CasPwd) {
		return object.CheckException(object.RequestError)
	}
	student, err := this.StudentService.BindCas(UnionId, StuId, CasPwd)
	if err != nil {
		return object.CheckException(err)
	}
	service.Trans2Account(student)
	return object.DataResult(student)
}
