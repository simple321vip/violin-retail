package models

// Goods 货品
type Goods struct {
	ID            int    `bson:"ID"`            // ID
	Name          string `bson:"Name"`          // 品名
	BigGoodType   int    `bson:"BigGoodType"`   // 大分类
	SmallGoodType int    `bson:"SmallGoodType"` // 小分类
	Brand         int    `bson:"Brand"`         // 品牌
	Unit          string `bson:"Unit"`          // 单位
	Price         int    `bson:"Price"`         // 建议零售价
	Quantity      int    `bson:"Quantity"`      // 出入库数量
	// 下面属性为可选
	Comment string `bson:"Comment"` // 备注
	Length  int    `bson:"Length"`  // 长度
	Width   int    `bson:"Width"`   // 宽度
}

// GoodType 分类
type GoodType struct {
	ID       int        `bson:"ID"`                       // ID
	Name     string     `bson:"Name"`                     // 分类名
	Rank     int        `bson:"Rank"`                     // 分类级别 1,2,3
	Comment  string     `bson:"Comment"`                  // 备注
	Children []GoodType `json:"children" bson:"children"` // 子分类
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
	return &GoodType{
		Children: make([]GoodType, 0),
	}
}

// SetID 实现SetID方法，用于接收并设置值
func (ms *GoodType) SetID(ID int) {
	ms.ID = ID
}

// NewBrand 品牌
func NewBrand() *Brand {
	return &Brand{}
}

// SetID 实现SetID方法，用于接收并设置值
func (ms *Brand) SetID(ID int) {
	ms.ID = ID
}
