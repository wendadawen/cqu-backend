package router

import (
	"cqu-backend/src/web/controller"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func InitRouter(app *iris.Application) {
	mvc.New(app.Party("/card")).Handle(controller.NewCardController()) // 校园卡管理
	mvc.New(app.Party("/exam")).Handle(controller.NewExamController()) // 考试管理
	mvc.New(app.Party("/rank")).Handle(controller.NewRankController()) // 排名管理
}
