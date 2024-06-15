package core

import (
	"service/global"
	"service/initialize"
	"strconv"
)

func Gin() {
	if global.GVA_SETTING.System.UseRedis {
		initialize.Redis()
	}

	routers := initialize.Routers()
	err := routers.Run(":" + strconv.Itoa(global.GVA_SETTING.System.Port))
	if err != nil {
		panic(err)
	}
}
