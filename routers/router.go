package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	v1 "service/api/v1"
	"service/dao/ipfs"
	_ "service/docs"
	"service/logic"
	"service/middleware"
)

func Init() *gin.Engine {
	// 支持跨域
	r := gin.Default()
	//r.Static("/uploads", "./uploads")
	r.GET("/ipfs", ipfs.CheckFileExistAndGet)
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
		// 登录
		router.POST("/login", logic.Login)
		// 注册用户
		router.POST("/register", v1.Register)
		// 生成挂号记录
		router.POST("/registration", v1.CreateRecord)
		// 更新挂号记录状态
		router.POST("/updateRecordStatus", v1.UpdateRecordStatus)
		router.GET("/getRecordByMedic", v1.GetRecordByMedic)
		// 获取患者挂号记录
		router.GET("/getRecordByPatient", v1.GetRecordByPatient)
		//  获取挂号记录列表
		router.POST("/getRecordList", v1.GetList)
		// 获取医生列表
		router.POST("/meidcList", v1.GetMedicList)
		// 获取科室列表
		router.GET("/departmentList", v1.GetDeparmentList)
		router.POST("/provider/upload", ipfs.UploadFileHandler)
		// 新增病例
		router.POST("/case/save", v1.InsertCase)
		// 查询病例
		router.GET("/case/getCaseByPatientId", v1.GetCaseByPatientId)
		// 新增授权
		router.POST("/accredit/save", v1.InertAccredit)
		// 查询授权
		router.POST("/accredit/accreditList", v1.GetAccredit)

	}
	return r
}
