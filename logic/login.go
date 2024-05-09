package logic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service/middleware"
	"service/models"
)

// 登录
func Login(c *gin.Context) {
	var user models.User
	var token string
	var code int = http.StatusInternalServerError
	var msg string

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ok := models.LoginCheck(user.UserName, user.PassWord)
	if ok {
		token, code = middleware.GenerateToken(user.UserName)
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
