package customer

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

// GetCustomers 获取客户一览
// **
func (ch *Handler) GetCustomers(c *gin.Context) {
	result := &common.Result{}
	gh := ch.GetHandler()
	collection := store.ClientMongo.Database(gh.DatabaseName).Collection(gh.CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	find, err := collection.Find(ctx, bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail(500, "系统内部错误"))
		return
	}
	var customers []models.Customer
	for find.Next(ctx) {
		var customer models.Customer
		err := find.Decode(&customer)
		if err != nil {
			logs.LG.Error(err.Error())
			return
		}
		customers = append(customers, customer)
	}
	c.JSON(http.StatusOK, customers)
}

// CreateCustomer 创建客户
// *
func (ch *Handler) CreateCustomer(c *gin.Context) {
	result := &common.Result{}
	var customer = models.NewCustomer()
	err := c.ShouldBindJSON(customer)
	gh := ch.GetHandler()

	collection := store.ClientMongo.Database(gh.DatabaseName).Collection(gh.CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	query := bson.M{
		"Phone": customer.Phone,
	}
	find, err := collection.Find(ctx, query)
	if find.TryNext(ctx) {
		c.JSON(http.StatusBadRequest, result.Fail(500, "已存在此客户，无需创建。"))
		return
	}
	_, err = gh.InsertOne(customer)

	if err != nil {
		logs.LG.Error(err.Error())
		c.JSON(http.StatusInternalServerError, result.Fail(500, "系统内部错误"))
		return
	}
}

// DeleteCustomer 删除客户
// *
func (ch *Handler) DeleteCustomer(c *gin.Context) {
	result := &common.Result{}
	ID, _ := strconv.Atoi(c.Param("ID"))
	gh := ch.GetHandler()
	err := gh.DeleteOne(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail(500, err.Error()))
		return
	}
	c.JSON(http.StatusOK, nil)
}

// UpdateCustomer 客户信息修改
// *
func (ch *Handler) UpdateCustomer(c *gin.Context) {
	result := &common.Result{}
	gh := ch.GetHandler()
	var customer = models.NewCustomer()
	err := c.ShouldBindJSON(customer)

	collection := store.ClientMongo.Database(gh.DatabaseName).Collection(gh.CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	query := bson.M{
		"Phone": customer.Phone,
		"ID":    bson.M{"$ne": customer.ID},
	}
	find, err := collection.Find(ctx, query)
	if find.TryNext(ctx) {
		c.JSON(http.StatusBadRequest, result.Fail(500, "此电话已经存在，请确认。"))
		return
	}

	err = gh.UpdateOne(customer)
	if err != nil {
		logs.LG.Error(err.Error())
		c.JSON(http.StatusInternalServerError, result.Fail(500, "系统内部错误"))
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (ch *Handler) GetHandler() *common.BaseHandler {
	gh := &common.BaseHandler{
		DatabaseName:   "test",
		CollectionName: "customer",
		Collection:     common.Customer,
	}
	return gh
}
