package controllers

import (
	"context"
	"log"
	"stock/src/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

// DbEngine 关系型数据库引擎
type DbEngine struct {
	MgEngine *mongo.Client //关系型数据库引擎
	Mdb      string
}

// NewDbEngine 实例工厂
func NewDbEngine() *DbEngine {
	return &DbEngine{}
}

// Open 开启连接池
func (d *DbEngine) Open(mg, mdb string, initdb bool) error {
	d.Mdb = mdb
	ops := options.Client().ApplyURI(mg)
	p := uint64(39000)
	ops.MaxPoolSize = &p
	ops.WriteConcern = writeconcern.New(writeconcern.J(true), writeconcern.W(1))
	ops.ReadPreference = readpref.PrimaryPreferred()
	db, err := mongo.NewClient(ops)

	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.Connect(ctx)
	if err != nil {
		return err
	}

	//err = db.Ping(ctx, readpref.PrimaryPreferred())
	//if err != nil {
	//	log.Println("ping err", err)
	//}

	d.MgEngine = db

	//初始化数据库
	if initdb {
		var session *mongo.Client
		session, err = mongo.NewClient(ops)
		if err != nil {
			panic(err)
		}
		err = session.Connect(context.Background())
		if err != nil {
			panic(err)
		}
		defer session.Disconnect(context.Background())
		//order表
		// res := InitDbAndColl(session, mdb, models.TEnterpriseIndicator, &models.Enterprise{})
		ord := session.Database(mdb).Collection(models.TEnterpriseIndicator)
		indexView := ord.Indexes()
		_, err = indexView.CreateMany(context.Background(), []mongo.IndexModel{
			// {
			// 	Keys: bsonx.Doc{{"biz_id", bsonx.Int32(1)}},
			// },
			// {
			// 	Keys: bsonx.Doc{{"seller_id", bsonx.Int32(1)}},
			// },
			// {
			// 	Keys: bsonx.Doc{{"buyer_id", bsonx.Int32(1)}},
			// },
			// {
			// 	Keys: bsonx.Doc{{"order_status", bsonx.Int32(1)}},
			// },
			// {
			// 	Keys: bsonx.Doc{{"order_sn", bsonx.Int32(1)}},
			// },
			// {
			// 	Keys: bsonx.Doc{{"create_date", bsonx.Int32(-1)}},
			// },
		})
		if err != nil {
			log.Println(err)
		}

	}

	return nil
}

// GetColl 获取表名
func (d *DbEngine) GetColl(coll string) *mongo.Collection {
	col, _ := d.MgEngine.Database(d.Mdb).Collection(coll).Clone()
	return col
}

// InitDbAndColl 建立数据库
func InitDbAndColl(session *mongo.Client, db, coll string, model map[string]interface{}) map[string]interface{} {
	tn, _ := session.Database(db).ListCollections(context.Background(), bson.M{"name": coll})
	if tn.Next(context.Background()) == false {
		session.Database(db).RunCommand(context.Background(), bson.D{primitive.E{Key: "create", Value: coll}})
	}
	result := session.Database(db).RunCommand(context.Background(), bson.D{primitive.E{Key: "collMod", Value: coll}, primitive.E{Key: "validator", Value: model}})
	var res map[string]interface{}
	err := result.Decode(&res)
	if err != nil {
		log.Println(err)
	}
	return res
}

// Close 关闭连接池
func (d *DbEngine) Close() {
	d.MgEngine.Disconnect(context.Background())
}
