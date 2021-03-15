package controllar

import (
	"github.com/gin-gonic/gin"
	"shop/model"
	"shop/utils"
)

func LoginPage(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

//返回注册页面
func RegisterPage(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

//
func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	email := c.PostForm("email")
	if len(phone) == 0 {
		utils.ReturnJson(c, "电话号码不能为空", 200, nil)
		return
	}

	var user model.UserInfo
	user.Password = password
	user.Phone = phone
	user.Email = email
	user.Username = username
	msg, err := user.Create()
	if err != nil {
		utils.ReturnJson(c, msg, 400, nil)
		return
	}
	utils.ReturnJson(c, msg, 200, nil)
	return
}
