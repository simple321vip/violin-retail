package order

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
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

// GetOrder 获取订单明细
// *
func (oh *Handler) GetOrder(c *gin.Context) {
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

// CreateOrder 创建订单
// *
func (oh *Handler) CreateOrder(c *gin.Context) {
	result := &common.Result{}
	DataBase := common.GetTenantDateBase(c)
	var order models.Order
	err := c.BindJSON(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	id, err := common.GetNextID(DataBase, "order")
	if err != nil {
		return
	}
	var orderProducts []models.OrderProduct
	orderProducts = append(orderProducts, models.OrderProduct{
		ProductID:  1,
		Quantity:   200,
		Price:      10,
		TotalPrice: 2000,
	}, models.OrderProduct{
		ProductID:  2,
		Quantity:   300,
		Price:      10,
		TotalPrice: 3000,
	})
	sup := models.Order{
		ID:                       id + 1,
		OrderTime:                time.Time{},
		CustomerID:               1,
		OrderType:                0,
		OrderProducts:            orderProducts,
		AccountsReceivable:       5000,
		ActualAccountsReceivable: 5000,
		Refund:                   0,
		ActualRefund:             0,
		Freight:                  0,
		Comment:                  "",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 开启事务
	if session, err := store.StartTransaction(); err == nil {
		// 执行事务
		err = mongo.WithSession(context.Background(), session, func(sessionContext mongo.SessionContext) error {

			// 商品库存修改
			productColl := store.ClientMongo.Database(DataBase).Collection("product")
			for _, orderProduct := range orderProducts {
				// 1. 定义查询条件
				filter := bson.D{{"_id", orderProduct.ProductID}}

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
		c.JSON(http.StatusOK, result.Success("success"))
	}

}

// CancelOrder 取消订单
// *
func (oh *Handler) CancelOrder(c *gin.Context) {
	result := &common.Result{}
	DataBase := common.GetTenantDateBase(c)
	id := 1

	order, err := common.FindOne[models.Order](DataBase, "order", id)
	if err != nil {
		c.JSON(http.StatusOK, result.Success("未发现订单信息"))
		return
	}
	if order.IsCancel == true {
		c.JSON(http.StatusOK, result.Success("该订单已被取消"))
		return
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
				filter := bson.D{{"_id", orderProduct.ProductID}}

				// 2. 获取该商品
				rst := productColl.FindOne(ctx, filter, options.FindOne())
				var product models.Product
				if err := rst.Decode(&product); err != nil {
					logs.LG.Error(err.Error())
					return err
				}

				// 3. 计算库存
				product.StockQuantity += orderProduct.Quantity

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

			// 设置订单取消
			order.IsCancel = true

			// 1. 定义查询条件
			filter := bson.D{{"_id", id}}

			// 2. 定义更新操作
			update := bson.D{{"$set", order}}

			collection.FindOneAndUpdate(ctx, filter, update)
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
		c.JSON(http.StatusOK, result.Success("success"))
	}
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
