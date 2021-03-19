package controllar

import (
	"github.com/gin-gonic/gin"
	"shop/model"
	"shop/utils"
	"time"
)

type OrderPage struct {
	OrderId          int64     `json:"order_id" gorm:"not null;unique;primarykey" form:"orderId"` //订单编号
	CreatedOrderTime time.Time `json:"created_order_time" form:"createOrderTime"`                 //下单时间
	Username         string    `json:"username" form:"username" gorm:"not null"`                  //用户名
	Address          string    `form:"address" gorm:"not null" json:"address"`                    //收货地址
	GoodsImage       string    `gorm:"not null" form:"goodsImage" json:"goods_image_path"`        //商品图片
	GoodsName        string    `gorm:"not null" form:"goodsName" json:"goods_name"`               //商品名称
	GoodsPrice       uint      `gorm:"not null" json:"goods_price" form:"goodsPrice"`             //商品价格
	GoodsModel       string    `gorm:"not null" json:"goods_model" form:"goodsModel"`             //商品型号
	GoodsQuantity    uint      `form:"goodsQuantity" gorm:"not null" json:"goods_quantity"`       //购买数量
	SumMoney         uint      `form:"sumMoney" gorm:"not null" json:"sum_money"`                 //付款总价
	OrderStatus      uint      `json:"order_status" form:"orderStatus" gorm:"not null;default:0"` //订单状态:0--未付款;1--未发货;2--未签收;3--已签收
}

//用户下单
func OrderGoods(c *gin.Context) {
	var UserOrder model.Order
	err := c.Bind(&UserOrder)
	if err != nil {
		utils.ReturnJson(c, "下单失败", 400, nil)
		return
	}
	//查询商品库存
	var Goods model.GoodsDetails
	Goods.GoodsId = UserOrder.GoodsId
	var msg string
	msg, err = Goods.ShowGoodsDetails()
	if err != nil {
		utils.ReturnJson(c, msg, 400, nil)
		return
	}
	if Goods.GoodsQuantity < UserOrder.GoodsQuantity {
		utils.ReturnJson(c, "库存不足", 400, nil)
		return
	}
	//创建一个随机的订单号
	now := time.Now()
	UserOrder.OrderId = now.UnixNano()
	//创建下单时间
	UserOrder.CreatedOrderTime = now
	//获取用户id
	token := c.Request.Header.Get("Authorization")
	claims, _ := utils.ParseToken(token)
	UserOrder.UserId = claims.UID
	//获取用户名
	var UserName model.UserInfo
	UserName.ID = UserOrder.UserId
	msg, err = UserName.RetrieveUserInfo()
	if err != nil {
		utils.ReturnJson(c, msg, 400, nil)
		return
	}
	UserOrder.Username = UserName.Username
	//检查用户收货地址
	if len(UserOrder.Address) == 0 {
		utils.ReturnJson(c, "收货地址不能为空", 400, nil)
		return
	}
	//获取商家id
	Goods.GoodsId = UserOrder.GoodsId
	msg, err = Goods.ShowGoodsDetails()
	if err != nil {
		utils.ReturnJson(c, msg, 400, nil)
		return
	}
	UserOrder.BusinessId = Goods.BusinessId
	//获取商品标题
	UserOrder.GoodsTitle = Goods.GoodsTitle
	//获取商品图片
	UserOrder.GoodsImage = Goods.GoodsImage
	//获取商品名称
	UserOrder.GoodsName = Goods.GoodsName
	//获取商品标签
	UserOrder.GoodsTag = Goods.GoodsTag
	//获取商品父级标签
	UserOrder.GoodsParentTag = Goods.GoodsParentTag
	//获取商品介绍
	UserOrder.GoodsExplain = Goods.GoodsExplain
	//获取商品价格
	UserOrder.GoodsPrice = Goods.GoodsPrice
	//获取商品型号
	UserOrder.GoodsModel = Goods.GoodsModel
	//计算付款金额
	if UserOrder.OrderType == 0 {
		UserOrder.SumMoney = UserOrder.GoodsPrice * UserOrder.GoodsQuantity
	}
	//添加订单到订单表
	msg, err = UserOrder.CreateOrder()
	if err != nil {
		utils.ReturnJson(c, msg, 400, nil)
		return
	}
	//更新商品表商品库存数量
	Goods.GoodsQuantity = Goods.GoodsQuantity - UserOrder.GoodsQuantity
	//fmt.Println(Goods)
	msg, err = Goods.UpdateGoodsQuantity()
	if err != nil {
		utils.ReturnJson(c, msg, 400, nil)
		return
	}
	utils.ReturnJson(c, msg, 200, nil)
}

//查看订单
func ShowOrder(c *gin.Context) {
	var OrderInfo model.Order
	err := c.Bind(&OrderInfo)
	if err != nil {
		utils.ReturnJson(c, "查看失败", 400, nil)
		return
	}
	msg, err := OrderInfo.ShowOrder()
	if err != nil {
		utils.ReturnJson(c, msg, 400, nil)
		return
	}
	var OrderPage OrderPage
	OrderPage.OrderId = OrderInfo.OrderId
	OrderPage.CreatedOrderTime = OrderInfo.CreatedOrderTime
	OrderPage.Username = OrderInfo.Username
	OrderPage.Address = OrderInfo.Address
	OrderPage.GoodsImage = OrderInfo.GoodsImage
	OrderPage.GoodsName = OrderInfo.GoodsName
	OrderPage.GoodsPrice = OrderInfo.GoodsPrice
	OrderPage.GoodsModel = OrderInfo.GoodsModel
	OrderPage.GoodsQuantity = OrderInfo.GoodsQuantity
	OrderPage.SumMoney = OrderInfo.SumMoney
	OrderPage.OrderStatus = OrderInfo.OrderStatus
	utils.ReturnJson(c, msg, 200, OrderPage)
}
