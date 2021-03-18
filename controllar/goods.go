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

//定义商品状态
const (
	Unsold = iota
	PreSale
	InStock
	Discontinued
	OffShelf
)

var GoodsStatus = map[int]string{
	Unsold:       "未售",
	PreSale:      "预售",
	InStock:      "在售",
	Discontinued: "停售",
	OffShelf:     "下架",
}

type GoodInfo struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

//浏览商品
func GoodsList(c *gin.Context) {
	var GoodsListArray model.GoodsList
	data, err := GoodsListArray.ShowGoodsList()
	if err != nil {
		utils.ReturnJson(c, "查看失败,请重试", 400, nil)
		return
	}
	utils.ReturnJson(c, "成功", 200, data)
}

//商品分类
func GoodsMenu(c *gin.Context) {
	var GoodsMenuArray model.GoodsMenu
	data, err := GoodsMenuArray.ShowGoodsMenu()
	if err != nil {
		utils.ReturnJson(c, "失败,请重试", 400, nil)
		return
	}
	utils.ReturnJson(c, "成功", 200, data)
}

//商品详情
func GoodsDetails(c *gin.Context) {
	var GoodsDetailsStruct model.GoodsDetails
	if err := c.Bind(&GoodsDetailsStruct); err != nil {
		utils.ReturnJson(c, "查看失败", 400, nil)
		return
	}
	//fmt.Println(GoodsDetailsStruct)
	msg, err := GoodsDetailsStruct.ShowGoodsDetails()
	if err != nil {
		utils.ReturnJson(c, msg, 400, nil)
		return
	}
	//fmt.Println(GoodsDetailsStruct)
	utils.ReturnJson(c, msg, 200, GoodsDetailsStruct)
}

//上传商品
func UpdateGoods(c *gin.Context) {
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
	now := time.Now()
	goods.GoodsId = now.UnixNano()
	//添加商品到商品详情表
	msg, err = goods.CreateGoodsDetails()
	if err != nil {
		utils.ReturnJson(c, msg, 400, nil)
		return
	}
	//添加商品到商品浏览表
	var goodsList model.GoodsList
	goodsList.GoodsId = goods.GoodsId
	goodsList.GoodsTitle = goods.GoodsTitle
	goodsList.GoodsImage = goods.GoodsImage
	msg, err = goodsList.CreateGoodsList()
	if err != nil {
		utils.ReturnJson(c, msg, 400, nil)
		return
	}
	//添加商品到商品分类表
	var goodsMenu model.GoodsMenu
	goodsMenu.GoodsId = goods.GoodsId
	goodsMenu.GoodsTitle = goods.GoodsTitle
	goodsMenu.GoodsTag = goods.GoodsTag
	goodsMenu.GoodsParentTag = goods.GoodsParentTag
	goodsMenu.GoodsImage = goods.GoodsImage
	msg, err = goodsMenu.CreateGoodsMenu()
	if err != nil {
		utils.ReturnJson(c, msg, 400, nil)
		return
	}
	utils.ReturnJson(c, msg, 200, nil)
}

//上传图片
func UpdateGoodsImages(c *gin.Context) {
	//保存上传图片
	//form, err := c.MultipartForm()
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
	//fmt.Println(now)
	nowTime := fmt.Sprint(now.UnixNano())
	//fmt.Println(nowTime)
	dst := "./assets/goodsImages/" + nowTime + fileExt
	fmt.Println(dst)
	err1 := c.SaveUploadedFile(file, dst)
	if err1 != nil {
		fmt.Println(err1)
		utils.ReturnJson(c, "文件保存失败", 400, nil)
		return
	}
	utils.ReturnJson(c, "成功", 200, map[string]interface{}{"GoodsImage": dst[1:]})
	//images := form.File["images"]
	//for _, v := range images {
	//	dst := "/assets/goodsImages/" + v.Filename
	//	err := c.SaveUploadedFile(v, dst)
	//	if err != nil {
	//		utils.ReturnJson(c, "图片保存失败", 400, nil)
	//		return
	//	}
	//	utils.ReturnJson(c, "图片保存成功", 200, dst)
	//}
}
