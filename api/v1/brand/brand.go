package brand

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

// GetBrand 获取所有品牌
// **
func (th *Handler) GetBrand(c *gin.Context) {
	result := &common.Result{}
	brands, err := th.getAllBrands()
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail(500, "系统内部错误"))
		return
	}
	c.JSON(http.StatusOK, brands)
}

// CreateBrand 新增品牌
// **
func (th *Handler) CreateBrand(c *gin.Context) {
	result := &common.Result{}
	var brand = models.NewBrand()
	err := c.ShouldBindJSON(brand)
	gh := th.GetHandler()

	_, err = gh.InsertOne(brand)
	if err != nil {
		logs.LG.Error(err.Error())
		c.JSON(http.StatusInternalServerError, result.Fail(500, "系统内部错误"))
		return
	}

	brands, err := th.getAllBrands()
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail(500, "系统内部错误"))
		return
	}
	c.JSON(http.StatusOK, brands)
}

// DeleteBrand 删除品牌
// *
func (th *Handler) DeleteBrand(c *gin.Context) {
	result := &common.Result{}
	ID, _ := strconv.Atoi(c.Param("ID"))
	gh := th.GetHandler()
	err := gh.DeleteOne(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail(500, err.Error()))
		return
	}
	brands, err := th.getAllBrands()
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail(500, "系统内部错误"))
		return
	}
	c.JSON(http.StatusOK, brands)
}

// UpdateBrand 品牌修改
// *
func (th *Handler) UpdateBrand(c *gin.Context) {
	result := &common.Result{}
	gh := th.GetHandler()
	var brand = models.NewBrand()
	err := c.ShouldBindJSON(brand)
	err = gh.UpdateOne(brand)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail(500, "系统内部错误"))
		return
	}
	brands, err := th.getAllBrands()
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail(500, "系统内部错误"))
		return
	}
	c.JSON(http.StatusOK, brands)
}

func (th *Handler) getAllBrands() ([]models.Brand, error) {
	gh := th.GetHandler()
	collection := store.ClientMongo.Database(gh.DatabaseName).Collection(gh.CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	find, err := collection.Find(ctx, bson.D{})
	if err != nil {
		logs.LG.Error(err.Error())
		return nil, err
	}
	var brands []models.Brand
	for find.Next(ctx) {
		var brand models.Brand
		err := find.Decode(&brand)
		if err != nil {
			logs.LG.Error(err.Error())
			return nil, err
		}
		brands = append(brands, brand)
	}
	return brands, nil
}

func (th *Handler) GetHandler() *common.BaseHandler {
	gh := &common.BaseHandler{
		DatabaseName:   "test",
		CollectionName: "brand",
		Collection:     common.Brand,
	}
	return gh
}
