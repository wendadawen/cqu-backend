package controller

import (
	"cqu-backend/src/object"
	"cqu-backend/src/service"
	"cqu-backend/src/tool"
	"github.com/kataras/iris/v12"
)

type ExamController struct {
	ExamService    *service.ExamService
	StudentService *service.StudentService
}

func NewExamController() *ExamController {
	return &ExamController{
		ExamService:    service.NewExamService(),
		StudentService: service.NewStudentService(),
	}
}

// Post
// @summary 查询考试安排
// @Success 200 {object} object.Result
// @tags 考试安排/exam
// @Router /exam [post]
// @param UnionId formData string true "UnionId" default(666)
func (this *ExamController) Post(Ctx iris.Context) object.Result {
	unionId := Ctx.FormValue("UnionId")
	if tool.CheckIsEmpty(unionId) {
		return object.CheckException(object.RequestError)
	}
	student := this.StudentService.GetStudent(unionId)
	if err := tool.CheckCasAccount(student); err != nil { // 绑定统一认证
		return object.CheckException(err)
	}
	exam, err := this.ExamService.ExamSchedule(student.StuId, student.CasPwd)
	if err != nil {
		return object.CheckException(err)
	}
	return object.DataResult(exam)
}
