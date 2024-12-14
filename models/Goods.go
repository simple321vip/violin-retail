package models

// Goods 货品
type Goods struct {
	ID        int                                                   `bson:"ID"`   // ID
	Name      string                                                `bson:"Name"` // 品名
	*GoodType `json:"GoodType,omitempty" bson:"GoodType,omitempty"` // 分类
	*Brand    `json:"Brand,omitempty" bson:"Brand,omitempty"`       // 品牌
	Unit      string                                                `bson:"Unit"` // 单位
	// 下面属性为可选
	Comment string `json:"Comment" bson:"Comment"` // 备注
	Length  int    `bson:"Length"`                 // 长度
	Width   int    `bson:"Width"`                  // 宽度
}

// GoodType 分类
type GoodType struct {
	ID      int    `bson:"ID"`      // ID
	Name    string `bson:"Name"`    // 分类名
	Comment string `bson:"Comment"` // 备注
}

// Brand 品牌
type Brand struct {
	ID      int    `bson:"ID"`      // ID
	Name    string `bson:"Name"`    // 品牌名
	Comment string `bson:"Comment"` // 备注
}

// NewGoods 货品
func NewGoods() *Goods {
	return &Goods{}
}

// SetID 实现SetID方法，用于接收并设置值
func (ms *Goods) SetID(ID int) {
	ms.ID = ID
}

// NewGoodType 分类
func NewGoodType() *GoodType {
	return &GoodType{}
}

// SetID 实现SetID方法，用于接收并设置值
func (ms *GoodType) SetID(ID int) {
	ms.ID = ID
}
