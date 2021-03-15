package model

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"shop/db"
)

type UserInfo struct {
	gorm.Model
	Username string
	Password string
	Phone    string
	Email    string
	Age      string
	Gender   string
	hobby    []string
}

func init() {
	db.DB.AutoMigrate(&UserInfo{})
}

func (u *UserInfo) Create() (msg string, error error) {
	if err := db.DB.
		Model(&u).Where("phone=?", u.Phone).Find(&u).Error; err != nil {
		fmt.Println(err.Error())
		return "查询失败", errors.New("查询失败")
	}
	if u.ID != 0 {
		return "该用户已注册", errors.New("该用户已注册")
	}
	if err := db.DB.Create(&u).Error; err != nil {
		return "查询失败", err
	}
	return "成功注册", nil
}

func (u *UserInfo) Update() {
	panic("implement me")
}

func (u *UserInfo) Retrieve() {
	//db.DB.Where("phone=?",u.Phone).Find(&u)
}

func (u *UserInfo) Delete() {
	panic("implement me")
}
