package controller

import (
	"cqu-backend/src/object"
	"cqu-backend/src/service"
	"cqu-backend/src/tool"
	"github.com/kataras/iris/v12"
)

type ScoreController struct {
	ScoreService   *service.ScoreService
	StudentService *service.StudentService
}

func NewScoreController() *ScoreController {
	return &ScoreController{
		ScoreService:   service.NewScoreService(),
		StudentService: service.NewStudentService(),
	}
}

// PostAll @summary 查询所有成绩
// @Success 200 {object} object.Result
// @tags 考试成绩/score
// @Router /score/all [post]
// @param UnionId formData string true "UnionId" default(666)
func (this *ScoreController) PostAll(Ctx iris.Context) object.Result {
	unionId := Ctx.FormValue("UnionId")
	if tool.CheckIsEmpty(unionId) {
		return object.CheckException(object.RequestError)
	}
	student := this.StudentService.GetStudent(unionId)
	if err := tool.CheckCasAccount(student); err != nil { // 绑定统一认证
		return object.CheckException(err)
	}
	score, err := this.ScoreService.AllScore(student.StuId, student.CasPwd)
	if err != nil {
		return object.CheckException(err)
	}
	return object.DataResult(score)
}
