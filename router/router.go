package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/controllar"
	"shop/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	//设置线程数量
	//runtime.GOMAXPROCS(1)
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
		//支付
		pay := apiNotToken.Group("/pay")
		{
			pay.POST("/bank", controllar.Bank)
			pay.POST("/in", controllar.PayIn)
		}
	}

	//需要携带Token
	apiUserNeedToken := r.Group("/v1/api")
	{
		apiUserNeedToken.Use(middleware.JWT())
		{
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
				user.GET("/showAddress", controllar.GetAddress)
				//删除收货地址
				user.POST("/deleteAddress")
			}
			//订单
			order := apiUserNeedToken.Group("/order")
			{
				//创建订单
				order.POST("/start", controllar.OrderGoods)
				//查看订单
				order.POST("/showOrder", controllar.ShowOrder)
				//秒杀
				order.POST("/SecondKill", controllar.SecondKill)
			}
			//商家信息
			business := apiUserNeedToken.Group("/business")
			{
				//修改商家信息
				business.POST("/modifyBusinessInfo")
				//todo 注销商户
				business.POST("/deleteBusiness")
			}
			//商品
			goods := apiUserNeedToken.Group("/goods")
			{
				//添加商品
				goods.POST("/createGoods", controllar.CreateGoods)
				//上传商品图片
				goods.POST("/addGoodsImage", controllar.AddGoodsImages)
				//更新商品
				goods.POST("/updateGoods", controllar.UpdateGoods)
				//todo 删除商品
				//创建秒杀活动
				goods.POST("/createSecondKillInfo", controllar.CreateSecondKill)
			}
		}
	}
	return r
}
