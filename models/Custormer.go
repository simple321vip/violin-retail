package models

type Customer struct {
	ID                 int     `bson:"_id"`
	Name               string  `bson:"Name"`               // 客户名称
	Rank               int     `bson:"Rank"`               // 客户等级
	Contacts           string  `bson:"Contacts"`           // 联系人
	Phone              string  `bson:"Phone"`              // 电话
	Address            string  `bson:"Address"`            // 送货地址
	AccountsReceivable float64 `bson:"AccountsReceivable"` // 应收账款
	Comment            string  `bson:"Comment"`            // 备注
}
