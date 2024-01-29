package controller

import (
	"cqu-backend/src/object"
	"cqu-backend/src/service"
	"cqu-backend/src/tool"
	"github.com/kataras/iris/v12"
)

type ClassController struct {
	ClassService   *service.ClassService
	StudentService *service.StudentService
}

func NewClassController() *ClassController {
	return &ClassController{
		ClassService:   service.NewClassService(),
		StudentService: service.NewStudentService(),
	}
}

// Post
// @summary 查询课表
// @Success 200 {object} object.Result
// @tags 课表/table
// @Router /table [post]
// @param UnionId formData string true "UnionId" default(666)
func (this *ClassController) Post(Ctx iris.Context) object.Result {
	UnionId := Ctx.FormValue("UnionId")
	if tool.CheckIsEmpty(UnionId) {
		return object.CheckException(object.RequestError)
	}
	student := this.StudentService.GetStudent(UnionId)
	if err := tool.CheckCasAccount(student); err != nil {
		return object.CheckException(object.CasNotBound)
	}
	res, err := this.ClassService.ClassSchedule(student.StuId, student.CasPwd)
	if err != nil {
		return object.CheckException(err)
	}
	return object.DataResult(res)
}
