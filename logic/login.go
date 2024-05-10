package logic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service/dao/mysql"
	"service/middleware"
)

// 登录
func Login(c *gin.Context) {
	var user mysql.User
	var token string
	var code int = http.StatusInternalServerError
	var msg string

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ok := LoginCheck(user.Email, user.PassWord)
	if ok {
		token, code = middleware.GenerateToken(user.Email)
	}

	msg = "登录失败"
	if code == http.StatusOK {
		msg = "登录成功"
	}

	c.JSON(code, gin.H{
		"message": msg,
		"code":    code,
		"token":   token,
	})
}

// 登录校验
func LoginCheck(userName, passWord string) bool {
	code, _ := mysql.GetUserByNameAndPwd(userName, passWord)
	if code == http.StatusOK {
		return true
	}
	return false
}
