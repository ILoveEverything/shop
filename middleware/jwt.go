package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/e"
	"shop/utils"
	"time"
)

//Token中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = e.SUCCESS
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			code = e.INVALID_PARAMS
			c.JSON(http.StatusUnauthorized, gin.H{"code": code, "msg": e.GetMsg(code), "data": data})
			c.Abort()
			return
		}
		claims, err := utils.ParseToken(token)
		//fmt.Println(claims.ExpiresAt)
		if err != nil {
			code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			c.JSON(http.StatusUnauthorized, gin.H{"code": code, "msg": e.GetMsg(code), "data": data})
			c.Abort()
			return
		}
		if time.Now().Unix() > claims.ExpiresAt {
			code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			c.JSON(http.StatusUnauthorized, gin.H{"code": code, "msg": e.GetMsg(code), "data": data})
			c.Abort()
			return
		}
		if claims.Issuer != "谢金武" {
			code = e.ERROR_AUTH
			c.JSON(http.StatusUnauthorized, gin.H{"code": code, "msg": e.GetMsg(code), "data": data})
			c.Abort()
			return
		}
		c.Next()
	}
}
