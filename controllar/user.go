package controllar

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shop/e"
	"shop/model"
	"shop/utils"
)

//返回前端用户收货地址
type UserAddressPage struct {
	AddressNo uint   `form:"addressNo" gorm:"not null" json:"address_no"` //收货地址编号
	Address   string `form:"address" gorm:"not null" json:"address"`      //收货地址
}

func LoginPage(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

//返回注册页面
func RegisterPage(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

//注册添加用户
func Register(c *gin.Context) {
	var User model.UserInfo
	err := c.Bind(&User)
	if err != nil {
		utils.ReturnJson(c, "注册失败", 400, nil)
		return
	}
	if len(User.Phone) == 0 {
		utils.ReturnJson(c, "电话号码不能为空", 400, nil)
		return
	}
	if len(User.Username) == 0 {
		utils.ReturnJson(c, "用户名不能为空", 400, nil)
		return
	}
	if len(User.Password) == 0 {
		utils.ReturnJson(c, "密码不能为空", 400, nil)
		return
	}
	if len(User.Email) == 0 {
		utils.ReturnJson(c, "邮箱不能为空", 400, nil)
		return
	}
	msg, err := User.CreateUser()
	if err != nil {
		utils.ReturnJson(c, msg, 400, nil)
		return
	}
	utils.ReturnJson(c, msg, 200, nil)
	return
}

//用户登录
func Login(c *gin.Context) {
	var User model.UserInfo
	if err := c.Bind(&User); err != nil {
		utils.ReturnJson(c, "登录失败", 400, nil)
		return
	}
	password := User.Password
	msg, err := User.RetrieveLogin()
	//fmt.Println(msg, err, "---", user.Password)
	if err != nil {
		utils.ReturnJson(c, msg, 400, nil)
		return
	}
	if User.Password != password {
		utils.ReturnJson(c, "密码错误", 400, nil)
		return
	}
	//签发Token
	token, err := utils.GenerateToken(User.ID)
	if err != nil {
		utils.ReturnJson(c, err.Error(), e.ERROR_AUTH_TOKEN, nil)
		return
	}
	utils.ReturnJson(c, msg, 200, token)
}

//查看用户信息
func ShowUserInfo(c *gin.Context) {
	var User model.UserInfo
	if err := c.Bind(&User); err != nil {
		utils.ReturnJson(c, "查看失败", 400, nil)
		return
	}
	fmt.Println("bindUserInfo:", User)
	msg, err := User.RetrieveUserInfo()
	if err != nil {
		utils.ReturnJson(c, msg, 400, nil)
		return
	}
	utils.ReturnJson(c, msg, 200, User)
}

//修改用户信息
func ModifyUserInfo(c *gin.Context) {
	var User model.UserInfo
	if err := c.Bind(&User); err != nil {
		utils.ReturnJson(c, "修改失败", 400, nil)
		return
	}
	//从Token中解析出用户id
	token := c.Request.Header.Get("Authorization")
	claims, _ := utils.ParseToken(token)
	//fmt.Println("claims:", claims)
	User.ID = claims.UID
	//fmt.Println("user.ID:", user.ID)
	msg, err := User.UpdateUserInfo()
	if err != nil {
		utils.ReturnJson(c, msg, 400, nil)
		return
	}
	utils.ReturnJson(c, msg, 200, nil)
}

//添加收货地址
func AddAddress(c *gin.Context) {
	var Address model.UserAddress
	err := c.Bind(&Address)
	if err != nil {
		utils.ReturnJson(c, "添加收货地址失败", 400, nil)
		return
	}
	token := c.Request.Header.Get("Authorization")
	claims, _ := utils.ParseToken(token)
	Address.UserID = claims.UID
	msg, err := Address.AddAddress()
	if err != nil {
		utils.ReturnJson(c, msg, 400, nil)
		return
	}
	utils.ReturnJson(c, msg, 200, nil)
}

//修改收货地址
//查询收货地址
func GetAddress(c *gin.Context) {
	var Address model.UserAddress
	//解析用户id
	token := c.Request.Header.Get("Authorization")
	claims, _ := utils.ParseToken(token)
	Address.UserID = claims.UID
	msg, err, Data := Address.ShowAddress()
	if err != nil {
		utils.ReturnJson(c, msg, 400, nil)
		return
	}
	var PageUserAddress UserAddressPage
	var PageUserAddressArray []UserAddressPage
	for _, v := range Data {
		PageUserAddress.AddressNo = v.AddressNo
		PageUserAddress.Address = v.Address
		PageUserAddressArray = append(PageUserAddressArray, PageUserAddress)
	}
	utils.ReturnJson(c, msg, 200, PageUserAddressArray)
}

//删除收货地址
//查询账户余额
//充值
//提现
//消费记录
