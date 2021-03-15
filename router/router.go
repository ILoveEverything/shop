package router

import (
	"github.com/gin-gonic/gin"
	"shop/controllar"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	gin.SetMode("debug")
	r.LoadHTMLGlob("views/template/*")
	r.GET("/login", controllar.LoginPage)
	r.GET("/register", controllar.RegisterPage)
	r.POST("/register", controllar.Register)
	return r
}
