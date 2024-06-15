package main

import (
	"service/core"
	"service/global"
	"service/initialize"
)

func main() {
	global.GVA_VIPER = core.Viper()
	global.GVA_LOGRUS = core.Logrus()
	global.GVA_DATABASE = initialize.Gorm()
	if global.GVA_DATABASE != nil {
		initialize.RegisterTables()
	}
	core.Gin()
}
