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

	// 多表查询
	// https://www.mebugs.com/post/mgolookup.html
	// 聚合查询
	// All stages except the $out, $merge, $geoNear, $changeStream, and $changeStreamSplitLargeEvent stages
	// can appear multiple times in a pipeline.
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

	//result := &common.Result{}
	//var err error
	//defer func() {
	//	if err != nil {
	//		logs.LG.Error(err.Error())
	//		c.JSON(http.StatusInternalServerError, nil)
	//	}
	//}()

	// 一个租户一个数据库
	//DataBase := common.GetTenantDateBase(c)
	//ID, _ := common.GetNextID(DataBase, "product")
	//Product := &models.Product{
	//	ID:            ID + 1,
	//	Name:          "实木地板",
	//	RetailPrice:   120.5,
	//	Unit:          "张",
	//	StockQuantity: 110,
	//	SizeID:        1,
	//}
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	//defer cancel()

	//// 开启事务
	//session, err := store.StartTransaction()
	//if err != nil {
	//	return
	//}
	//
	//// 执行事务
	//err = mongo.WithSession(context.Background(), session, func(sessionContext mongo.SessionContext) error {
	//
	//	collection = store.ClientMongo.Database(DataBase).Collection("product")
	//
	//	bsonData, err := bson.Marshal(Product)
	//
	//	if err != nil {
	//		return err
	//	}
	//	_, err = collection.InsertOne(ctx, bsonData)
	//	if err != nil {
	//		return err
	//	}
	//	return nil
	//})
	//
	//if err != nil {
	//	return
	//}
	//
	//// 提交事务
	//err = store.CommitTransaction(session)
	//if err != nil {
	//	return
	//}
	//
	//c.JSON(http.StatusOK, result.Success("success"))
}
