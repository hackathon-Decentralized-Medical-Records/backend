package v1

import (
	"github.com/gin-gonic/gin"
	"service/dao/mysql"
	"service/utils/httputils"
)

type requestParam struct {
	UserName string `json:"username" `
	PassWord string `json:"password" `
	Email    string `json:"email" `
	Role     int    `json:"role"`
	// 职业
	Profession string `json:"profession" `
	// 科室
	DepartmentId uint    `json:"departmentId" `
	WorkTime     string  `json:"workTime" `
	EndTime      string  `json:"endTime" `
	Price        float64 `json:"price" `
	Address      string  `json:"address"`
}

// 注册用户
// @Summary 注册用户
// @Tags 用户
// @description 注册用户
// @param username body string true "用户名"
// @param password body string true "密码"
// @param email body string true "邮箱"
// @param role body int true "角色"
// @param profession body string false "职业"
// @param departmentId body int false "科室"
// @Success 200 {string} string "ok"
// @Router /api/register [post]
func Register(c *gin.Context) {
	var req requestParam
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(httputils.StatusBadRequest, gin.H{
			"code": httputils.StatusBadRequest,
			"msg":  httputils.StatusText(httputils.StatusBadRequest),
		})
		return
	}

	// 校验用户是否存在
	code := mysql.CheckUserByEmail(req.Email)
	if code != httputils.StatusOK {
		c.JSON(code, gin.H{
			"code": code,
			"msg":  httputils.StatusText(code),
		})
		return
	}

	user := mysql.User{
		UserName: req.UserName,
		PassWord: req.PassWord,
		Email:    req.Email,
		Role:     req.Role,
		Address:  req.Address,
	}
	// 注册用户表
	mysql.RegisterUser(&user)

	// 注册医生
	if user.Role == 0 {
		createMedic(&req, user.ID)
	}

	// 注册患者
	if user.Role == 1 {
		createPatient(&req, user.ID)
	}

	c.JSON(httputils.StatusOK, gin.H{
		"code": httputils.StatusOK,
		"msg":  httputils.StatusText(httputils.StatusOK),
	})
}

func createMedic(req *requestParam, id uint) {
	medic := mysql.Medic{
		Name:         req.UserName,
		Profession:   req.Profession,
		WorkTime:     req.WorkTime,
		EndTime:      req.EndTime,
		Price:        req.Price,
		DepartmentID: req.DepartmentId,
		UserID:       id,
	}
	mysql.InsertMedic(&medic)
}

func createPatient(req *requestParam, id uint) {
	patient := mysql.Patient{
		Name:    req.UserName,
		User_Id: id,
	}
	mysql.InsertPatient(&patient)
}
