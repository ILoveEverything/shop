package model

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"os"
	"shop/db"
)

//商品浏览表
type GoodsList struct {
	GoodsId    int64  `gorm:"not null" json:"goods_id" form:"goodsId"`            //商品id
	GoodsTitle string `gorm:"not null" form:"goodsTitle" json:"goods_title"`      //商品标题
	GoodsImage string `gorm:"not null" form:"goodsImage" json:"goods_image_path"` //商品图片
}

//商品详情表
type GoodsDetails struct {
	gorm.Model
	BusinessId     uint   `form:"businessId" gorm:"not null" json:"business_id"`       //商家id
	GoodsId        int64  `json:"goods_id" gorm:"not null;unique" form:"goodsId"`      //商品id
	GoodsTitle     string `gorm:"not null" form:"goodsTitle" json:"goods_title"`       //商品标题
	GoodsImage     string `gorm:"not null" form:"goodsImage" json:"goods_image_path"`  //商品图片
	GoodsName      string `gorm:"not null" form:"goodsName" json:"goods_name"`         //商品名称
	GoodsTag       string `gorm:"not null" form:"goodsTag" json:"goods_tag"`           //商品标签
	GoodsParentTag string `json:"goods_parent_tag" form:"goodsParentTag"`              //商品父级标签
	GoodsExplain   string `gorm:"not null" json:"goods_explain" form:"goodsExplain"`   //商品介绍
	GoodsPrice     uint   `gorm:"not null" json:"goods_price" form:"goodsPrice"`       //商品价格
	GoodsModel     string `gorm:"not null" json:"goods_model" form:"goodsModel"`       //商品型号
	GoodsQuantity  uint   `gorm:"not null" json:"goods_quantity" form:"goodsQuantity"` //商品库存
	GoodsStatus    int    `gorm:"not null" json:"goods_status" form:"goodsStatus"`     //商品状态:0--未售;1--预售;2--在售;3--停售;4--下架;5--删除
}

//商品分类
type GoodsMenu struct {
	GoodsId        int64  `gorm:"not null" json:"goods_id" form:"goodsId"`            //商品id
	GoodsTitle     string `gorm:"not null" form:"goodsTitle" json:"goods_title"`      //商品标题
	GoodsTag       string `gorm:"not null" form:"goodsTag" json:"goods_tag"`          //商品标签
	GoodsParentTag string `json:"goods_parent_tag" form:"goodsParentTag"`             //商品父级标签
	GoodsImage     string `gorm:"not null" form:"goodsImage" json:"goods_image_path"` //商品图片
}

func init() {
	if err := db.DB.AutoMigrate(&GoodsDetails{}); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

//获取单个商品详情表数据
func (g *GoodsDetails) ShowGoodsDetails() (msg string, err error) {
	if err := db.DB.Model(&GoodsDetails{}).Where("goods_id=?", g.GoodsId).Find(&g).Error; err != nil {
		return "查找失败", err
	}
	if g.ID == 0 {
		return "未查到", errors.New("未查到")
	}
	return "查找成功", nil
}

//获取所有商品详情表数据
func (g *GoodsDetails) ShowAllGoodsDetails() (data []*GoodsDetails, msg string, err error) {
	var AllGoods []*GoodsDetails
	if err := db.DB.Model(&GoodsDetails{}).Find(&AllGoods).Error; err != nil {
		return nil, "查找失败", err
	}
	return AllGoods, "查找成功", nil
}

//添加商品到商品详情表
func (g *GoodsDetails) CreateGoodsDetails() (msg string, err error) {
	if err = db.DB.Model(&GoodsDetails{}).Create(&g).Error; err != nil {
		return "商品添加失败", err
	}
	return "商品添加成功", nil
}

//更新商品库存
func (g *GoodsDetails) UpdateGoodsQuantity() (msg string, err error) {
	err = db.DB.Model(&GoodsDetails{}).Where("id=?", g.ID).Updates(map[string]interface{}{"goods_quantity": g.GoodsQuantity}).Error
	if err != nil {
		return "更改失败", err
	}
	return "成功", nil
}
