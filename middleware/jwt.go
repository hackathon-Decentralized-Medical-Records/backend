package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"service/global"
	"time"
)

var JWTKEY = []byte(global.GVA_SETTING.JWT.JwtKey)

type Claims struct {
	UserName string `json:"userName"`
	jwt.StandardClaims
}

// 生成token
func GenerateToken(loginId string) (string, int) {
	expireTime := time.Now().Add(2 * time.Hour) //2小时过期

	// 生成token
	claims := Claims{
		loginId,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ginblog",
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(JWTKEY)
	if err != nil {
		return "", http.StatusInternalServerError
	}
	return token, http.StatusOK
}

// 验证token
func ParseToken(token string) (*Claims, int) {
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTKEY, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, http.StatusBadRequest
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, http.StatusUnauthorized
			} else {
				return nil, http.StatusBadRequest
			}
		}
	}
	if claims, ok := t.Claims.(*Claims); ok && t.Valid {
		return claims, http.StatusOK
	}
	return nil, http.StatusUnauthorized
}

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}
		claims, code := ParseToken(token)
		if code != http.StatusOK {
			c.JSON(code, gin.H{
				"msg": "token 验证失败",
			})
			c.Abort()
			return
		}

		// 校验token是否过期
		if time.Now().Unix() > claims.ExpiresAt {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "token 已过期",
			})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
