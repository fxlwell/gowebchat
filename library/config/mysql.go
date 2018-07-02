package config

import (
	"fmt"
	"github.com/go-ini/ini"
	lib_mysql "github.com/jcsz/gowebchat/library/mysql"
	"strings"
)

var (
	MysqlConf map[string]lib_mysql.MysqlConf = make(map[string]lib_mysql.MysqlConf)
)

const (
	/* common selection */
	M_CFG_COMMON_USERNAME = "Username"
	M_CFG_COMMON_PASSWORD = "Password"
	M_CFG_COMMON_DATABASE = "Database"
	M_CFG_COMMON_ADDR     = "Addr"
)

func Parse_mysql_config() {
	cfg, err := ini.Load(conf_path + "/mysql.ini")
	if err != nil {
		panic(err)
	}
	for _, sec := range cfg.Sections() {
		section_name := sec.Name()
		if section_name == "DEFAULT" {
			continue
		}
		node_cluser := strings.Split(section_name, "-")
		if len(node_cluser) != 2 {
			panic(fmt.Sprintf("mysql.ini -> section name '%s' is invalid. myqsl node and cluser must be separated by '-' like 'default-master'", section_name))
		}
		tmpMap := lib_mysql.MysqlConf{}
		tmpMap.Username = sec.Key(M_CFG_COMMON_USERNAME).MustString("")
		tmpMap.Password = sec.Key(M_CFG_COMMON_PASSWORD).MustString("")
		tmpMap.Database = sec.Key(M_CFG_COMMON_DATABASE).MustString("")
		tmpMap.Addr = sec.Key(M_CFG_COMMON_ADDR).MustString("")
		MysqlConf[section_name] = tmpMap
	}
	return
}
