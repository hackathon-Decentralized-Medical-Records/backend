package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"service/config"
	"service/global"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true, // 允许携带 Cookie
	})
}

func checkCors(currentOrigin string) *config.CORSWhitelist {
	for _, whitelist := range global.GVA_SETTING.Cors.Whitelist {
		// 遍历配置中的跨域头，寻找匹配项
		if currentOrigin == whitelist.AllowOrigin {
			return &whitelist
		}
	}
	return nil
}
