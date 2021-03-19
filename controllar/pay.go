package controllar

import (
	"github.com/gin-gonic/gin"
	"shop/model"
	"shop/utils"
)

type Pay struct {
	PayOrderId int64 `json:"pay_order_id" gorm:"not null;unique" form:"payOrderId"`     //支付订单号
	PayMoney   uint  `json:"pay_money" gorm:"not null" form:"payMoney"`                 //支付金额
	PayType    uint  `json:"pay_type" gorm:"not null" form:"payType"`                   //支付方式:1--银行卡,2--支付宝,3--微信
	PayStatus  bool  `json:"pay_status" gorm:"not null;default:false" form:"payStatus"` //支付状态:true--成功,false--失败
}

//银行
func Bank(c *gin.Context) {
	var PayOrderInfo Pay
	err := c.Bind(&PayOrderInfo)
	if err != nil {
		utils.ReturnJson(c, "接收失败", 400, nil)
		return
	}
	PayOrderInfo.PayType = 1
	PayOrderInfo.PayStatus = true
	utils.ReturnJson(c, "成功", 200, PayOrderInfo)
}

//接收
func PayIn(c *gin.Context) {
	var PayOrderIn Pay
	err := c.Bind(&PayOrderIn)
	if err != nil {
		utils.ReturnJson(c, "数据接收失败", 400, nil)
		return
	}
	//判断支付状态
	if PayOrderIn.PayStatus == false {
		utils.ReturnJson(c, "支付失败", 400, nil)
		return
	}
	//判断支付金额
	var PayOrder model.Order
	PayOrder.OrderId = PayOrderIn.PayOrderId
	msg, err := PayOrder.ShowOrder()
	if err != nil {
		utils.ReturnJson(c, msg, 400, nil)
		return
	}
	if PayOrder.SumMoney != PayOrderIn.PayMoney {
		utils.ReturnJson(c, "支付金额错误", 400, nil)
		return
	}
	//修改订单状态
	PayOrder.OrderStatus = 1
	msg, err = PayOrder.UpdateOrderStatus()
	if err != nil {
		utils.ReturnJson(c, msg, 400, nil)
		return
	}
	utils.ReturnJson(c, "支付成功", 200, PayOrderIn)
}
