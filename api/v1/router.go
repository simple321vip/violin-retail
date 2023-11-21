package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"violin-home.cn/retail/api/v1/customer"
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
	v1 := r.Group("violin-retail/api/v1", Interceptor())
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
		v1.GET("/order/put", oh.CreateOrder)
		v1.GET("/order/cancel", oh.CancelOrder)

		// 客户
		v1.GET("/customer/put", ch.CreateCustomer)
		v1.GET("/customer/update", ch.UpdateCustomer)

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
