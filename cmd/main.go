package main

import (
	"github.com/spf13/viper"
	"service/config"
	"service/dao/mysql"
	"service/log"
	"service/routers"
)

func init() {
	config.Init()
	log.Init(log.LogMode(viper.GetInt("LogConfig.mode")))
	mysql.Init()
}

func main() {
	r := routers.Init()

	r.Run(viper.GetString("server.port"))
}
