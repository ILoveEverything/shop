package model

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"os"
	"shop/db"
)

//用户信息
type UserInfo struct {
	gorm.Model
	Username string `json:"username" form:"username" gorm:"not null;unique"` //用户名
	Password string `json:"password" form:"password" gorm:"not null"`        //密码
	Phone    string `json:"phone" form:"phone" gorm:"not null;unique"`       //手机号
	Email    string `json:"email" form:"email" gorm:"not null;unique"`       //邮箱
	Age      string `json:"age" form:"age"`                                  //年龄
	Gender   string `json:"gender" form:"gender"`                            //性别
}

//用户收货地址
type UserAddress struct {
	gorm.Model
	UserID    uint   `form:"userId" gorm:"not null" json:"user_id"`       //用户id
	AddressNo uint   `form:"addressNo" gorm:"not null" json:"address_no"` //收货地址编号
	Address   string `form:"address" gorm:"not null" json:"address"`      //收货地址
}

func init() {
	err := db.DB.AutoMigrate(&UserInfo{}, &UserAddress{})
	if err != nil {
		os.Exit(-1)
	}
}

//创建用户
func (u *UserInfo) CreateUser() (msg string, err error) {
	var user UserInfo
	if err := db.DB.Model(&UserInfo{}).Where("phone=?", u.Phone).Find(&user).Error; err != nil {
		//fmt.Println(err.Error())
		return "查询失败", errors.New("查询失败")
	}
	if u.ID != 0 {
		return "该手机号已被绑定", errors.New("该手机号已被绑定")
	}
	if err := db.DB.Model(&UserInfo{}).Create(&u).Error; err != nil {
		//fmt.Println("创建用户出错:", err)
		return "查询失败", err
	}
	return "成功注册", nil
}

//修改用户信息
func (u *UserInfo) UpdateUserInfo() (msg string, err error) {
	var user UserInfo
	if len(u.Phone) != 0 {
		if err := db.DB.Model(&UserInfo{}).Where("phone=?", u.Phone).Find(&user).Error; err != nil {
			return "修改失败", err
		}
		if user.ID != 0 {
			return "该手机号已被绑定", errors.New("该手机号已被绑定")
		}
	}
	if len(u.Email) != 0 {
		if err := db.DB.Model(&UserInfo{}).Where("email=?", u.Email).Find(&user).Error; err != nil {
			return "修改失败", err
		}
		if user.ID != 0 {
			return "该邮箱已被绑定", errors.New("该邮箱已被绑定")
		}
	}
	if len(u.Username) != 0 {
		if err := db.DB.Model(&UserInfo{}).Where("username=?", u.Username).Find(&user).Error; err != nil {
			return "修改失败", err
		}
		if user.ID != 0 {
			return "用户名重复", errors.New("用户名重复")
		}
	}
	if err := db.DB.Model(&UserInfo{}).Where("id=?", u.ID).Updates(&u).Error; err != nil {
		return "修改失败", err
	}
	return "修改成功", nil
}

//用户登录查询数据库密码
func (u *UserInfo) RetrieveLogin() (msg string, err error) {
	if err := db.DB.Model(&UserInfo{}).Where("phone=?", u.Phone).Find(&u).Error; err != nil {
		fmt.Println(err.Error())
		return "查询失败", errors.New("查询失败")
	}
	if u.ID == 0 {
		return "该用户未注册", errors.New("该用户未注册")
	}
	return "成功", nil
}

//查看用户信息
func (u *UserInfo) RetrieveUserInfo() (msg string, err error) {
	if err := db.DB.Model(&UserInfo{}).Where("id=?", u.ID).Find(&u).Error; err != nil {
		//fmt.Println(err.Error())
		return "查询失败", errors.New("查询失败")
	}
	if u.ID == 0 {
		return "查无此人", errors.New("查无此人")
	}
	return "成功", nil
}

//注销账户
func (u *UserInfo) DeleteUser() {
	panic("implement me")
}

//添加收货地址
func (u *UserAddress) AddAddress() (msg string, err error) {
	if err := db.DB.Model(&UserAddress{}).Create(&u).Error; err != nil {
		return "添加收货地址失败", errors.New("添加收货地址失败")
	}
	return "添加收货地址成功", nil
}

//查看收货地址
func (u *UserAddress) ShowAddress() (msg string, err error, data []*UserAddress) {
	var AddressArray []*UserAddress
	err = db.DB.Model(&UserAddress{}).Where("user_id=?", u.UserID).Find(&AddressArray).Error
	if err != nil {
		return "查找失败", err, nil
	}
	if len(AddressArray) == 0 {
		return "未查到", errors.New("未查到"), nil
	}
	return "查找成功", nil, AddressArray
}
