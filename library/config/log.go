package config

import (
	"github.com/go-ini/ini"
)

var (
	LogPath string
)

const (
	L_CFG_SECTION = "Log"
	L_CFG_PATH    = "Path"
)

func Parse_log_config() {
	cfg, err := ini.Load(conf_path + "/log.ini")
	if err != nil {
		panic(err)
	}
	LogPath = cfg.Section(L_CFG_SECTION).Key(L_CFG_PATH).MustString("./var/logs")
	return
}
