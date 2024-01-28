package config

import "cqu-backend/src/config/setting"

type cquConfig struct {
	TermCurrentMy  string
	TermCurrentMis string
}

var CquConfig = cquConfig{
	TermCurrentMy:  setting.CquConfig.GetString("term_current_my"),
	TermCurrentMis: setting.CquConfig.GetString("term_current_mis"),
}
