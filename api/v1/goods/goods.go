package goods

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
	"violin-home.cn/retail/common"
	"violin-home.cn/retail/common/logs"
	"violin-home.cn/retail/models"
	"violin-home.cn/retail/store"
)

type Handler struct {
}

// GetGoods 获取商品一栏
// **
func (nh *Handler) GetGoods(c *gin.Context) {
	result := &common.Result{}
	var err error
	defer func() {
		if err != nil {
			logs.LG.Error(err.Error())
			c.JSON(http.StatusInternalServerError, nil)
		}
	}()

	DataBase := common.GetTenantDateBase(c)
	collection := store.ClientMongo.Database(DataBase).Collection("product")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	unwind1 := bson.M{
		"$unwind": bson.M{
			"path":                       "$Color", // 将主表查询结果和从表查询结果1对1关联
			"preserveNullAndEmptyArrays": true,     // 空数组记录保留，不会丢失主表数据
		},
	}

	unwind2 := bson.M{
		"$unwind": bson.M{
			"path":                       "$Size", // 将主表查询结果和从表查询结果1对1关联
			"preserveNullAndEmptyArrays": true,    // 空数组记录保留，不会丢失主表数据
		},
	}

	// outer left join
	lookup1 := bson.M{
		"$lookup": bson.M{
			"from":         "color",   // 关联表 color
			"localField":   "ColorID", // 主表 关联字段
			"foreignField": "_id",     // 关联表color 关联字段
			"as":           "Color",   // 查询后返回结果名称，一对多，该结果为数组，当使用unwind时候，变成1对1形式，变成对象
		},
	}
	lookup2 := bson.M{
		"$lookup": bson.M{
			"from":         "size",
			"localField":   "SizeID",
			"foreignField": "_id",
			"as":           "Size",
		},
	}

	pipeline := []bson.M{lookup1, lookup2, unwind1, unwind2}

	// 执行聚合查询
	cursor, err := collection.Aggregate(context.Background(), pipeline)
	if err == nil {
		var goods []models.Product
		for cursor.Next(ctx) {
			var product models.Product
			err := cursor.Decode(&product)
			if err != nil {
				logs.LG.Error(err.Error())
				return
			}
			goods = append(goods, product)
		}
		c.JSON(http.StatusOK, result.Success(goods))
	}
}

// IncreaseGoods 新增商品
// **
func (nh *Handler) IncreaseGoods(c *gin.Context) {
	//var goods = models.NewGoods()
	//err := c.ShouldBindJSON(goods)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, nil)
	//	return
	//}
	//DataBase := common.GetTenantDateBase(c)
	//ID, _ := common.GetNextID(DataBase, "goods")
	//goods.ID = ID + 1
	//
	//err = common.InsertOne(DataBase, "goods", goods)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, err.Error())
	//	return
	//}
	//one, err := common.FindOne[models.Goods](DataBase, "goods", goods.ID)
	//if err != nil {
	//	return
	//}
	//c.JSON(http.StatusOK, one)
}

// IncreaseGoodType 新增分类
// **
func (nh *Handler) IncreaseGoodType(c *gin.Context) {
	//var goodType = models.NewGoodType()
	//err := c.ShouldBindJSON(&goodType)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, nil)
	//	return
	//}
	//DataBase := common.GetTenantDateBase(c)
	//ID, _ := common.GetNextID(DataBase, "goodType")
	//goodType.ID = ID + 1
	//
	//err = common.InsertOne(DataBase, "goodType", goodType)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, err.Error())
	//	return
	//}
	//some, err := common.Find[models.GoodType](DataBase, "goodType", bson.D{})
	//if err != nil {
	//	return
	//}
	//c.JSON(http.StatusOK, some)
}

// DeleteGoodType 删除分类
// **
func (nh *Handler) DeleteGoodType(c *gin.Context) {
	//gh := &common.Handler{
	//	DatabaseName:   "test",
	//	CollectionName: "goodType",
	//}
	//_, err := gh.DeleteOne(c)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, err.Error())
	//	return
	//}
	//
	//find, err := gh.Find(bson.D{})
	//if err != nil {
	//	return
	//}
	//c.JSON(http.StatusOK, find)
}
