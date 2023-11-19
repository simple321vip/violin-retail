package models

type Product struct {
	ID            int                                                   `bson:"_id"`
	Name          string                                                `bson:"Name"`                                             // 商品名称
	Unit          string                                                `bson:"Unit"`                                             // 商品单位
	StockQuantity int                                                   `bson:"StockQuantity"`                                    // 库存数量
	RetailPrice   float64                                               `bson:"RetailPrice"`                                      // 零售价格
	ImageUrl      string                                                `bson:"ImageUrl"`                                         // 商品图片地址
	CategoryID    int                                                   `json:"CategoryID,omitempty" bson:"CategoryID,omitempty"` // 商品种类ID
	ColorID       int                                                   `json:"ColorID,omitempty" bson:"ColorID,omitempty"`       // 商品颜色ID
	*Color        `json:"Color,omitempty" bson:"Color,omitempty"`       // 商品颜色ID
	SizeID        int                                                   `json:"SizeID,omitempty" bson:"SizeID,omitempty"` // 商品规格ID
	*Category     `json:"Category,omitempty" bson:"Category,omitempty"` // 商品种类ID
	*Size         `json:"Size,omitempty" bson:"Size,omitempty"`         // 商品种类ID
}
