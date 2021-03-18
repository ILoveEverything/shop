package config

import (
	"fmt"
	"github.com/go-ini/ini"
	"log"
	"os"
	"time"
)

var (
	Cfg          *ini.File
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PageSize     int
	JwtSecret    string
)

func init() {
	var err error
	//加载配置文件
	Cfg, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("加载配置文件app.ini出错", err)
	}
	LoadBase()
	LoadServer()
	LoadApp()
	//启动创建商品图片文件夹
	if err := os.MkdirAll("assets/goodsImages", os.ModePerm); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	//启动创建评论图片文件夹
	if err := os.MkdirAll("assets/evaluateImages", os.ModePerm); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

//获取配置文件RUN_MODE的值
func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

//加载配置文件Server块的值
func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	//获取到服务器端口
	HTTPPort = sec.Key("HTTP_PORT").MustInt(9999)
	//获取读取超时时间
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	//获取写入超时时间
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

//加载配置文件App块的值
func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}
	//JWT令牌
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	//todo
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
