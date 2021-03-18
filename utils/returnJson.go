package utils

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

//公共的返回
func ReturnJson(c *gin.Context, msg string, code int, data interface{}) {
	//http状态麻
	var status int
	switch status {

	}

	c.JSON(200, &Response{
		Msg:  msg,
		Code: code,
		Data: data,
	})
}
