package utils

import (
	"github.com/gin-gonic/gin"
)

//公共的返回
func ReturnJson(c *gin.Context, msg string, code int, data interface{}) {
	//http状态麻
	var status int
	switch status {

	}

	c.JSON(200, gin.H{
		"msg":  msg,
		"code": code,
		"data": data,
	})
}
