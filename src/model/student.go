package model

import (
	"cqu-backend/src/config/datasource"
	"log"
)

func init() {
	engine := datasource.InstanceMaster()
	err := engine.CreateTables(&Student{})
	if err != nil {
		log.Fatalf("[CreatTable Student Error] %+v\n", err)
	}
}

type Student struct {
	UnionId     string `xorm:"not null pk comment('unionId') unique VARCHAR(255)"`
	StuId       string `xorm:"comment('学号') VARCHAR(255)"`
	Uid         string `xorm:"comment('学生 uid') VARCHAR(255)"`
	JwcPwd      string `xorm:"comment('教务处密码') VARCHAR(255)" json:"isBoundJwc"`
	CardPwd     string `xorm:"comment('一卡通密码') VARCHAR(255)" json:"isBoundCard"`
	LibPwd      string `xorm:"comment('图书馆密码') VARCHAR(255)" json:"isBoundLib"`
	CasPwd      string `xorm:"comment('统一认证密码') VARCHAR(255)" json:"isBoundCas"`
	Room        string `xorm:"comment('宿舍编码') VARCHAR(255)"`
	Campus      string `xorm:"comment('校区') VARCHAR(255)"`
	Dom         string `xorm:"comment('宿舍') VARCHAR(255)"`
	RoomNum     string `xorm:"comment('电费查询账号') VARCHAR(255)"`
	Balance     string `xorm:"comment('一卡通余额') VARCHAR(255)"`
	Unsettle    string `xorm:"comment('一卡通查询方式') VARCHAR(255)"`
	ElectricFee string `xorm:"comment('电费余额') VARCHAR(255)" json:"-"`
	Name        string `xorm:"comment('姓名') VARCHAR(255)"`
	College     string `xorm:"comment('学院') VARCHAR(255)"`
	Class       string `xorm:"comment('班级') VARCHAR(255)"`
	Major       string `xorm:"comment('专业') VARCHAR(255)"`
}
