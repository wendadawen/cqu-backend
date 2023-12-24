package tool

import (
	"cqu-backend/src/config/setting"
	"fmt"
)

type ProxyType string

// 代理常量
const (
	MY    ProxyType = "my"
	LIB   ProxyType = "lib"
	CARD  ProxyType = "card"
	MIS   ProxyType = "mis" // 研究生
	FORCE ProxyType = "force"
	CAS   ProxyType = "cas"
	YOUTH ProxyType = "youth" //志愿者
)

func ShouldProxy(webType ProxyType) bool {
	return setting.ProxyConfig.GetBool(string(webType))
}

var ProxyUrl = fmt.Sprintf("http://%s:%s@%s",
	setting.ProxyConfig.GetString("proxy_user"),
	setting.ProxyConfig.GetString("proxy_pwd"),
	setting.ProxyConfig.GetString("proxy_server"))
