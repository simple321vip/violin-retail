package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"violin-home.cn/retail/common/logs"
)

type Body struct {
	ID int `json:"ID"`
}

// CheckAndGetID CHECK ID
func CheckAndGetID(c *gin.Context) error {
	ID, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		return err
	}
	var body Body
	err = c.ShouldBind(&body)
	if err != nil {
		return fmt.Errorf("field %s cannot be set", "ID")
	}

	if ID != body.ID {
		logs.LG.Error(err.Error())
		// c.JSON(http.StatusInternalServerError, result.Fail(500, "请求路径ID和Body中ID不匹配"))
		return fmt.Errorf("field %s cannot be set", "ID")
	}
	return nil
}
