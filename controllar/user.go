package controllar

import (
	"fmt"
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

//注册添加用户
func Register(c *gin.Context) {
	//username := c.PostForm("username")
	//password := c.PostForm("password")
	//phone := c.PostForm("phone")
	//email := c.PostForm("email")
	var user model.UserInfo
	err := c.Bind(&user)
	if err != nil {
		utils.ReturnJson(c, "注册失败", 400, nil)
		return
	}
	if len(user.Phone) == 0 {
		utils.ReturnJson(c, "电话号码不能为空", 400, nil)
		return
	}
	if len(user.Username) == 0 {
		utils.ReturnJson(c, "用户名不能为空", 400, nil)
		return
	}
	if len(user.Password) == 0 {
		utils.ReturnJson(c, "密码不能为空", 400, nil)
		return
	}
	if len(user.Email) == 0 {
		utils.ReturnJson(c, "邮箱不能为空", 400, nil)
		return
	}
	msg, err := user.CreateUser()
	if err != nil {
		utils.ReturnJson(c, msg, 400, nil)
		return
	}
	utils.ReturnJson(c, msg, 200, nil)
	return
}

//用户登录
func Login(c *gin.Context) {
	//password := c.PostForm("password")
	//phone := c.PostForm("phone")
	var user model.UserInfo
	//user.Username = username
	//user.Password = password
	//user.Phone = phone
	if err := c.Bind(&user); err != nil {
		utils.ReturnJson(c, "登录失败", 400, nil)
		return
	}
	password := user.Password
	msg, err := user.RetrieveLogin()
	//fmt.Println(msg, err, "---", user.Password)
	if err != nil {
		utils.ReturnJson(c, msg, 400, nil)
		return
	}
	if user.Password != password {
		utils.ReturnJson(c, "密码错误", 400, nil)
		return
	}
	utils.ReturnJson(c, msg, 200, nil)
}

//查看用户信息
func ShowUserInfo(c *gin.Context) {
	//phone := c.PostForm("phone")
	//var user model.UserInfo
	//user.Phone = phone
	////fmt.Println(user)
	//msg, err := user.RetrieveUserInfo()
	//if err != nil {
	//	utils.ReturnJson(c, msg, 400, nil)
	//	return
	//}
	//utils.ReturnJson(c, msg, 200, user)
	var user model.UserInfo
	if err := c.Bind(&user); err != nil {
		utils.ReturnJson(c, "查看失败", 400, nil)
		return
	}
	msg, err := user.RetrieveUserInfo()
	if err != nil {
		utils.ReturnJson(c, msg, 400, nil)
		return
	}
	utils.ReturnJson(c, msg, 200, user)
}

//修改用户信息
func ModifyUserInfo(c *gin.Context) {
	//oldPhone := c.PostForm("oldPhone")
	//username := c.PostForm("username")
	//password := c.PostForm("password")
	//phone := c.PostForm("phone")
	//email := c.PostForm("email")
	//age := c.PostForm("age")
	//gender := c.PostForm("gender")
	//var user model.UserInfo
	//user.Username = username
	//user.Password = password
	//user.Phone = phone
	//user.Email = email
	//user.Age = age
	//user.Gender = gender
	////打印查看
	//fmt.Println(user, oldPhone)
	//msg, err := user.UpdateUserInfo(oldPhone)
	//if err != nil {
	//	utils.ReturnJson(c, msg, 400, nil)
	//	return
	//}
	//utils.ReturnJson(c, msg, 200, nil)
	var user model.UserInfo
	if err := c.Bind(&user); err != nil {
		utils.ReturnJson(c, "修改失败", 400, nil)
		return
	}
	oldPhone := c.PostForm("oldPhone")
	fmt.Println(user)
	msg, err := user.UpdateUserInfo(oldPhone)
	if err != nil {
		utils.ReturnJson(c, msg, 400, nil)
		return
	}
	utils.ReturnJson(c, msg, 200, nil)
}
