package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	dsn := "root:mysql@tcp(127.0.0.1:3306)/shopping?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		TablePrefix:   "xjw_", // 表名前缀，`User`表为`t_users`
		SingularTable: true,   // 使用单数表名，启用该选项后，`User` 表将是`user`
	}})
	if err != nil {
		fmt.Println("数据库连接出错")
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}
