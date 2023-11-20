package order

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
	//result := &common.Result{}
	//DataBase := common.GetTenantDateBase(c)

}

// CancelOrder 取消订单
// *
func (oh *Handler) CancelOrder(c *gin.Context) {
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
