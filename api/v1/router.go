package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"violin-home.cn/retail/api/v1/customer"
	"violin-home.cn/retail/api/v1/door"
	"violin-home.cn/retail/api/v1/goods"
	"violin-home.cn/retail/api/v1/house"
	"violin-home.cn/retail/api/v1/order"
	"violin-home.cn/retail/common/logs"
	"violin-home.cn/retail/router"
	"violin-home.cn/retail/store"
)

func init() {
	router.Register(&Router{})
}

type Router struct {
}

func (sr *Router) Route(r *gin.Engine) {

	gh := &goods.Handler{}
	hh := &house.Handler{}
	ch := &customer.Handler{}
	oh := &order.Handler{}
	dh := &door.Handler{}
	v1 := r.Group("retail/api/v1", Interceptor())
	{
		// 商品
		v1.GET("/goods", gh.GetGoods)
		v1.GET("/goods/put", gh.IncreaseGoods)

		// 出入库
		v1.GET("/house", hh.GetHouseList)
		v1.GET("/house/:houseId", hh.GetHouse)
		v1.GET("/house/in", hh.HouseIn)
		//v1.GET("/house/out", hh.HouseOut)

		v1.GET("/supplier", hh.GetSuppliers)
		v1.GET("/supplier/put", hh.CreateSupplier)

		// 订单
		v1.POST("/order", oh.CreateOrder)
		v1.GET("/order", oh.GetOrder)
		v1.GET("/orders", oh.GetOrderList)
		v1.GET("/order/cancel", oh.CancelOrder)

		// 客户
		v1.POST("/customer", ch.CreateCustomer)
		v1.DELETE("/customer/:ID", ch.DeleteCustomer)
		v1.PUT("/customer/:ID", ch.UpdateCustomer)
		v1.GET("/customer", ch.GetCustomers)

		// 柜门
		v1.GET("/doorSheet", dh.GetDoorList)
		v1.GET("/doorSheet/:ID", dh.GetDoorSheet)
		v1.DELETE("/doorSheet/:ID", dh.DeleteDoorSheet)
		v1.PUT("/doorSheet/:ID", dh.UpdateDoorSheet)
		v1.POST("/doorSheet", dh.CreateDoorSheet)

		// 货物
		//v1.POST("/doorSheet/:id", dh.GetDoorList)
		//v1.DELETE("/customer/:id", dh.DeleteCustomer)
		//v1.PUT("/customer/:id", dh.UpdateCustomer)
		v1.POST("/goods", gh.IncreaseGoods)
		v1.POST("/goodType", gh.IncreaseGoodType)
		v1.DELETE("/goodType", gh.DeleteGoodType)
	}

}

// Interceptor
// 调用链
// c.Next() 执行下一个调用，直到后续调用完成后，才会执行c.Next() 下面代码
// c.Abort() 不在执行后续调用方法，但是 c.Abort() 后续代码会执行完毕
func Interceptor() gin.HandlerFunc {

	return func(c *gin.Context) {
		logs.LG.Info("Interceptor begin")
		//c.Next()
		return
		logs.LG.Info("Interceptor ...")

		tenantId := c.GetHeader("tenantid")

		if tenantId == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "tenantid is empty"})
			c.Abort()
			return
		}

		authorization := c.GetHeader("authorization")
		var tmp []string
		if tmp = strings.Split(authorization, ":"); len(tmp) < 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			c.Abort()
			return
		}
		token := tmp[1]

		result, err := store.ClientRedis.Get(tenantId).Result()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": err.Error()})
			c.Abort()
			return
		}

		if result != token {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "authorized error"})
			c.Abort()
			return
		}

		c.Next()
	}
}
