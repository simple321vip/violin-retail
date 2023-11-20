package models

import "time"

// House 进货 退货
type House struct {
	ID              int            `bson:"_id"`             // ID
	HouseTime       time.Time      `bson:"HouseTime"`       // 出入库时间
	SupplierID      int            `bson:"SupplierID"`      // 供货商ID
	HouseType       int            `bson:"HouseType"`       // 出库或入库 0为出库，1为入库
	HouseProduct    []HouseProduct `bson:"HouseProduct"`    // 数量
	AccountsPayable float64        `bson:"AccountsPayable"` // 应付账款
	ActualPaid      float64        `bson:"ActualPaid"`      // 实付付账款
	Freight         float64        `bson:"Freight"`         // 运费
	Comment         string         `bson:"Name"`            // 备注
}

// Supplier 供货商
type Supplier struct {
	ID              int     `bson:"_id"`             // ID
	Name            string  `bson:"Name"`            // 供货商
	Contacts        string  `bson:"Contacts"`        // 联系人
	Phone           string  `bson:"Phone"`           // 电话
	Address         string  `bson:"Address"`         // 送货地址
	AccountsPayable float64 `bson:"AccountsPayable"` // 初始应付账款
	Comment         string  `bson:"Comment"`         // 备注
}

// HouseProduct 出入库商品
type HouseProduct struct {
	ProductID  int     `bson:"ProductID"`  // 商品ID
	Quantity   int     `bson:"Quantity"`   // 出入库数量
	HousePrice float64 `bson:"HousePrice"` // 价格
}
