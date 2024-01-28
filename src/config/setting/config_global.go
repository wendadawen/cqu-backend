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
	ProxyConfig  *viper.Viper
	CquConfig    *viper.Viper
)

func init() {
	conf = viper.New()
	conf.AddConfigPath("C:/Users/23202/Desktop/WeCqu/cqu-backend/")
	//conf.AddConfigPath("./")
	conf.SetConfigName("application") // 正式
	err := conf.ReadInConfig()
	if err != nil {
		log.Fatalf("[Read Config] %+v\n", err.Error())
	}
	conf.WatchConfig()
	conf.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("[Configuration Changed] %+v\n", in.Name)
		setting()
	})
	setting()
}

func setting() {
	ServerConfig = conf.Sub("server")
	MysqlConfig = conf.Sub("mysql")
	ProxyConfig = conf.Sub("proxy")
	CquConfig = conf.Sub("cqu")
}
