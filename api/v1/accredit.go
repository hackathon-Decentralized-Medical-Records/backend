package v1

import (
	"github.com/gin-gonic/gin"
	"service/model"
	"service/service"
	"strconv"
	"strings"
)

type requestAccredit struct {
	MedicId uint ` json:"medicId" `
	// 挂号id
	RegistrationId uint ` json:"registrationId" `
	// 病例id
	CaseId string ` json:"caseId"`
}

// InertAccredit 新增授权
// @Summary 新增授权
// @Description 新增授权
// @Tags accredit
// @param medicId body uint true "medicId"
// @param registrationId body uint true "registrationId"
// @param caseId body string true "caseId"
// @Success 200 {string} string "{"msg": "success"}"
// @Router /accredit/save [post]
func InertAccredit(c *gin.Context) {
	var entity requestAccredit
	if err := c.ShouldBindJSON(&entity); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	// 判断是否包含逗号
	if strings.Contains(entity.CaseId, ",") {
		// 逗号分割
		caseIds := strings.Split(entity.CaseId, ",")
		for _, caseId := range caseIds {
			// 转成uint
			parseUint, err := strconv.ParseUint(caseId, 10, 32)
			if err != nil {
				c.JSON(400, gin.H{"msg": err.Error()})
				return
			}
			u := uint(parseUint)

			accredut := model.Accredut{
				MedicId:        entity.MedicId,
				RegistrationId: entity.RegistrationId,
				CaseId:         u,
			}
			if err := service.InsertAccredit(&accredut); err != nil {
				c.JSON(400, gin.H{"msg": err.Error()})
				return
			}
		}
	} else {
		// 转成uint
		parseUint, err := strconv.ParseUint(entity.CaseId, 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}
		u := uint(parseUint)

		entity := model.Accredut{
			MedicId:        entity.MedicId,
			RegistrationId: entity.RegistrationId,
			CaseId:         u,
		}

		if err := service.InsertAccredit(&entity); err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}
	}

	c.JSON(200, gin.H{"msg": "success"})
}

// GetAccredit 查询授权
// @Summary 查询授权
// @Description 查询授权
// @Tags accredit
// @Param medicId query uint false "medicId"
// @Param registrationId query uint false "registrationId"
// @Param caseId query uint false "caseId"
// @Success 200 {string} string "{"data": "accredits"}"
// @Router /accredit/getAccredit [POST]
func GetAccredit(c *gin.Context) {
	var entity model.Accredut
	if err := c.ShouldBindQuery(&entity); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	if accredits, err := service.GetAccredit(&entity); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	} else {
		c.JSON(200, gin.H{"data": accredits})
	}
}
