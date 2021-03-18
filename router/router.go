package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/controllar"
	"shop/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	//限制文件上传大小
	r.MaxMultipartMemory = 8 << 20
	//加载静态文件
	r.StaticFS("assets", http.Dir("./assets"))
	//设置model模式
	gin.SetMode("debug")

	//不需要携带Token
	apiNotToken := r.Group("/v1/api")
	{
		//用户信息
		user := apiNotToken.Group("/user")
		{
			//注册
			user.POST("/register", controllar.Register)
			//登录
			user.POST("/login", controllar.Login)
			//查看用户信息
			user.POST("/showUserInfo", controllar.ShowUserInfo)
		}
		//商家信息
		business := apiNotToken.Group("/business")
		{
			//商家注册
			business.POST("/register")
			//商家登录
			business.POST("login")
		}
		//商品
		goods := apiNotToken.Group("/goods")
		{
			//商品分类
			goods.GET("/menu", controllar.GoodsMenu)
			//浏览商品
			goods.GET("/list", controllar.GoodsList)
			//商品轮播图
			goods.GET("/banner")
			//商品详情
			goods.POST("/Details", controllar.GoodsDetails)
		}
	}

	//需要携带Token
	apiUserNeedToken := r.Group("/v1/api")
	{
		apiUserNeedToken.Use(middleware.JWT())
		//用户信息操作
		user := apiUserNeedToken.Group("/user")
		{
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
		}
		//todo 订单
		order := apiNotToken.Group("/order")
		{
			order.GET("/menu")
			order.GET("/list")
			order.GET("/banner")
		}
	}
	//商家操作需要携带Token
	apiBusinessNeedToken := r.Group("/v1/api")
	{
		apiBusinessNeedToken.Use(middleware.JWT())
		//商家信息
		business := apiBusinessNeedToken.Group("/business")
		{
			//修改商家信息
			business.POST("/modifyBusinessInfo")
			//todo 注销商户
			business.POST("/deleteBusiness")
		}
		//商品
		goods := apiBusinessNeedToken.Group("/goods")
		{
			//上传商品
			goods.POST("/updateGoods", controllar.UpdateGoods)
			//上传商品图片
			goods.POST("/updateGoodsImage", controllar.UpdateGoodsImages)
			//修改商品
			goods.POST("/modifyGoods")
			//删除商品
		}
	}

	return r
}
