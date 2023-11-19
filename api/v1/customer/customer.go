package customer

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

// GetCustomers 获取客户一览
// **
func (ch *Handler) GetCustomers(c *gin.Context) {
	result := &common.Result{}

	DataBase := common.GetTenantDateBase(c)
	collection := store.ClientMongo.Database(DataBase).Collection("customer")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if find, err := collection.Find(ctx, bson.D{}); err == nil {
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
		c.JSON(http.StatusOK, result.Success(customers))
	}

}

// CreateCustomer 创建客户
// *
func (ch *Handler) CreateCustomer(c *gin.Context) {
	result := &common.Result{}
	DataBase := common.GetTenantDateBase(c)

	ID, _ := common.GetNextID(DataBase, "customer")
	Customer := models.Customer{
		ID:                 ID + 1,
		Name:               "李刚",
		Rank:               0,
		Contacts:           "李刚",
		Phone:              "13332247026",
		Address:            "北京天安门",
		AccountsReceivable: 0,
		Comment:            "我爱北京天安门",
	}

	err := common.InsertOne(DataBase, "customer", Customer)
	if err != nil {
		logs.LG.Error(err.Error())
		c.JSON(http.StatusInternalServerError, result.Fail(500, "系统内部错误"))
		return
	}
	c.JSON(http.StatusOK, result.Success("success"))
}

// DeleteCustomer 删除客户
// *
func (ch *Handler) DeleteCustomer(c *gin.Context) {
	result := &common.Result{}
	DataBase := common.GetTenantDateBase(c)

	ID := 1
	err := common.DeleteOne(DataBase, "customer", ID)
	if err != nil {
		logs.LG.Error(err.Error())
		c.JSON(http.StatusInternalServerError, result.Fail(500, "系统内部错误"))
		return
	}
	c.JSON(http.StatusOK, result.Success("success"))
}

// UpdateCustomer 客户信息修改
// *
func (ch *Handler) UpdateCustomer(c *gin.Context) {
	result := &common.Result{}
	DataBase := common.GetTenantDateBase(c)
	ID := 1

	Customer := models.Customer{
		ID:                 ID,
		Name:               "xiaowang",
		Rank:               0,
		Contacts:           "李刚1",
		Phone:              "133",
		Address:            "2222",
		AccountsReceivable: 0,
		Comment:            "11111",
	}

	err := common.UpdateOne(DataBase, "customer", ID, Customer)
	if err != nil {
		logs.LG.Error(err.Error())
		c.JSON(http.StatusInternalServerError, result.Fail(500, "系统内部错误"))
		return
	}
	c.JSON(http.StatusOK, result.Success("success"))
}
