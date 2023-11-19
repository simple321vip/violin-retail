package models

type Size struct {
	ID            int    `bson:"_id"`           // 尺寸ID
	Specification string `bson:"Specification"` // 尺寸规格
}
