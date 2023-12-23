package config

import "cqu-backend/src/config/setting"

type serverConfig struct {
	Port string
}

var ServerConfig = serverConfig{
	Port: setting.ServerConfig.GetString("port"),
}
