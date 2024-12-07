package common

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
	"time"
	"violin-home.cn/retail/store"
)

type Max struct {
	ID int `bson:"ID"` // 种类ID
}

type BaseHandler struct {
	DatabaseName   string `bson:"DatabaseName"`   // DB NAME
	CollectionName string `bson:"CollectionName"` // Coll NAME
	Collection     interface{}
}

type T interface {
}

func (ch *BaseHandler) GetNextID() (int, error) {
	collection := store.ClientMongo.Database(ch.DatabaseName).Collection(ch.CollectionName)
	opts := options.FindOne()
	opts.SetSort(bson.D{{"_id", -1}})

	rst := collection.FindOne(context.TODO(), bson.D{}, opts)

	var max Max
	if err := rst.Decode(&max); err != nil {
		return 0, err
	}
	return max.ID + 1, nil
}

func GetTenantDateBase(c *gin.Context) string {
	TenantID := c.GetHeader("TenantID")
	if TenantID == "" {
		return "test"
	}
	return TenantID
}

// Find 单一文档查询, 返回结构体数组
// *
//func (ch *BaseHandler) Find(filter interface{})  error {
//collection := store.ClientMongo.Database(ch.DatabaseName).Collection(ch.CollectionName)
//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//defer cancel()
//
//find, err := collection.Find(ctx, filter)
//if err != nil {
//	return find, err
//}
////ch.SetCollection()
////var ts []
//for find.Next(ctx) {
//	var t T
//	err := find.Decode(&t)
//	if err != nil {
//		logs.LG.Error(err.Error())
//		return nil, err
//	}
//	ts = append(ts, t)
//}
//return ts, nil
//}

// InsertOne 插入指定对象，包含事务，连续多个插入时，请勿使用。
// *
func (ch *BaseHandler) InsertOne(model interface{}) (T, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	// 开启事务
	if session, err := store.StartTransaction(); err == nil {
		// 执行事务
		err = mongo.WithSession(context.Background(), session, func(sessionContext mongo.SessionContext) error {
			ID, err := ch.GetNextID()
			err = ch.setID(model, ID)
			if err != nil {
				return err
			}
			// 柜门表单
			dhDoc := store.ClientMongo.Database(ch.DatabaseName).Collection(ch.CollectionName)
			_, err = dhDoc.InsertOne(ctx, model)
			if err != nil {
				return err
			}
			// 提交事务
			err = store.CommitTransaction(session)
			return nil
		})
		if err != nil {
			return model, fmt.Errorf("Inner error: " + err.Error())
		}
	}
	return model, nil
}

// UpdateOne 指定_id删除，已包含事务，连续多个删除时，请勿使用。
// *
func (ch *BaseHandler) UpdateOne(model interface{}) error {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 开启事务
	session, err := store.StartTransaction()
	if err != nil {
		return err
	}

	// 执行事务
	err = mongo.WithSession(context.Background(), session, func(sessionContext mongo.SessionContext) error {
		collection := store.ClientMongo.Database(ch.DatabaseName).Collection(ch.CollectionName)
		bsonData, err := bson.Marshal(model)
		var doc *bson.D
		err = bson.Unmarshal(bsonData, &doc)
		if err != nil {
			return err
		}
		ID, err := ch.getID(model)
		if err != nil {
			return err
		}

		// 1. 定义查询条件
		filter := bson.D{{"ID", ID}}

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

// FindOne 指定_id 查找。
// *
func (ch *BaseHandler) FindOne(ID int) (*mongo.SingleResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	collection := store.ClientMongo.Database(ch.DatabaseName).Collection(ch.CollectionName)
	filter := bson.D{{"ID", ID}}
	find := collection.FindOne(ctx, filter)

	return find, nil
}

func (ch *BaseHandler) DeleteOne(ID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	// 开启事务
	if session, err := store.StartTransaction(); err == nil {
		// 执行事务
		err = mongo.WithSession(context.Background(), session, func(sessionContext mongo.SessionContext) error {
			// 柜门表单
			collection := store.ClientMongo.Database(ch.DatabaseName).Collection(ch.CollectionName)
			filter := bson.M{
				"ID": ID,
			}
			_, err = collection.DeleteOne(ctx, filter)
			if err != nil {
				return err
			}
			// 提交事务
			err = store.CommitTransaction(session)
			return nil
		})
		if err != nil {
			return fmt.Errorf("Inner error: " + err.Error())
		}
	}
	return nil
}

func (ch *BaseHandler) setID(model T, ID int) error {
	val := reflect.ValueOf(model)
	//val = reflect.Indirect(val)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("provided object is not a struct")
	}
	//vId := val.Field(0)
	//vb := reflect.ValueOf(ID)
	//vId.Set(vb)
	fieldVal := val.FieldByName("ID")
	if !fieldVal.IsValid() {
		return fmt.Errorf("field %s not found", "ID")
	}
	if !fieldVal.CanSet() {
		return fmt.Errorf("field %s cannot be set", "ID")
	}
	inVal := reflect.ValueOf(ID)
	if fieldVal.Type() != inVal.Type() {
		return fmt.Errorf("provided value type %s doesn't match field type %s", inVal.Type(), fieldVal.Type())
	}
	//
	fieldVal.SetInt(int64(ID))
	return nil
}

func (ch *BaseHandler) getID(model T) (int, error) {
	val := reflect.ValueOf(model)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return -1, fmt.Errorf("provided object is not a struct")
	}
	fieldVal := val.FieldByName("ID")
	if !fieldVal.IsValid() {
		return -1, fmt.Errorf("field %s not found", "ID")
	}
	intValue := fieldVal.Interface()
	ID, ok := intValue.(int)
	if !ok {
		return -1, fmt.Errorf("get ID error")
	} else {
		return ID, nil
	}
}

//func (ch *BaseHandler) SetCollection(m interface{}) (int, error) {
//	// 使用reflect.TypeOf获取结构体的类型信息
//	t := reflect.TypeOf(m)
//
//	runes := []rune(ch.CollectionName)
//	runes[0] = unicode.ToUpper(runes[0])
//	collectionName := string(runes)
//
//	// 通过结构体名称获取结构体的类型
//	t := reflect.TypeOf(GetTypeByName(collectionName))
//	if t.Kind() == reflect.Ptr {
//		t = t.Elem()
//	}
//
//	val := reflect.ValueOf(model)
//	if val.Kind() == reflect.Ptr {
//		val = val.Elem()
//	}
//	if val.Kind() != reflect.Struct {
//		return -1, fmt.Errorf("provided object is not a struct")
//	}
//	fieldVal := val.FieldByName("ID")
//	if !fieldVal.IsValid() {
//		return -1, fmt.Errorf("field %s not found", "ID")
//	}
//	intValue := fieldVal.Interface()
//	ID, ok := intValue.(int)
//	if !ok {
//		return -1, fmt.Errorf("get ID error")
//	} else {
//		return ID, nil
//	}
//}

// GetTypeByName 通过结构体名称获取结构体类型
//func GetTypeByName(name string) interface{} {
//	switch name {
//	case "ExampleStruct":
//		return ExampleStruct{}
//	default:
//		panic("unknown type: " + name)
//	}
//}
