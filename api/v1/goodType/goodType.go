package goodType

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

// GetGoodType 获取分类一栏
// **
func (th *Handler) GetGoodType(c *gin.Context) {
	result := &common.Result{}
	gh := th.GetHandler()
	collection := store.ClientMongo.Database(gh.DatabaseName).Collection(gh.CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	find, err := collection.Find(ctx, bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail(500, "系统内部错误"))
		return
	}
	var goodTypes []models.GoodType
	for find.Next(ctx) {
		var goodType models.GoodType
		err := find.Decode(&goodType)
		if err != nil {
			logs.LG.Error(err.Error())
			return
		}
		goodTypes = append(goodTypes, goodType)
	}
	c.JSON(http.StatusOK, goodTypes)
}

// CreateGoodType 新增商品分类
// **
func (th *Handler) CreateGoodType(c *gin.Context) {
	result := &common.Result{}
	var goodType = models.NewGoodType()
	err := c.ShouldBindJSON(goodType)
	gh := th.GetHandler()

	_, err = gh.InsertOne(goodType)
	if err != nil {
		logs.LG.Error(err.Error())
		c.JSON(http.StatusInternalServerError, result.Fail(500, "系统内部错误"))
		return
	}
}

// DeleteGoodType 删除货品
// *
func (th *Handler) DeleteGoodType(c *gin.Context) {
	result := &common.Result{}
	ID, _ := strconv.Atoi(c.Param("ID"))
	gh := th.GetHandler()
	err := gh.DeleteOne(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail(500, err.Error()))
		return
	}
	c.JSON(http.StatusOK, nil)
}

// UpdateGoodType 分类信息修改
// *
func (th *Handler) UpdateGoodType(c *gin.Context) {
	result := &common.Result{}
	gh := th.GetHandler()
	var goodType = models.NewGoodType()
	err := c.ShouldBindJSON(goodType)
	err = gh.UpdateOne(goodType)
	if err != nil {
		logs.LG.Error(err.Error())
		c.JSON(http.StatusInternalServerError, result.Fail(500, "系统内部错误"))
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (th *Handler) GetHandler() *common.BaseHandler {
	gh := &common.BaseHandler{
		DatabaseName:   "test",
		CollectionName: "goodType",
		Collection:     common.GoodType,
	}
	return gh
}
