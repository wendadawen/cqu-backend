package router

import (
	"cqu-backend/src/web/controller"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func InitRouter(app *iris.Application) {
	mvc.New(app.Party("/student")).Handle(controller.NewStudentController()) // 学生管理
	mvc.New(app.Party("/card")).Handle(controller.NewCardController())       // 校园卡管理
	mvc.New(app.Party("/exam")).Handle(controller.NewExamController())       // 考试管理
	mvc.New(app.Party("/rank")).Handle(controller.NewRankController())       // 排名管理
	mvc.New(app.Party("/score")).Handle(controller.NewScoreController())     // 成绩管理
	mvc.New(app.Party("/table")).Handle(controller.NewClassController())     // 课表管理
}
