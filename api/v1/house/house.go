package house

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

// GetHouseList 获取出入库一览
// *
func (hh *Handler) GetHouseList(c *gin.Context) {
	result := &common.Result{}
	DataBase := common.GetTenantDateBase(c)
	houseCol := store.ClientMongo.Database(DataBase).Collection("house")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if find, err := houseCol.Find(ctx, bson.D{}); err == nil {
		var houses []models.House
		for find.Next(ctx) {
			var house models.House
			err := find.Decode(&house)
			if err != nil {
				logs.LG.Error(err.Error())
				return
			}
			houses = append(houses, house)
		}
		c.JSON(http.StatusOK, result.Success(houses))
	}
}

// GetHouse 获取指定出入库明细
// *
func (hh *Handler) GetHouse(c *gin.Context) {
	result := &common.Result{}
	DataBase := common.GetTenantDateBase(c)
	houseCol := store.ClientMongo.Database(DataBase).Collection("house")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if find, err := houseCol.Find(ctx, bson.D{}); err == nil {
		var houses []models.House
		for find.Next(ctx) {
			var house models.House
			err := find.Decode(&house)
			if err != nil {
				logs.LG.Error(err.Error())
				return
			}
			houses = append(houses, house)
		}
		c.JSON(http.StatusOK, result.Success(houses))
	}
}

// HouseIn 入库
// *
func (hh *Handler) HouseIn(c *gin.Context) {
	//result := &common.Result{}
	//DataBase := common.GetTenantDateBase(c)
	//
	//var houseProducts []models.HouseProduct
	//houseProducts = append(houseProducts, models.HouseProduct{
	//	ProductID:  1,
	//	Quantity:   200,
	//	HousePrice: 2000,
	//}, models.HouseProduct{
	//	ProductID:  2,
	//	Quantity:   300,
	//	HousePrice: 3000,
	//})
	//
	//id, err := common.GetNextID(DataBase, "house")
	//if err != nil {
	//	return
	//}
	//sup := models.House{
	//	ID:              id + 1,
	//	HouseTime:       time.Time{},
	//	HouseType:       0,
	//	HouseProduct:    houseProducts,
	//	AccountsPayable: 0,
	//	ActualPaid:      0,
	//	Freight:         0,
	//	Comment:         "",
	//}
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	//defer cancel()
	//
	//// 开启事务
	//if session, err := store.StartTransaction(); err == nil {
	//	// 执行事务
	//	err = mongo.WithSession(context.Background(), session, func(sessionContext mongo.SessionContext) error {
	//
	//		// 商品库存修改
	//		productColl := store.ClientMongo.Database(DataBase).Collection("product")
	//		for _, houseProduct := range houseProducts {
	//			// 1. 定义查询条件
	//			filter := bson.D{{"_id", houseProduct.ProductID}}
	//
	//			// 2. 获取该商品
	//			rst := productColl.FindOne(ctx, filter, options.FindOne())
	//			var product models.Product
	//			if err := rst.Decode(&product); err != nil {
	//				logs.LG.Error(err.Error())
	//				return err
	//			}
	//
	//			// 3. 计算库存
	//			product.StockQuantity += houseProduct.Quantity
	//
	//			// 4. 定义更新操作
	//			update := bson.D{{"$set", bson.M{
	//				"StockQuantity": product.StockQuantity,
	//			}}}
	//
	//			// 5. 更新库存
	//			_, err = productColl.UpdateOne(ctx, filter, update)
	//			if err != nil {
	//				logs.LG.Error(err.Error())
	//				return err
	//			}
	//		}
	//
	//		// 插入出入库记录
	//		collection := store.ClientMongo.Database(DataBase).Collection("house")
	//
	//		bsonData, err := bson.Marshal(sup)
	//
	//		if err != nil {
	//			logs.LG.Error(err.Error())
	//			return err
	//		}
	//		_, err = collection.InsertOne(ctx, bsonData)
	//		if err != nil {
	//			return err
	//		}
	//		return nil
	//	})
	//	if err != nil {
	//		return
	//	}
	//
	//	// 提交事务
	//	err = store.CommitTransaction(session)
	//	if err != nil {
	//		return
	//	}
	//	c.JSON(http.StatusOK, result.Success("success"))
	//}
}
