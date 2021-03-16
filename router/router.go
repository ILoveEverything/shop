package router

import (
	"github.com/gin-gonic/gin"
	"shop/controllar"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	gin.SetMode("debug")
	//加载模板文件
	//r.LoadHTMLGlob("views/template/*")
	api := r.Group("/v1/api")
	{
		//todo 用户信息
		user := api.Group("/user")
		{
			//user.GET("/login", controllar.LoginPage)
			//user.GET("/register", controllar.RegisterPage)
			//注册
			user.POST("/register", controllar.Register)
			//登录
			user.POST("/login", controllar.Login)
			//修改用户信息
			user.POST("/modifyUserInfo", controllar.ModifyUserInfo)
			//查看用户信息
			user.POST("/showUserInfo", controllar.ShowUserInfo)
			//添加收货地址
			//修改收货地址
			//查询收货地址
			//删除收货地址
			//充值
			//提现
			//消费记录
		}
		//todo 商品
		goods := api.Group("/goods")
		{
			goods.GET("/menu")
			goods.GET("/list")
			goods.GET("/banner")
		}
		//todo 商品详情
		goodsDetails := api.Group("goodsDetails")
		{
			goodsDetails.POST("/")
		}
		//todo 订单
		order := api.Group("/order")
		{
			order.GET("/menu")
			order.GET("/list")
			order.GET("/banner")
		}
	}
	return r
}
