package models

import "time"

// Order 订单
type Order struct {
	ID int `bson:"_id"` // ID

	OrderTime                time.Time                                             `bson:"OrderTime"`                          // 订单创建时间
	OrderType                int                                                   `bson:"OrderType"`                          // 0为销售，1为退货
	OrderProducts            []OrderProduct                                        `json:"OrderProducts" bson:"OrderProducts"` // 商品一览
	AccountsReceivable       float64                                               `bson:"AccountsReceivable"`                 // 应收账款
	ActualAccountsReceivable float64                                               `bson:"ActualAccountsReceivable"`           // 实收账款
	Refund                   float64                                               `bson:"Refund"`                             // 退款
	ActualRefund             float64                                               `bson:"ActualRefund"`                       // 实际退款
	Freight                  float64                                               `bson:"Freight"`                            // 运费
	Comment                  string                                                `json:"Comment" bson:"Name"`                // 备注
	IsCancel                 bool                                                  `bson:"IsCancel"`                           // 是否取消
	CustomerId               int                                                   `bson:"CustomerId"`                         // 商品种类ID
	*Customer                `json:"Customer,omitempty" bson:"Customer,omitempty"` // 商品种类ID

}

// OrderProduct 出入库商品
type OrderProduct struct {
	ID         int     `bson:"ProductID"`  // 商品ID
	Quantity   int     `bson:"Quantity"`   // 出入库数量
	Price      float64 `bson:"Price"`      // 单价
	TotalPrice float64 `bson:"TotalPrice"` // 总价
	Checked    bool    `bson:"Checked"`    // 是否选中
}
