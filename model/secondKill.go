package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"os"
	"shop/db"
	"strconv"
	"time"
)

//秒杀详情表
type SecondKillGoodsDetails struct {
	gorm.Model
	BusinessId    uint      `form:"business_id" gorm:"comment:商家id;not null" json:"business_id"`       //商家id
	GoodsId       int64     `form:"goods_id" gorm:"comment:商品id;not null;" json:"goods_id"`            //商品id
	GoodsQuantity uint      `form:"goods_quantity" gorm:"comment:秒杀数量;not null" json:"goods_quantity"` //商品库存
	GoodsStatus   int       `form:"goods_status" gorm:"comment:秒杀状态;not null" json:"goods_status"`     //商品状态:0--未开始;1--已开始;2--结束;
	StartAt       time.Time `json:"start_at" gorm:"comment:开始时间;not null"`                             //活动开始时间
	EndAt         time.Time `json:"end_at" gorm:"comment:结束时间;not null"`                               //活动结束时间
}

func init() {
	if err := db.DB.AutoMigrate(&SecondKillGoodsDetails{}); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

//创建秒杀产品
func (s *SecondKillGoodsDetails) Create() (msg string, err error) {
	if err := db.DB.Model(&SecondKillGoodsDetails{}).Create(&s).Error; err != nil {
		return "添加秒杀商品失败", err
	}
	return "成功", nil
}

//创建秒杀缓存redis列表
func CreateRedis(s *SecondKillGoodsDetails) (msg string, err error) {
	for i := 1; i <= int(s.GoodsQuantity); i++ {
		var ctx = context.TODO()
		_, err := db.RedisDB.RPush(ctx, "secondKill:"+strconv.FormatInt(s.GoodsId, 10), i).Result()
		if err != nil {
			return "创建商品redis秒杀失败", err
		}
	}
	return "成功", nil
}

func (s *SecondKillGoodsDetails) Update() (msg string, err error) {
	err = db.DB.Model(&GoodsDetails{}).Where("goods_id=?", s.GoodsId).Updates(map[string]interface{}{"goods_quantity": s.GoodsQuantity}).Error
	if err != nil {
		return "更改失败", err
	}
	return "成功", nil
}

func (s *SecondKillGoodsDetails) Retrieve() (msg string, err error) {
	if err := db.DB.Model(&GoodsDetails{}).Where("goods_id=?", s.GoodsId).Find(&s).Error; err != nil {
		return "查找失败", err
	}
	return "查找成功", nil
}

func (s *SecondKillGoodsDetails) Delete() (msg string, err error) {
	panic("implement me")
}
