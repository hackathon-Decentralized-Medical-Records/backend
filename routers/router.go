package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"service/logic"
	"service/middleware"
)

func Init() {
	// 支持跨域
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true, // 允许携带 Cookie
	}))

	// 鉴权 TODO 因为当前就一个接口 所以暂时不做鉴权 在user都授予了角色
	auth := r.Group("/api")
	auth.Use(middleware.JWTAuth())
	{

	}

	router := r.Group("/api")
	{
		router.POST("/login", logic.Login)
	}

	port := viper.GetString("server.port")
	_ = r.Run(port)

	//// 检测用户是否登录
	//v1.GET("/check-login", nil)
	//// 登录
	//v1.GET("/nonce", nil)
	//v1.GET("/login", func(context *gin.Context) {
	//	fmt.Println("login")
	//})
	//// 获取历史位置
	//v1.POST("/location", nil)
}
