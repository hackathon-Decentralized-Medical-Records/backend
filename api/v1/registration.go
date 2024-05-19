package v1

import (
	"github.com/gin-gonic/gin"
	"service/dao/mysql"
)

// 生成挂号记录
// @Summary 生成挂号记录
// @Description 生成挂号记录
// @Tags registration
// @param medicId body uint true "医生id"
// @param patientId body uint true "患者id"
// @param status body int true "挂号状态"
// @Success 200 {string} string "{"message": "挂号成功"}"
// @Router /api/registration [post]
func CreateRecord(c *gin.Context) {
	var registration mysql.Registration
	err := c.ShouldBindJSON(&registration)
	if err != nil {
		c.JSON(400, gin.H{"message": "参数错误"})
		return
	}
	mysql.InsertRegistration(&registration)
	c.JSON(200, gin.H{"message": "挂号成功"})
}

// 更新挂号记录状态
// @Summary 更新挂号记录状态
// @Description 更新挂号记录状态
// @Tags registration
// @param id body uint true "挂号记录id"
// @param status body int true "挂号状态"
// @Success 200 {string} string "{"message": "更新成功"}"
// @Router /api/updateRecordStatus [post]
func UpdateRecordStatus(c *gin.Context) {
	var registration mysql.Registration
	err := c.ShouldBindJSON(&registration)
	if err != nil {
		c.JSON(400, gin.H{"message": "参数错误"})
		return
	}
	mysql.UpdateStatus(registration.ID, registration.Status)
	c.JSON(200, gin.H{"message": "更新成功"})
}

// 获取挂号记录
// @Summary 获取医生挂号记录
// @Description 获取医生挂号记录
// @Tags registration
// @param medicId body uint true "医生id"
// @Success 200 {string} string "{"data": "挂号记录"}"
// @Router /api/getRecordByMedic [get]
func GetRecordByMedic(c *gin.Context) {
	var registration mysql.Registration
	err := c.ShouldBindJSON(&registration)
	if err != nil {
		c.JSON(400, gin.H{"message": "参数错误"})
		return
	}

	records := mysql.GetRecordByMedic(registration.Medic_Id)
	c.JSON(200, gin.H{"data": records})
}

// 获取患者挂号记录
// @Summary 获取患者挂号记录
// @Description 获取患者挂号记录
// @Tags registration
// @param patientId body uint true "患者id"
// @Success 200 {string} string "{"data": "挂号记录"}"
// @Router /api/getRecordByPatient [get]
func GetRecordByPatient(c *gin.Context) {
	var registration mysql.Registration
	err := c.ShouldBindJSON(&registration)
	if err != nil {
		c.JSON(400, gin.H{"message": "参数错误"})
		return
	}

	records := mysql.GetRecordByPatient(registration.Patient_Id)
	c.JSON(200, gin.H{"data": records})
}

// 获取挂号记录列表  （可能没有必要）。
// @Summary 获取挂号记录列表
// @Description 获取挂号记录列表
// @Tags registration
// @param patientId body uint false "患者id"
// @param medicId body uint false "医生id"
// @param status body int false "挂号状态"
// @Success 200 {string} string "{"data": "挂号记录"}"
// @Router /api/getRecordList [get]
func GetList(c *gin.Context) {
	var registration mysql.Registration
	err := c.ShouldBindJSON(&registration)
	if err != nil {
		c.JSON(400, gin.H{"message": "参数错误"})
		return
	}
	records := mysql.GetList(&registration)
	c.JSON(200, gin.H{"data": records})
}
