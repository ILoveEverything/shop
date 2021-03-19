package controllar

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"shop/db"
	"shop/model"
	"shop/utils"
	"strconv"
	"sync"
	"time"
)

var (
	suMapLuck sync.Mutex
	ctx       = context.TODO()
	SK        model.SecondKillGoodsDetails
)

type UserSecondKill struct {
	GoodsId   int64  `form:"goods_id" gorm:"comment:商品id;not null" json:"goods_id"`  //商品id
	Address   string `json:"address" form:"address" gorm:"comment:收货地址;"`            //收货地址
	OrderType uint   `json:"order_type" form:"order_type" gorm:"not null;default:0"` //订单类型:0--普通订单;1--限时秒杀;2--优惠券订单
}

//创建秒杀活动
func CreateSecondKill(c *gin.Context) {
	if err := c.Bind(&SK); err != nil {
		utils.ReturnJson(c, "Bind失败", 400, nil)
		return
	}
	//解析商家id
	token := c.Request.Header.Get("authorization")
	claims, _ := utils.ParseToken(token)
	SK.BusinessId = claims.UID
	//创建秒杀产品
	msg, err := SK.Create()
	msg, err = model.CreateRedis(&SK)
	if err != nil {
		utils.ReturnJson(c, msg, 400, nil)
		return
	}
	utils.ReturnJson(c, msg, 200, nil)
}

//秒杀
func SecondKill(c *gin.Context) {
	var suMap = make(map[uint]int)
	var userSK UserSecondKill
	if err := c.Bind(&userSK); err != nil {
		utils.ReturnJson(c, "Bind失败", 400, nil)
		return
	}
	//获取订单类型
	if userSK.OrderType != 1 {
		utils.ReturnJson(c, "订单类型不对", 400, nil)
		return
	}
	//获取用户id
	token := c.Request.Header.Get("Authorization")
	claims, _ := utils.ParseToken(token)
	for {
		suMapLuck.Lock()
		if suMap[claims.UID] == 1 {
			suMapLuck.Unlock()
			time.Sleep(time.Millisecond)
			continue
		}
		suMap[claims.UID] = 1
		suMapLuck.Unlock()
		break
	}
	defer delete(suMap, claims.UID)
	//读取已秒杀用户id
	get, err := db.RedisDB.LRange(ctx, strconv.FormatInt(userSK.GoodsId, 10), 0, -1).Result()
	if err != nil {
		fmt.Println(err)
		//第一次查询无内容会报错
		//utils.ReturnJson(c, "读redis存储秒杀记录-->秒杀失败", 400, nil)
		//return
	}
	for _, v := range get {
		if v == strconv.Itoa(int(claims.UID)) {
			utils.ReturnJson(c, "不能再次参与秒杀", 400, nil)
			return
		}
	}
	//抛出秒杀资格
	var s model.SecondKillGoodsDetails
	s.GoodsId = userSK.GoodsId
	result, err := db.RedisDB.LPop(ctx, "secondKill:"+strconv.FormatInt(s.GoodsId, 10)).Result()
	if err != nil {
		utils.ReturnJson(c, "抛出秒杀资格-->秒杀失败", 400, nil)
		return
	}
	if result != "" {
		//记录已秒杀的用户 todo 修改活动时间
		_, err := db.RedisDB.LPush(ctx, strconv.FormatInt(userSK.GoodsId, 10), claims.UID).Result()
		if err != nil {
		returnAuthority1:
			err := returnSecondKillAuthority(s, result)
			if err != nil {
				goto returnAuthority1
			}
			utils.ReturnJson(c, "记录已秒杀的用户出错-->秒杀失败", 400, nil)
			return
		}
		//秒杀订单
		var secondKillOrder model.Order
		//创建订单id
		now := time.Now()
		secondKillOrder.OrderId = now.UnixNano()
		//创建订单时间
		secondKillOrder.CreatedOrderTime = now
		secondKillOrder.UserId = claims.UID
		//获取用户名
		var userInfo model.UserInfo
		userInfo.ID = secondKillOrder.UserId
		msg, err := userInfo.RetrieveUserInfo()
		if err != nil {
		returnAuthority2:
			err := returnSecondKillAuthority(s, result)
			if err != nil {
				goto returnAuthority2
			}
		delLog1:
			err = delSecondKillUserLog(claims, userSK)
			if err != nil {
				goto delLog1
			}
			utils.ReturnJson(c, "获取用户信息失败-->秒杀失败", 400, nil)
			return
		}
		secondKillOrder.Username = userInfo.Username
		//获取收货地址
		if len(userSK.Address) == 0 {
			utils.ReturnJson(c, "收货地址不能为空", 400, nil)
			return
		}
		secondKillOrder.Address = userSK.Address
		//获取商家id
		var goods model.GoodsDetails
		secondKillOrder.GoodsId = userSK.GoodsId
		goods.GoodsId = secondKillOrder.GoodsId
		msg, err = goods.ShowGoodsDetails()
		if err != nil {
		returnAuthority3:
			err := returnSecondKillAuthority(s, result)
			if err != nil {
				goto returnAuthority3
			}
		delLog2:
			err = delSecondKillUserLog(claims, userSK)
			if err != nil {
				goto delLog2
			}
			utils.ReturnJson(c, "获取商品信息失败-->秒杀失败", 400, nil)
			return
		}
		secondKillOrder.BusinessId = goods.BusinessId
		//获取商品id
		secondKillOrder.GoodsId = userSK.GoodsId
		//获取商品标题
		secondKillOrder.GoodsTitle = goods.GoodsTitle
		//获取商品图片
		secondKillOrder.GoodsImage = goods.GoodsImage
		//获取商品名称
		secondKillOrder.GoodsName = goods.GoodsName
		//获取商品标签
		secondKillOrder.GoodsTag = goods.GoodsTag
		//获取商品父级标签
		secondKillOrder.GoodsParentTag = goods.GoodsParentTag
		//获取商品介绍
		secondKillOrder.GoodsExplain = goods.GoodsExplain
		//获取商品价格
		secondKillOrder.GoodsPrice = goods.GoodsPrice
		//获取商品类型
		secondKillOrder.GoodsModel = goods.GoodsModel
		//订单类型
		secondKillOrder.OrderType = userSK.OrderType
		//订单商品数量
		secondKillOrder.GoodsQuantity = 1
		//订单商品金额
		secondKillOrder.SumMoney = goods.GoodsPrice
		//创建订单
		msg, err = secondKillOrder.CreateOrder()
		if err != nil {
		returnAuthority4:
			err := returnSecondKillAuthority(s, result)
			if err != nil {
				goto returnAuthority4
			}
		delLog3:
			err = delSecondKillUserLog(claims, userSK)
			if err != nil {
				goto delLog3
			}
			utils.ReturnJson(c, msg, 400, &secondKillOrder)
			return
		}
		utils.ReturnJson(c, msg, 200, &secondKillOrder)
	}
}

//返回抛出的秒杀资格
func returnSecondKillAuthority(s model.SecondKillGoodsDetails, result string) error {
	_, err := db.RedisDB.LPush(ctx, "secondKill:"+strconv.FormatInt(s.GoodsId, 10), result).Result()
	return err
}

//删除已记录的秒杀用户id
func delSecondKillUserLog(claims *utils.Claims, userSK UserSecondKill) error {
	_, err := db.RedisDB.LRem(ctx, strconv.FormatInt(userSK.GoodsId, 10), 1, claims.UID).Result()
	return err
}
