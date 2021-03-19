package model

import "gorm.io/gorm"

//商家信息
type BusinessInfo struct {
	gorm.Model
	Username string `json:"username" form:"username" gorm:"not null;unique"` //用户名
	Password string `json:"password" form:"password" gorm:"not null"`        //密码
	Phone    string `json:"phone" form:"phone" gorm:"not null;unique"`       //手机号
	Email    string `json:"email" form:"email" gorm:"not null;unique"`       //邮箱
	Age      string `json:"age" form:"age"`                                  //年龄
	Gender   string `json:"gender" form:"gender"`                            //性别
	Shop     string `json:"shop" form:"shop" gorm:"not null"`                //店铺
}
