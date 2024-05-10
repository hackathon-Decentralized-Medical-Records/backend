package main

import (
	"service/config"
	"service/dao/mysql"
	"service/log"
	"service/routers"
)

func init() {
	config.Init()
	config.LoadConfig()
	log.Init(log.LogMode(config.LogConfigMode))
	mysql.Init()
}

func main() {
	r := routers.Init()
	r.Run(config.ServerPort)
}
