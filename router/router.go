package router

import (
	"github.com/gin-gonic/gin"
	"shop/controllar"
	"shop/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	gin.SetMode("debug")
	apiNotToken := r.Group("/v1/api")
	{
		//用户信息
		user := apiNotToken.Group("/user")
		{
			//注册
			user.POST("/register", controllar.Register)
			//登录
			user.POST("/login", controllar.Login)
		}
		//todo 商品
		goods := apiNotToken.Group("/goods")
		{
			goods.GET("/menu")
			goods.GET("/list")
			goods.GET("/banner")
		}
		//todo 商品详情
		goodsDetails := apiNotToken.Group("goodsDetails")
		{
			goodsDetails.POST("/")
		}
		//todo 订单
		order := apiNotToken.Group("/order")
		{
			order.GET("/menu")
			order.GET("/list")
			order.GET("/banner")
		}
	}
	apiNeedToken := r.Group("/v1/api")
	{
		apiNeedToken.Use(middleware.JWT())
		user := apiNeedToken.Group("/user")
		{
			//查看用户信息
			user.POST("/showUserInfo", controllar.ShowUserInfo)
			//修改用户信息
			user.POST("/modifyUserInfo", controllar.ModifyUserInfo)
			//todo 注销用户
			user.POST("/deleteUser")
			//添加收货地址
			user.POST("/addAddress", controllar.AddAddress)
			//修改收货地址
			user.POST("/modifyAddress")
			//查询收货地址
			user.POST("/showAddress")
			//删除收货地址
			user.POST("/deleteAddress")
			//查询账户余额
			user.POST("/showMoney")
			//充值
			user.POST("/addMoney")
			//提现
			user.POST("/subMoney")
			//消费记录
			user.POST("logMoney")
		}
	}
	return r
}
