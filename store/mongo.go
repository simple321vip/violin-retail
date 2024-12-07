package store

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"violin-home.cn/retail/common/logs"
)

var ClientMongo *mongo.Client

func NewMongoClient() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	ClientMongo, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo:27017"))
}

// StartTransaction 开启事务
func StartTransaction() (mongo.Session, error) {

	// 创建会话
	session, err := ClientMongo.StartSession()
	if err != nil {
		return session, err
	}
	logs.LG.Debug("创建会话")

	//开始事务
	err = session.StartTransaction()

	if err != nil {
		return session, err
	}
	logs.LG.Debug("开启事务")
	return session, nil
}

func CommitTransaction(session mongo.Session) error {

	// 提交事务
	err := session.CommitTransaction(context.Background())

	if err != nil {
		return err
	}
	logs.LG.Debug("提交事务")
	// 关闭会话
	session.EndSession(context.Background())

	if err != nil {
		return err
	}
	logs.LG.Debug("关闭会话")
	return nil
}
