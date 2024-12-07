package house

import (
	"github.com/gin-gonic/gin"
)

// GetSuppliers 供货商一览
// **
func (hh *Handler) GetSuppliers(c *gin.Context) {
	//result := &common.Result{}
	//
	//DataBase := common.GetTenantDateBase(c)
	//
	//suppliers, err := common.Find[models.Supplier](DataBase, "supplier", bson.D{})
	//if err != nil {
	//	logs.LG.Error(err.Error())
	//	return
	//}
	//c.JSON(http.StatusOK, result.Success(suppliers))
}

// CreateSupplier 创建供货商
// *
func (hh *Handler) CreateSupplier(c *gin.Context) {
	//result := &common.Result{}
	//DataBase := common.GetTenantDateBase(c)
	//
	//ID, _ := common.GetNextID(DataBase, "supplier")
	//supplier := models.Supplier{
	//	ID:              ID + 1,
	//	Name:            "兔宝宝有限公司",
	//	Contacts:        "徐明华",
	//	Phone:           "1333333333",
	//	Address:         "广东潮汕",
	//	AccountsPayable: 0,
	//	Comment:         "兔宝宝直属进货商",
	//}
	//
	//err := common.InsertOne(DataBase, "supplier", supplier)
	//if err != nil {
	//	logs.LG.Error(err.Error())
	//	c.JSON(http.StatusInternalServerError, result.Fail(500, "系统内部错误"))
	//	return
	//}
	//c.JSON(http.StatusOK, result.Success("success"))
}

// DeleteSupplier 删除供货商
// *
func (hh *Handler) DeleteSupplier(c *gin.Context) {
	//result := &common.Result{}
	//DataBase := common.GetTenantDateBase(c)
	//
	//ID := 1
	//err := common.DeleteOne(DataBase, "supplier", ID)
	//if err != nil {
	//	logs.LG.Error(err.Error())
	//	c.JSON(http.StatusInternalServerError, result.Fail(500, "系统内部错误"))
	//	return
	//}
	//c.JSON(http.StatusOK, result.Success("success"))
}

// UpdateSupplier 供货商信息修改
// *
func (hh *Handler) UpdateSupplier(c *gin.Context) {
	//result := &common.Result{}
	//DataBase := common.GetTenantDateBase(c)
	//ID := 1
	//
	//supplier := models.Supplier{
	//	ID:              ID,
	//	Name:            "兔宝宝有限公司",
	//	Contacts:        "徐明华",
	//	Phone:           "1333333333",
	//	Address:         "广东潮汕臭水沟",
	//	AccountsPayable: 0,
	//	Comment:         "兔宝宝直属进货商",
	//}
	//
	//err := common.UpdateOne(DataBase, "supplier", ID, supplier)
	//if err != nil {
	//	logs.LG.Error(err.Error())
	//	c.JSON(http.StatusInternalServerError, result.Fail(500, "系统内部错误"))
	//	return
	//}
	//c.JSON(http.StatusOK, result.Success("success"))
}
