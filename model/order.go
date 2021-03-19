package model

import (
	"fmt"
	"os"
	"shop/db"
	"time"
)

//用户订单表
type Order struct {
	OrderId                   int64     `json:"order_id" gorm:"not null;unique;primarykey" form:"orderId"`              //订单编号
	CreatedOrderTime          time.Time `json:"created_order_time" form:"createOrderTime"`                              //下单时间
	UserId                    uint      `form:"userId" gorm:"not null" json:"user_id"`                                  //用户id
	Username                  string    `json:"username" form:"username" gorm:"not null"`                               //用户名
	Address                   string    `form:"address" gorm:"not null" json:"address"`                                 //收货地址
	BusinessId                uint      `form:"businessId" gorm:"not null" json:"business_id"`                          //商家id
	GoodsId                   int64     `json:"goods_id" gorm:"not null" form:"goodsId"`                                //商品id
	GoodsTitle                string    `gorm:"not null" form:"goodsTitle" json:"goods_title"`                          //商品标题
	GoodsImage                string    `gorm:"not null" form:"goodsImage" json:"goods_image_path"`                     //商品图片
	GoodsName                 string    `gorm:"not null" form:"goodsName" json:"goods_name"`                            //商品名称
	GoodsTag                  string    `gorm:"not null" form:"goodsTag" json:"goods_tag"`                              //商品标签
	GoodsParentTag            string    `json:"goods_parent_tag" form:"goodsParentTag"`                                 //商品父级标签
	GoodsExplain              string    `gorm:"not null" json:"goods_explain" form:"goodsExplain"`                      //商品介绍
	GoodsPrice                uint      `gorm:"not null" json:"goods_price" form:"goodsPrice"`                          //商品价格
	GoodsModel                string    `gorm:"not null" json:"goods_model" form:"goodsModel"`                          //商品型号
	OrderType                 uint      `json:"order_type" form:"orderType"  gorm:"not null;default:0"`                 //订单类型:0--普通订单;1--限时秒杀;2--优惠券订单
	GoodsQuantity             uint      `form:"goodsQuantity" gorm:"not null" json:"goods_quantity"`                    //购买数量
	SumMoney                  uint      `form:"sumMoney" gorm:"not null" json:"sum_money"`                              //付款总价
	OrderStatus               uint      `json:"order_status" form:"orderStatus" gorm:"not null;default:0"`              //订单状态:0--未付款;1--未发货;2--未签收;3--已签收;
	OrderGetGoodsTime         time.Time `json:"order_get_goods_time" form:"orderGetGoodsTime"`                          //订单签收时间
	OrderReturnStatus         uint      `json:"order_return_status" form:"orderReturnStatus" gorm:"not null;default:0"` //订单退货状态:0--未退货;1--申请退货;2--退货;3--退货完毕
	OrderReturnBeginTime      time.Time `json:"order_return_begin_time" form:"orderReturnBeginTime"`                    //订单申请退货时间
	OrderReturnFinishTime     time.Time `json:"order_return_finish_time" form:"orderReturnFinishTime"`                  //退货完成时间
	OrderFirstEvaluate        string    `form:"orderFirstEvaluate" json:"order_first_evaluate"`                         //订单首次评论
	OrderFirstEvaluateImages  string    `form:"orderFirstEvaluateImages" json:"order_first_evaluate_images"`            //订单首次评论图片
	OrderAppendEvaluate       string    `form:"orderAppendEvaluate" json:"order_append_evaluate"`                       //订单追加评论
	OrderAppendEvaluateImages string    `form:"orderAppendEvaluateImages" json:"order_append_evaluate_images"`          //订单追加评论图片
}

func init() {
	err := db.DB.AutoMigrate(&Order{})
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

//添加订单
func (o *Order) CreateOrder() (msg string, err error) {
	err = db.DB.Model(&Order{}).Create(&o).Error
	if err != nil {
		return "创建订单失败", err
	}
	return "成功", nil
}

//查看订单
func (o *Order) ShowOrder() (msg string, err error) {
	if err := db.DB.Model(&Order{}).Where("order_id=?", o.OrderId).Find(&o).Error; err != nil {
		return "查找失败", err
	}
	return "查找成功", nil
}

//修改订单状态
func (o *Order) UpdateOrderStatus() (msg string, err error) {
	if err := db.DB.Model(&Order{}).Where("order_id=?", o.OrderId).Updates(map[string]interface{}{"order_status": o.OrderStatus}).Error; err != nil {
		return "订单状态修改错误", err
	}
	return "订单状态修改成功", nil
}
