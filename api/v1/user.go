package v1

import (
	"github.com/gin-gonic/gin"
	"service/dao/mysql"
	"service/utils/httputils"
)

// 注册用户
func Register(c *gin.Context) {
	var user mysql.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(httputils.StatusBadRequest, gin.H{
			"code": httputils.StatusBadRequest,
			"msg":  httputils.StatusText(httputils.StatusBadRequest),
		})
		return
	}

	// 校验用户是否存在
	code := mysql.CheckUserByEmail(user.Email)
	if code != httputils.StatusOK {
		c.JSON(code, gin.H{
			"code": code,
			"msg":  httputils.StatusText(code),
		})
		c.Abort()
		return
	}

	mysql.RegisterUser(&user)
	c.JSON(httputils.StatusOK, gin.H{
		"code": httputils.StatusOK,
		"msg":  httputils.StatusText(httputils.StatusOK),
	})
}
