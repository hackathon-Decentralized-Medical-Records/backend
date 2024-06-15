package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service/middleware"
	. "service/model"
)

// 登录
// @Summary 登录
// @Description 登录
// @Tags 用户
// @param email body string true "邮箱"
// @param password body string true "密码"
// @Success 200 {string} string "{"message": "登录成功"}"
// @Router /api/login [post]
func Login(c *gin.Context) {
	var user User
	var token string
	var code int = http.StatusInternalServerError
	var msg string

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ok, re := LoginCheck(user.Email, user.PassWord)
	if ok {
		token, code = middleware.GenerateToken(user.Email)
	}

	msg = "登录失败"
	if code == http.StatusOK {
		msg = "登录成功"
	}

	// 返回user信息+token信息
	c.JSON(code, gin.H{
		"message": msg,
		"code":    code,
		"data": gin.H{
			"data":  re,
			"token": token,
		},
	})

	//c.JSON(code, gin.H{
	//	"message": msg,
	//	"code":    code,
	//	"data": {
	//		user: user,
	//		token: token,
	//	},
	//})
}

// 登录校验
func LoginCheck(userName, passWord string) (bool, Response) {
	code, _, response := GetUserByNameAndPwd(userName, passWord)
	if code == http.StatusOK {
		return true, response
	}
	return false, response
}
