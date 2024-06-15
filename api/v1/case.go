package v1

import (
	"github.com/gin-gonic/gin"
	"service/service"
	"strconv"
)

// InsertCase 新增病例
// @Summary 新增病例
// @Description 新增病例
// @Tags case
// @param patientId body uint true "patientId"
// @param content body string true "content"
// @Success 200 {string} string "{"msg": "success"}"
// @Router /case/save [post]
func InsertCase(c *gin.Context) {
	var entity service.Case
	if err := c.ShouldBindJSON(&entity); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	if err := service.InsertCase(&entity); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"msg": "success"})
}

// GetCaseByPatientId 查询病例
// @Summary 查询病例
// @Description 查询病例
// @Tags case
// @Param patientId query uint true "patientId"
// @Success 200 {string} string "{"data": "cases"}"
// @Router /case/getCaseByPatientId [get]
func GetCaseByPatientId(c *gin.Context) {
	value := c.Query("patientId")
	// 转成uint
	parseUint, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	u := uint(parseUint)
	if cases, err := service.GetCase(u); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	} else {
		c.JSON(200, gin.H{"data": cases})
	}
}
