package config

import (
	"github.com/go-ini/ini"
	lib_log "github.com/jcsz/gowebchat/library/log"
	lib_tool "github.com/jcsz/gowebchat/library/tool"
	"os"
)

var (
	LogPath string
	LogConf map[string]lib_log.Conf = make(map[string]lib_log.Conf)
)

const (
	L_CFG_SECTION = "Log"
	L_CFG_PATH    = "Path"
	L_CFG_FILE    = "File"
	L_CFG_LEVEL   = "Level"
	L_CFG_SUFFIX  = "Suffix"
)

func Parse_log_config() {
	cfg, err := ini.Load(conf_path + "/log.ini")
	if err != nil {
		panic(err)
	}

	LogPath = cfg.Section(L_CFG_SECTION).Key(L_CFG_PATH).MustString("./var/logs")

	if !lib_tool.File_exists(LogPath) {
		os.MkdirAll(LogPath, 0775)
	}

	for _, sec := range cfg.Sections() {
		section_name := sec.Name()
		if section_name == "DEFAULT" || section_name == "Log" {
			continue
		}
		tmp_log_conf := lib_log.Conf{}
		tmp_log_conf.File = LogPath + "/" + sec.Key(L_CFG_FILE).MustString("")
		tmp_log_conf.Level = sec.Key(L_CFG_LEVEL).MustInt(7)
		tmp_log_conf.Suffix = sec.Key(L_CFG_SUFFIX).MustString("")
		LogConf[section_name] = tmp_log_conf
	}
	return
}
