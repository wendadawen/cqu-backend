# 重庆大学学生信息后端接口
## 程序架构
- src：源代码
  - bo：代码中临时需要的一些结构体
  - config：配置
    - datasource：数据库来源
    - setting：使用单独的setting包是为了初始化
  - dao：数据层
  - model：数据库对应的数据模型
  - object：一些其他方面的操作
  - service：服务层
  - spider：爬虫程序
    - card：校园卡爬虫
    - cas：统一认证爬虫
    - my：智慧教务爬虫
    - mis：研究生系统爬虫
    - api.go：定义爬虫的通用登录接口和请求接口
  - tool：工具库
  - web：网络层
    - controller：控制器
    - router：路由
- test：代码测试
- application.yaml：配置信息
- main.go：启动程序

## 接口文档
- 使用swag init生成接口文档
## 运行
- 使用go run main.go自动生成数据库表
- 需要到application.yaml配置数据库信息
## Linux运行
- 执行`SET CGO_ENABLED=0 && SET GOOS=linux&& SET GOARCH=amd64&& go build -o output/wechat_linux_linux`
- 登录服务器把文件放在特等位置
- 执行`chmod 777 wechat_linux_linux` `sh restart.sh`