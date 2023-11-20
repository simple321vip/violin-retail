package models

import "time"

// Order 订单
type Order struct {
	ID                       int       `bson:"_id"`                      // ID
	OrderTime                time.Time `bson:"OrderTime"`                // 出入库时间
	CustomerID               int       `bson:"CustomerID"`               // 客户ID
	OrderType                int       `bson:"OrderType"`                // 出库或入库 0为销售，1为退货
	Product                  []Product `bson:"Product"`                  // 商品一览
	AccountsReceivable       float64   `bson:"AccountsReceivable"`       // 应收账款
	ActualAccountsReceivable float64   `bson:"ActualAccountsReceivable"` // 实收账款
	Freight                  float64   `bson:"Freight"`                  // 运费
	Comment                  string    `bson:"Name"`                     // 备注
}
