package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service/model"
	"service/service"
)

// @Summary 获取医生列表
// @Tags 医生
// @Router /api/meidcList [post]
func GetMedicList(c *gin.Context) {
	var medic model.Medic
	if err := c.ShouldBindQuery(&medic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	medicList := service.GetMedicList(&medic)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": medicList,
	})
}
