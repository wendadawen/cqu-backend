package main

import (
	"cqu-backend/src/config"
	"cqu-backend/src/web/router"
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	"net/http"
	"time"
)

// 使用 swag init生成
// @title template模板
// @version 1.0
// @description Go 语言编程之旅：一起用 Go 做项目
func main() {
	app := iris.Default()
	app.Configure(iris.WithOptimizations)
	router.InitRouter(app)
	app.Get("/swagger/{any:path}", swagger.CustomWrapHandler(&swagger.Config{
		URL: "http://localhost:" + config.ServerConfig.Port + "/swagger/doc.json",
	}, swaggerFiles.Handler))
	srv := &http.Server{
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		Addr:         ":" + config.ServerConfig.Port,
	}
	_ = app.Run(iris.Server(srv))
}
