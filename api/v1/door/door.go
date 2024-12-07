package door

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

// GetDoorList 获取订单一览
// *
func (dh *Handler) GetDoorList(c *gin.Context) {
	result := &common.Result{}
	gh := dh.GetHandler()
	collection := store.ClientMongo.Database(gh.DatabaseName).Collection(gh.CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	query := bson.M{
		"$or": []bson.M{
			{
				"Phone": bson.M{
					"$regex":   c.Query("Phone"),
					"$options": "i",
				},
			},
			{
				"Name": bson.M{
					"$regex":   c.Query("Phone"),
					"$options": "i",
				},
			},
		},
	}
	c.Query("Phone")
	//filter := bson.D{{"Phone", {"$regex": "/" +  + "/"}}}
	find, err := collection.Find(ctx, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail(500, "系统内部错误"))
		return
	}
	var doorSheets []models.DoorSheet
	for find.Next(ctx) {
		var doorSheet models.DoorSheet
		err := find.Decode(&doorSheet)
		if err != nil {
			logs.LG.Error(err.Error())
			c.JSON(http.StatusInternalServerError, result.Fail(500, "系统内部错误"))
			return
		}
		doorSheets = append(doorSheets, doorSheet)
	}
	c.JSON(http.StatusOK, doorSheets)
}

// GetDoorSheet 获取订单一览
// *
func (dh *Handler) GetDoorSheet(c *gin.Context) {
	result := &common.Result{}
	gh := dh.GetHandler()
	ID, _ := strconv.Atoi(c.Param("ID"))
	one, err := gh.FindOne(ID)
	if err != nil {
		logs.LG.Error(err.Error())
		c.JSON(http.StatusInternalServerError, result.Fail(500, "系统内部错误"))
		return
	}
	var doorSheet = models.DoorSheet{}
	err = one.Decode(&doorSheet)
	if err != nil {
		logs.LG.Error(err.Error())
		c.JSON(http.StatusInternalServerError, result.Fail(500, "系统内部错误"))
	}

	c.JSON(http.StatusOK, doorSheet)

}

// CreateDoorSheet 创建柜门一览
// *
func (dh *Handler) CreateDoorSheet(c *gin.Context) {
	result := &common.Result{}
	var doorSheet = models.NewDoorSheet()
	err := c.ShouldBindJSON(doorSheet)
	gh := dh.GetHandler()
	// 正常for循环
	for i := 0; i < doorSheet.Amount; i++ {
		doorSheet.Doors = append(doorSheet.Doors, models.Door{
			ID:      i,
			Name:    "柜门" + strconv.Itoa(i+1),
			Length:  0,
			Width:   0,
			Area:    0,
			Amount:  0,
			Comment: "",
		})
	}
	t, err := gh.InsertOne(doorSheet)

	if err != nil {
		logs.LG.Error(err.Error())
		c.JSON(http.StatusInternalServerError, result.Fail(500, "系统内部错误"))
		return
	}

	c.JSON(http.StatusOK, t)
}

// DeleteDoorSheet
// *
func (dh *Handler) DeleteDoorSheet(c *gin.Context) {
	result := &common.Result{}
	ID, _ := strconv.Atoi(c.Param("ID"))
	gh := dh.GetHandler()
	err := gh.DeleteOne(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, result.Fail(500, err.Error()))
		return
	}
	c.JSON(http.StatusOK, nil)
}

// UpdateDoorSheet
// *
func (dh *Handler) UpdateDoorSheet(c *gin.Context) {
	result := &common.Result{}
	var doorSheet = models.NewDoorSheet()
	err := c.ShouldBindJSON(doorSheet)
	gh := dh.GetHandler()

	preAmount := len(doorSheet.Doors)
	// 正常for循环
	for i := preAmount; i < doorSheet.Amount; i++ {
		doorSheet.Doors = append(doorSheet.Doors, models.Door{
			ID:      i,
			Name:    "柜门" + strconv.Itoa(i+1),
			Length:  0,
			Width:   0,
			Area:    0,
			Amount:  0,
			Comment: "",
		})
	}

	err = gh.UpdateOne(doorSheet)
	if err != nil {
		logs.LG.Error(err.Error())
		c.JSON(http.StatusInternalServerError, result.Fail(500, "系统内部错误"))
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (dh *Handler) GetHandler() *common.BaseHandler {
	gh := &common.BaseHandler{
		DatabaseName:   "test",
		CollectionName: "doorsheet",
		//Collection:     common.,
	}
	return gh
}
