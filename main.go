package main

import (
	"fmt"
	"github.com/jcsz/gowebchat/library/config"
)

func main() {
	fmt.Println(config.LogPath)
}

func init() {
	config.Init()
}
