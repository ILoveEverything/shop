package controllar

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shop/e"
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
	var user model.UserInfo
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
	//签发Token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		utils.ReturnJson(c, err.Error(), e.ERROR_AUTH_TOKEN, nil)
		return
	}
	utils.ReturnJson(c, msg, 200, token)
}

//查看用户信息
func ShowUserInfo(c *gin.Context) {
	var user model.UserInfo
	if err := c.Bind(&user); err != nil {
		utils.ReturnJson(c, "查看失败", 400, nil)
		return
	}
	fmt.Println("bindUserInfo:", user)
	msg, err := user.RetrieveUserInfo()
	if err != nil {
		utils.ReturnJson(c, msg, 400, nil)
		return
	}
	utils.ReturnJson(c, msg, 200, user)
}

//修改用户信息
func ModifyUserInfo(c *gin.Context) {
	var user model.UserInfo
	if err := c.Bind(&user); err != nil {
		utils.ReturnJson(c, "修改失败", 400, nil)
		return
	}
	//从Token中解析出用户id
	token := c.Request.Header.Get("Authorization")
	claims, _ := utils.ParseToken(token)
	//fmt.Println("claims:", claims)
	user.ID = claims.UID
	//fmt.Println("user.ID:", user.ID)
	msg, err := user.UpdateUserInfo()
	if err != nil {
		utils.ReturnJson(c, msg, 400, nil)
		return
	}
	utils.ReturnJson(c, msg, 200, nil)
}

//添加收货地址
func AddAddress(c *gin.Context) {
	var address model.UserAddress
	err := c.Bind(&address)
	if err != nil {
		utils.ReturnJson(c, "添加收货地址失败", 400, nil)
		return
	}
	token := c.Request.Header.Get("Authorization")
	claims, _ := utils.ParseToken(token)
	address.UserID = claims.UID
	msg, err := address.AddAddress()
	if err != nil {
		utils.ReturnJson(c, msg, 400, nil)
		return
	}
	utils.ReturnJson(c, msg, 200, nil)
}

//修改收货地址
//查询收货地址
//删除收货地址
//查询账户余额
//充值
//提现
//消费记录
