package main

import (
	"fmt"
	lib_cfg "github.com/jcsz/gowebchat/library/config"
	lib_log "github.com/jcsz/gowebchat/library/log"
	_ "os"
	"time"
)

func main() {
	for i := 0; i < 10000; i++ {
		lib_log.GetNode("main").Infof("111111111111")
		lib_log.GetNode("default").Infof("2222222222")
		time.Sleep(time.Second)
	}
}

//初始化
func init() {
	/* init log */
	lib_cfg.Parse_log_config()
	if err := lib_log.Init(lib_cfg.LogConf); err != nil {
		panic(err)
	}

	/* init mysql */
	/*
		lib_cfg.Parse_mysql_config()
		if err := lib_mysql.Init(lib_cfg.MysqlConf); err != nil {
			lib_log.GetNode("main").Infof(err)
		}*/

	fmt.Println(lib_cfg.MysqlConf)
	//os.Exit(0)
}
