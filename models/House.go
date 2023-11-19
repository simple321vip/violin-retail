package models

type House struct {
	ID   int    `bson:"_id"`  // 颜色ID
	Name string `bson:"Name"` // 颜色
}
