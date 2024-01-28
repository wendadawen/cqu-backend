package config

import "cqu-backend/src/config/setting"

type cquConfig struct {
	TermCurrentMy string
}

var CquConfig = cquConfig{
	TermCurrentMy: setting.CquConfig.GetString("term_current_my"),
}
