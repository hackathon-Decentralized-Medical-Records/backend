package v1

import (
	"github.com/gin-gonic/gin"
	"service/dao/mysql"
)

// GetDeparmentList 获取科室列表
// @Summary 获取科室列表
// @Description 获取科室列表
// @Tags 科室
// @Router /api/departmentList [get]
func GetDeparmentList(c *gin.Context) {

	department := mysql.GetDepartment()
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": department,
	})
}
