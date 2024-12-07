package models

import "time"

// Door 柜门
type Door struct {
	ID      int     `bson:"ID"`                     // ID
	Name    string  `bson:"Name"`                   // 柜门名称
	Length  int     `bson:"Length"`                 // 长度
	Width   int     `bson:"Width"`                  // 宽度
	Area    float64 `bson:"Area"`                   // 面积
	Amount  int     `bson:"Amount"`                 // 数量
	Comment string  `json:"Comment" bson:"Comment"` // 备注
	//*DoorSheet `json:"DoorSheet,omitempty" bson:"DoorSheet,omitempty"` // 商品种类ID
}

// DoorSheet 柜门单
type DoorSheet struct {
	ID        int       `bson:"ID"`                 // 柜门表单ID
	Phone     string    `bson:"Phone"`              // 柜门表单Phone
	Name      string    `bson:"Name"`               // 柜门表单Name
	Comment   string    `bson:"Comment"`            // 柜门表单Comment
	Amount    int       `bson:"Amount"`             // 柜门表单Amount
	OrderTime time.Time `bson:"OrderTime"`          // 订单创建时间
	Doors     []Door    `json:"Doors" bson:"Doors"` // 柜门一览
	TotalArea float64   `bson:"TotalArea"`          // 退款
}

// NewDoorSheet 默认
func NewDoorSheet() *DoorSheet {
	Doors := make([]Door, 0)
	return &DoorSheet{
		TotalArea: 0,
		Doors:     Doors,
	}
}
