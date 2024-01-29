package controller

import (
	"cqu-backend/src/object"
	"cqu-backend/src/service"
	"cqu-backend/src/tool"
	"github.com/kataras/iris/v12"
)

type RankController struct {
	RankService    *service.RankService
	StudentService *service.StudentService
}

func NewRankController() *RankController {
	return &RankController{
		RankService:    service.NewRankService(),
		StudentService: service.NewStudentService(),
	}
}

// Post
// @summary 查询成绩排名，研究生排名仅供参考
// @Success 200 {object} object.Result
// @tags 成绩排名/rank
// @Router /rank [post]
// @param UnionId formData string true "UnionId" default(666)
func (this *RankController) Post(Ctx iris.Context) object.Result {
	unionId := Ctx.FormValue("UnionId")
	if tool.CheckIsEmpty(unionId) {
		return object.CheckException(object.RequestError)
	}
	student := this.StudentService.GetStudent(unionId)
	if err := tool.CheckCasAccount(student); err != nil {
		return object.CheckException(err)
	}
	rank, err := this.RankService.Rank(student.StuId, student.CasPwd)
	if err != nil {
		return object.CheckException(err)
	}
	return object.DataResult(rank)
}
