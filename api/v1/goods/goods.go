package goods

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

// GetGoods 获取商品一栏
// **
func (nh *Handler) GetGoods(c *gin.Context) {
	result := &common.Result{}
	goods, err := nh.getAllGoods()
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail(500, "系统内部错误"))
		return
	}
	c.JSON(http.StatusOK, goods)
}

// CreateGoods 新增商品
// **
func (nh *Handler) CreateGoods(c *gin.Context) {
	result := &common.Result{}
	var goods = models.NewGoods()
	err := c.ShouldBindJSON(goods)
	gh := nh.GetHandler()

	_, err = gh.InsertOne(goods)
	if err != nil {
		logs.LG.Error(err.Error())
		c.JSON(http.StatusInternalServerError, result.Fail(500, "系统内部错误"))
		return
	}
}

// DeleteGoods 删除货品
// *
func (nh *Handler) DeleteGoods(c *gin.Context) {
	result := &common.Result{}
	ID, _ := strconv.Atoi(c.Param("ID"))
	gh := nh.GetHandler()
	err := gh.DeleteOne(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail(500, err.Error()))
		return
	}
	c.JSON(http.StatusOK, nil)
}

// UpdateGoods 货品信息修改
// *
func (nh *Handler) UpdateGoods(c *gin.Context) {
	result := &common.Result{}
	gh := nh.GetHandler()
	var goods = models.NewGoods()
	err := c.ShouldBindJSON(goods)

	//collection := store.ClientMongo.Database(gh.DatabaseName).Collection(gh.CollectionName)
	//ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	//defer cancel()
	//query := bson.M{
	//	"Phone": customer.Phone,
	//	"ID":    bson.M{"$ne": customer.ID},
	//}
	//find, err := collection.Find(ctx, query)
	//if find.TryNext(ctx) {
	//	c.JSON(http.StatusBadRequest, result.Fail(500, "此电话已经存在，请确认。"))
	//	return
	//}

	err = gh.UpdateOne(goods)
	if err != nil {
		logs.LG.Error(err.Error())
		c.JSON(http.StatusInternalServerError, result.Fail(500, "系统内部错误"))
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (nh *Handler) getAllGoods() ([]models.Goods, error) {
	gh := nh.GetHandler()
	collection := store.ClientMongo.Database(gh.DatabaseName).Collection(gh.CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	find, err := collection.Find(ctx, bson.D{})
	if err != nil {
		logs.LG.Error(err.Error())
		return nil, err
	}
	var goods []models.Goods
	for find.Next(ctx) {
		var good models.Goods
		err := find.Decode(&good)
		if err != nil {
			logs.LG.Error(err.Error())
			return nil, err
		}
		goods = append(goods, good)
	}
	return goods, nil
}

func (nh *Handler) GetHandler() *common.BaseHandler {
	gh := &common.BaseHandler{
		DatabaseName:   "test",
		CollectionName: "goods",
		Collection:     common.Goods,
	}
	return gh
}
