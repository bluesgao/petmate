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

func InitMongoClient(url string) (*mongo.Client, error) {
	log.Print("mongo client init start...")

	//连接mongo服务
	var err error
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if MongoCli, err = mongo.Connect(ctx, options.Client().ApplyURI(url)); err != nil {
		log.Printf("mongo connect error:%+v", err)
		os.Exit(2)
		//return nil, err
	}
	//判断服务是否可用
	if err = MongoCli.Ping(ctx, readpref.Primary()); err != nil {
		log.Printf("mongo client ping error:%+v", err)
		os.Exit(2)
		//return nil, err
	}
	log.Print("mongo client init end...")
	return MongoCli, nil
}
