package main

import (
	"fmt"
	"github.com/jcsz/gowebchat/library/config"
)

func main() {
	fmt.Println(config.LogPath)
	fmt.Println(config.MysqlConf)
}

func init() {
	config.Init()
}
