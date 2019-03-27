package model

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

var MongoCli *mongo.Client

func InitMongoClient(url string, database string, username string, password string) (*mongo.Client, error) {
	log.Print("mongo client init start...")

	//连接mongo服务
	//认证参数设置，否则连不上
	opts := &options.ClientOptions{}
	opts.SetAuth(options.Credential{
		AuthMechanism: "SCRAM-SHA-1",
		AuthSource:    database,
		Username:      username,
		Password:      password})

	var err error

	if MongoCli, err = mongo.Connect(getContext(), options.Client().ApplyURI(url), opts); err != nil {
		log.Printf("mongo connect error:%+v", err)
		os.Exit(2)
		//return nil, err
	}
	//判断服务是否可用
	if err = MongoCli.Ping(getContext(), readpref.Primary()); err != nil {
		log.Printf("mongo client ping error:%+v", err)
		os.Exit(2)
		//return nil, err
	}
	log.Print("mongo client init end...")
	return MongoCli, nil
}

func getContext() (ctx context.Context) {
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	return
}
