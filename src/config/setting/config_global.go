package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

var conf *viper.Viper

var (
	ServerConfig *viper.Viper
	MysqlConfig  *viper.Viper
)

func init() {
	conf = viper.New()
	conf.AddConfigPath("./")
	conf.SetConfigName("application")
	err := conf.ReadInConfig()
	if err != nil {
		log.Fatalln("Read Config Error: ", err)
	}
	conf.WatchConfig()
	conf.OnConfigChange(func(in fsnotify.Event) {
		log.Println("Configuration Changed:", in.Name)
		setting()
	})
	setting()
}

func setting() {
	ServerConfig = conf.Sub("server")
	MysqlConfig = conf.Sub("mysql")
}
