package datasource

import (
	"cqu-backend/src/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
	"xorm.io/core"

	"log"
	"sync"
)

var (
	masterEngine   *xorm.Engine
	masterInitOnce sync.Once
)

func InstanceMaster() *xorm.Engine {
	masterInitOnce.Do(func() {
		// 载入配置
		cfg := config.MysqlConfig
		driverSource := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4",
			cfg.UserName, cfg.Password, cfg.Host, cfg.Port, cfg.DatabaseName)
		//log.Println(driverSource)
		engine, err := xorm.NewEngine(config.DriverName, driverSource)
		if err != nil {
			// 如果初始化出现错误，退出
			log.Fatalln("InstanceMaster error: ", err)
		} else {
			engine.SetConnMaxLifetime(1500 * time.Second)
			engine.SetMaxOpenConns(15)
			engine.SetMaxIdleConns(15)
			engine.ShowSQL(true)
			engine.SetLogLevel(core.LOG_WARNING)
			masterEngine = engine
		}
	})
	return masterEngine
}
