package order

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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
	result := &common.Result{}
	orders, err := oh.getOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail(500, "系统内部错误"))
		return
	}
	c.JSON(http.StatusOK, orders)
}

// GetOrder 获取订单明细
// *
func (oh *Handler) GetOrder(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("ID"), 10, 64)
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
	var order = models.NewOrder()
	err := c.ShouldBindJSON(order)
	gh := oh.GetHandler()

	_, err = gh.InsertOne(order)
	if err != nil {
		logs.LG.Error(err.Error())
		c.JSON(http.StatusInternalServerError, result.Fail(500, "系统内部错误"))
		return
	}
	orders, err := oh.getOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail(500, "系统内部错误"))
		return
	}
	c.JSON(http.StatusOK, orders)
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

func (oh *Handler) getOrders() ([]models.Order, error) {
	gh := oh.GetHandler()
	collection := store.ClientMongo.Database(gh.DatabaseName).Collection(gh.CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	find, err := collection.Find(ctx, bson.D{})
	if err != nil {
		logs.LG.Error(err.Error())
		return nil, err
	}
	var orders = make([]models.Order, 0)
	for find.Next(ctx) {
		var order models.Order
		err := find.Decode(&order)
		if err != nil {
			logs.LG.Error(err.Error())
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (oh *Handler) GetHandler() *common.BaseHandler {
	gh := &common.BaseHandler{
		DatabaseName:   "test",
		CollectionName: "order",
	}
	return gh
}
