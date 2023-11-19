package models

type Category struct {
	ID   int    `bson:"_id"`  // 种类ID
	Name string `bson:"Name"` // 种类名称
}
