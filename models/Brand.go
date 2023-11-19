package models

type Brand struct {
	ID   int    `bson:"_id"`
	Age  int    `bson:"age"`
	Name string `bson:"name"`
	Addr string `bson:"address"`
}
