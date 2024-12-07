package models

type Customer struct {
	ID                 int         `bson:"ID"`
	Name               string      `bson:"Name"`                         // 客户名称
	Rank               int         `bson:"Rank"`                         // 客户等级
	Contacts           string      `bson:"Contacts"`                     // 联系人
	Phone              string      `bson:"Phone"`                        // 电话
	Address            string      `bson:"Address"`                      // 送货地址
	AccountsReceivable float64     `bson:"AccountsReceivable"`           // 应收账款
	Comment            string      `bson:"Comment"`                      // 备注
	DoorSheets         []DoorSheet `json:"DoorSheets" bson:"DoorSheets"` // 柜门单
}

// NewCustomer 默认
func NewCustomer() *Customer {
	return &Customer{}
}

// SetID 实现SetID方法，用于接收并设置值
func (ms *Customer) SetID(ID int) {
	ms.ID = ID
}
