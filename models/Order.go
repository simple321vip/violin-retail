package models

import "time"

// Order 订单
type Order struct {
	ID                       int          `bson:"_id"`                                // ID
	OrderTime                time.Time    `bson:"OrderTime"`                          // 订单创建时间
	OrderGoods               []OrderGoods `json:"OrderProducts" bson:"OrderProducts"` // 商品一览
	AccountsReceivable       float64      `bson:"AccountsReceivable"`                 // 应收账款
	ActualAccountsReceivable float64      `bson:"ActualAccountsReceivable"`           // 实收账款
	Refund                   float64      `bson:"Refund"`                             // 退款
	ActualRefund             float64      `bson:"ActualRefund"`                       // 实际退款
	Freight                  float64      `bson:"Freight"`                            // 运费
	Comment                  string       `json:"Comment" bson:"Name"`                // 备注
	IsCancel                 bool         `bson:"IsCancel"`                           // 是否取消
	Status                   int          `bson:"Status"`                             // 订单状态
	CustomerPhone            int          `bson:"CustomerPhone"`                      // 顾客电话
	CustomerName             string       `bson:"CustomerName"`                       // 顾客姓名
	CustomerAddress          string       `bson:"CustomerAddress"`                    // 顾客地址
	Version                  int          `bson:"Version"`                            // 版本号

}

// OrderGoods 订单商品
type OrderGoods struct {
	ID            int     `bson:"ProductID"`        // 商品ID
	Name          string  `bson:"Name"`             // 商品名
	BigGoodType   string  `bson:"BigGoodType"`      // 大分类
	SmallGoodType string  `bson:"SmallGoodType"`    // 小分类
	Color         string  `bson:"Color,omitempty"`  // 颜色
	Size          string  `bson:"Size,omitempty"`   // 尺寸
	Brand         string  `bson:"Brand,omitempty"`  // 品牌
	Unit          string  `bson:"Unit,omitempty"`   // 单位
	Amount        string  `bson:"Amount,omitempty"` // 数量
	Price         float64 `bson:"Price"`            // 单价
	TotalPrice    float64 `bson:"TotalPrice"`       // 总价
}

// NewOrder 订单
func NewOrder() *Order {
	return &Order{}
}
