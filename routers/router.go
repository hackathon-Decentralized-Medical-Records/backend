package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	v1 "service/api/v1"
	_ "service/docs"
	"service/logic"
	"service/middleware"
)

func Init() *gin.Engine {
	// 支持跨域
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true, // 允许携带 Cookie
	}))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 鉴权 TODO 因为当前就一个接口 所以暂时不做鉴权 在user都授予了角色
	auth := r.Group("/api")
	auth.Use(middleware.JWTAuth())
	{

	}

	router := r.Group("/api")
	{
		router.POST("/login", logic.Login)
		router.POST("/register", v1.Register)
		router.POST("/registration", v1.CreateRecord)
		router.POST("/updateRecordStatus", v1.UpdateRecordStatus)
		router.GET("/getRecordByMedic", v1.GetRecordByMedic)
		router.GET("/getRecordByPatient", v1.GetRecordByPatient)
		router.POST("/getRecordList", v1.GetList)
	}
	return r

}
