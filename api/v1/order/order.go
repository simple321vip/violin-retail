package order

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"strconv"
	"time"
	"violin-home.cn/retail/common"
	"violin-home.cn/retail/common/logs"
	"violin-home.cn/retail/models"
	"violin-home.cn/retail/store"
)

type Handler struct {
}

// GetOrderList 获取订单一览
// *
func (oh *Handler) GetOrderList(c *gin.Context) {
	//unit := c.Param("unit")
	start := time.Now().Add(-time.Hour * 24 * 30)
	result := &common.Result{}
	DataBase := common.GetTenantDateBase(c)
	collection := store.ClientMongo.Database(DataBase).Collection(common.Order)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	unwind := bson.M{
		"$unwind": bson.M{
			"path":                       "$Customer", // 将主表查询结果和从表查询结果1对1关联
			"preserveNullAndEmptyArrays": true,        // 空数组记录保留，不会丢失主表数据
		},
	}

	// outer left join
	lookup := bson.M{
		"$lookup": bson.M{
			"from":         "customer",   // 关联表 color
			"localField":   "CustomerID", // 主表 关联字段
			"foreignField": "_id",        // 关联表color 关联字段
			"as":           "Customer",   // 查询后返回结果名称，一对多，该结果为数组，当使用unwind时候，变成1对1形式，变成对象
		},
	}

	match := bson.M{
		"$match": bson.M{
			"OrderTime": bson.M{
				"$gt": start,
			},
		},
	}

	pipeline := []bson.M{lookup, unwind, match}

	// 执行聚合查询
	cursor, err := collection.Aggregate(context.Background(), pipeline)
	if err == nil {
		var orders []models.Order
		for cursor.Next(ctx) {
			var order models.Order
			err := cursor.Decode(&order)
			if err != nil {
				logs.LG.Error(err.Error())
				return
			}
			orders = append(orders, order)
		}
		c.JSON(http.StatusOK, result.Success(orders))
	}
}

// GetOrder 获取订单明细
// *
func (oh *Handler) GetOrder(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("orderID"), 10, 64)
	result := &common.Result{}
	gh := oh.GetHandler()
	one, err := gh.FindOne(int(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail(500, ""))
		return
	}

	//customer, err := common.FindOne[models.Customer](DataBase, common.Customer, order.CustomerId)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, result.Fail(500, ""))
	//	return
	//}
	//order.Customer = &customer

	c.JSON(http.StatusOK, result.Success(one))
}

// CreateOrder 创建订单
// *
func (oh *Handler) CreateOrder(c *gin.Context) {
	result := &common.Result{}
	DataBase := common.GetTenantDateBase(c)
	var order models.Order
	err := c.ShouldBind(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	gh := oh.GetHandler()

	id, err := gh.GetNextID()
	if err != nil {
		return
	}

	sup := models.Order{
		ID:                       id + 1,
		OrderTime:                time.Now(),
		CustomerId:               order.Customer.ID,
		OrderType:                order.OrderType,
		OrderProducts:            order.OrderProducts,
		AccountsReceivable:       order.AccountsReceivable,
		ActualAccountsReceivable: order.ActualAccountsReceivable,
		Refund:                   order.Refund,
		ActualRefund:             order.ActualRefund,
		Freight:                  order.Freight,
		Comment:                  order.Comment,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 开启事务
	if session, err := store.StartTransaction(); err == nil {
		// 执行事务
		err = mongo.WithSession(context.Background(), session, func(sessionContext mongo.SessionContext) error {

			// 商品库存修改
			productColl := store.ClientMongo.Database(DataBase).Collection("product")
			for _, orderProduct := range order.OrderProducts {
				// 1. 定义查询条件
				filter := bson.D{{"_id", orderProduct.ID}}

				// 2. 获取该商品
				rst := productColl.FindOne(ctx, filter, options.FindOne())
				var product models.Product
				if err := rst.Decode(&product); err != nil {
					logs.LG.Error(err.Error())
					return err
				}

				// 3. 计算库存
				product.StockQuantity -= orderProduct.Quantity

				// 4. 定义更新操作
				update := bson.D{{"$set", bson.M{
					"StockQuantity": product.StockQuantity,
				}}}

				// 5. 更新库存
				_, err = productColl.UpdateOne(ctx, filter, update)
				if err != nil {
					logs.LG.Error(err.Error())
					return err
				}
			}

			// 创建订单记录
			collection := store.ClientMongo.Database(DataBase).Collection("order")

			bsonData, err := bson.Marshal(sup)

			if err != nil {
				logs.LG.Error(err.Error())
				return err
			}
			_, err = collection.InsertOne(ctx, bsonData)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return
		}

		// 提交事务
		err = store.CommitTransaction(session)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, result.Success(sup.ID))
	}

}

// CancelOrder 取消订单
// *
func (oh *Handler) CancelOrder(c *gin.Context) {
	//result := &common.Result{}
	//DataBase := common.GetTenantDateBase(c)
	//id := 1

	//gh := oh.GetHandler()
	//
	//order, err := common.FindOne[models.Order](DataBase, "order", id)
	//if err != nil {
	//	c.JSON(http.StatusOK, result.Success("未发现订单信息"))
	//	return
	//}
	//if order.IsCancel == true {
	//	c.JSON(http.StatusOK, result.Success("该订单已被取消"))
	//	return
	//}

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
	//		for _, orderProduct := range order.OrderProducts {
	//			// 1. 定义查询条件
	//			filter := bson.D{{"_id", orderProduct.ID}}
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
	//			product.StockQuantity += orderProduct.Quantity
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
	//		// 创建订单记录
	//		collection := store.ClientMongo.Database(DataBase).Collection("order")
	//
	//		// 设置订单取消
	//		order.IsCancel = true
	//
	//		// 1. 定义查询条件
	//		filter := bson.D{{"_id", id}}
	//
	//		// 2. 定义更新操作
	//		update := bson.D{{"$set", order}}
	//
	//		collection.FindOneAndUpdate(ctx, filter, update)
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

// UpdateOrder 变更订单
// *
func (oh *Handler) UpdateOrder(c *gin.Context) {
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

func (oh *Handler) GetHandler() *common.BaseHandler {
	gh := &common.BaseHandler{
		DatabaseName:   "test",
		CollectionName: "order",
	}
	return gh
}
