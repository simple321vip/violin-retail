package common

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"violin-home.cn/retail/store"
)

type Max struct {
	ID int `bson:"_id"` // 种类ID
}

type T interface {
}

func GetNextID(databaseName string, collectionName string) (int, error) {
	collection := store.ClientMongo.Database(databaseName).Collection(collectionName)
	opts := options.FindOne()
	opts.SetSort(bson.D{{"_id", -1}})

	rst := collection.FindOne(context.TODO(), bson.D{}, opts)
	var max Max
	if err := rst.Decode(&max); err != nil {
		return 0, nil
	}
	return max.ID, nil
}

func GetTenantDateBase(c *gin.Context) string {
	TenantID := c.GetHeader("TenantID")
	if TenantID == "" {
		return "test"
	}
	return TenantID
}

// InsertOne 插入指定对象，包含事务，连续多个插入时，请勿使用。
// *
func InsertOne[T any](databaseName string, collectionName string, model T) error {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 开启事务
	session, err := store.StartTransaction()
	if err != nil {
		return err
	}

	// 执行事务
	err = mongo.WithSession(context.Background(), session, func(sessionContext mongo.SessionContext) error {

		collection := store.ClientMongo.Database(databaseName).Collection(collectionName)

		bsonData, err := bson.Marshal(model)

		if err != nil {
			return err
		}
		_, err = collection.InsertOne(ctx, bsonData)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	// 提交事务
	err = store.CommitTransaction(session)
	if err != nil {
		return err
	}
	return nil
}

// DeleteOne 指定_id删除，已包含事务，连续多个删除时，请勿使用。
// *
func DeleteOne(databaseName string, collectionName string, ID int) error {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 开启事务
	session, err := store.StartTransaction()
	if err != nil {
		return err
	}

	// 执行事务
	err = mongo.WithSession(context.Background(), session, func(sessionContext mongo.SessionContext) error {
		collection := store.ClientMongo.Database(databaseName).Collection(collectionName)
		filter := bson.M{
			"_id": ID,
		}
		_, err := collection.DeleteOne(ctx, filter)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	// 提交事务
	err = store.CommitTransaction(session)
	if err != nil {
		return err
	}
	return nil
}

// UpdateOne 指定_id删除，已包含事务，连续多个删除时，请勿使用。
// *
func UpdateOne[T any](databaseName string, collectionName string, ID int, model T) error {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 开启事务
	session, err := store.StartTransaction()
	if err != nil {
		return err
	}

	// 执行事务
	err = mongo.WithSession(context.Background(), session, func(sessionContext mongo.SessionContext) error {
		collection := store.ClientMongo.Database(databaseName).Collection(collectionName)
		bsonData, err := bson.Marshal(model)
		var doc *bson.D
		err = bson.Unmarshal(bsonData, &doc)
		if err != nil {
			return err
		}

		// 1. 定义查询条件
		filter := bson.D{{"_id", ID}}

		// 2. 定义更新操作
		update := bson.D{{"$set", doc}}

		collection.FindOneAndUpdate(ctx, filter, update)

		return nil
	})

	if err != nil {
		return err
	}

	// 提交事务
	err = store.CommitTransaction(session)
	if err != nil {
		return err
	}
	return nil
}
