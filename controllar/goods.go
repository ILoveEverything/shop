package controllar

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
	"shop/model"
	"shop/utils"
	"strings"
	"time"
)

var Goods model.GoodsDetails

//返回前端商品信息
type GoodsInfoPage struct {
	BusinessId     uint   `form:"businessId" gorm:"not null" json:"business_id"`      //商家id
	GoodsId        int64  `json:"goods_id" gorm:"not null;unique" form:"goodsId"`     //商品id
	GoodsTitle     string `gorm:"not null" form:"goodsTitle" json:"goods_title"`      //商品标题
	GoodsImage     string `gorm:"not null" form:"goodsImage" json:"goods_image_path"` //商品图片
	GoodsName      string `gorm:"not null" form:"goodsName" json:"goods_name"`        //商品名称
	GoodsTag       string `gorm:"not null" form:"goodsTag" json:"goods_tag"`          //商品标签
	GoodsParentTag string `json:"goods_parent_tag" form:"goodsParentTag"`             //商品父级标签
	GoodsExplain   string `gorm:"not null" json:"goods_explain" form:"goodsExplain"`  //商品介绍
	GoodsPrice     uint   `gorm:"not null" json:"goods_price" form:"goodsPrice"`      //商品价格
	GoodsModel     string `gorm:"not null" json:"goods_model" form:"goodsModel"`      //商品型号
}

//浏览商品
func GoodsList(c *gin.Context) {
	var GoodsListArray []model.GoodsList
	var GoodsList model.GoodsList
	data, msg, err := Goods.ShowAllGoodsDetails()
	if err != nil {
		utils.ReturnJson(c, msg, 400, nil)
		return
	}
	for _, v := range data {
		GoodsList.GoodsId = v.GoodsId
		GoodsList.GoodsTitle = v.GoodsTitle
		GoodsList.GoodsImage = v.GoodsImage
		GoodsListArray = append(GoodsListArray, GoodsList)
	}
	utils.ReturnJson(c, "成功", 200, GoodsListArray)
}

//商品分类
func GoodsMenu(c *gin.Context) {
	var GoodsMenuArray []model.GoodsMenu
	var GoodsMenu model.GoodsMenu
	data, msg, err := Goods.ShowAllGoodsDetails()
	if err != nil {
		utils.ReturnJson(c, msg, 400, nil)
		return
	}
	for _, v := range data {
		GoodsMenu.GoodsId = v.GoodsId
		GoodsMenu.GoodsTitle = v.GoodsTitle
		GoodsMenu.GoodsTitle = v.GoodsTitle
		GoodsMenu.GoodsTag = v.GoodsTag
		GoodsMenu.GoodsParentTag = v.GoodsParentTag
		GoodsMenu.GoodsImage = v.GoodsImage
		GoodsMenuArray = append(GoodsMenuArray, GoodsMenu)
	}
	utils.ReturnJson(c, "成功", 200, GoodsMenuArray)
}

//商品详情
func GoodsDetails(c *gin.Context) {
	if err := c.Bind(&Goods); err != nil {
		utils.ReturnJson(c, "查看失败", 400, nil)
		return
	}
	msg, err := Goods.ShowGoodsDetails()
	if err != nil {
		utils.ReturnJson(c, msg, 400, nil)
		return
	}
	var GoodsInfo GoodsInfoPage
	GoodsInfo.BusinessId = Goods.BusinessId
	GoodsInfo.GoodsId = Goods.GoodsId
	GoodsInfo.GoodsTitle = Goods.GoodsTitle
	GoodsInfo.GoodsImage = Goods.GoodsImage
	GoodsInfo.GoodsName = Goods.GoodsName
	GoodsInfo.GoodsTag = Goods.GoodsTag
	GoodsInfo.GoodsParentTag = Goods.GoodsParentTag
	GoodsInfo.GoodsExplain = Goods.GoodsExplain
	GoodsInfo.GoodsPrice = Goods.GoodsPrice
	GoodsInfo.GoodsModel = Goods.GoodsModel
	utils.ReturnJson(c, msg, 200, GoodsInfo)
}

//添加商品
func CreateGoods(c *gin.Context) {
	var (
		msg string
		err error
	)
	var goods model.GoodsDetails
	if err = c.Bind(&goods); err != nil {
		fmt.Println(err)
		utils.ReturnJson(c, "上传失败", 400, nil)
		return
	}
	//判断上传信息是否完整
	if len(goods.GoodsTitle) == 0 {
		utils.ReturnJson(c, "商品标题不能为空", 400, nil)
		return
	}
	if len(goods.GoodsImage) == 0 {
		utils.ReturnJson(c, "请输入商品图片", 400, nil)
		return
	}
	if len(goods.GoodsName) == 0 {
		utils.ReturnJson(c, "商品名称不能为空", 400, nil)
		return
	}
	if len(goods.GoodsTag) == 0 {
		utils.ReturnJson(c, "商品标签不能为空", 400, nil)
		return
	}
	if len(goods.GoodsExplain) == 0 {
		utils.ReturnJson(c, "商品介绍为空", 400, nil)
		return
	}
	if len(goods.GoodsModel) == 0 {
		utils.ReturnJson(c, "商品型号为空", 400, nil)
		return
	}
	if goods.GoodsQuantity <= 0 {
		utils.ReturnJson(c, "商品库存不能为空", 400, nil)
		return
	}
	if goods.GoodsPrice <= 0 {
		utils.ReturnJson(c, "请重新输入正确的价格", 400, nil)
		return
	}
	if goods.GoodsStatus < 0 || goods.GoodsStatus > 5 {
		utils.ReturnJson(c, "请重新输入商品状态编号", 400, nil)
		return
	}
	//获取商家id
	token := c.Request.Header.Get("Authorization")
	claims, _ := utils.ParseToken(token)
	goods.BusinessId = claims.UID
	//创建商品编号
	now := time.Now()
	goods.GoodsId = now.UnixNano()
	//添加商品到商品详情表
	msg, err = goods.CreateGoodsDetails()
	if err != nil {
		utils.ReturnJson(c, msg, 400, nil)
		return
	}
	utils.ReturnJson(c, msg, 200, nil)
}

//上传图片
func AddGoodsImages(c *gin.Context) {
	file, err := c.FormFile("images")
	if err != nil {
		utils.ReturnJson(c, "图片上传失败", 400, nil)
		return
	}
	//判断文件后缀
	fileExt := strings.ToLower(path.Ext(file.Filename))
	if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
		utils.ReturnJson(c, "上传失败!只允许png,jpg,gif,jpeg文件", 400, nil)
		return
	}
	now := time.Now()
	nowTime := fmt.Sprint(now.UnixNano())
	dst := "./assets/goodsImages/" + nowTime + fileExt
	fmt.Println(dst)
	err1 := c.SaveUploadedFile(file, dst)
	if err1 != nil {
		fmt.Println(err1)
		utils.ReturnJson(c, "文件保存失败", 400, nil)
		return
	}
	utils.ReturnJson(c, "成功", 200, map[string]interface{}{"GoodsImage": dst[1:]})
}

//更改商品信息 todo
func UpdateGoods(c *gin.Context) {

}
