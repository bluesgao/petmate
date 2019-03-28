package model

import (
	"log"
	"time"
	"context"
)

type User struct {
	Userid   string `json:"userid"`
	Username string `json:"username"`
	Password string `json:password`
}

func (user *User) Create() error {
	log.Printf("create user: %+v \n", user)
	key := "u_" + user.Username

	fields := make(map[string]interface{})
	fields["username"] = user.Username
	fields["password"] = user.Password
	fields["userid"] = user.Userid

	if ret, err := RedisCli.HMSet(key, fields).Result(); err != nil {
		log.Printf("redis create user error: %+v \n", err)
	} else {
		log.Printf("redis create user ret: %+v \n", ret)
	}

	//选择数据库和集合
	collection := MongoCli.Database("cowboy").Collection("article")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//插入一条数据
	if ret, err := collection.InsertOne(ctx, user); err != nil {
		log.Printf("mongo create user error: %+v \n", err)
	} else {
		log.Printf("mongo create user ret: %+v \n", ret)
	}

	return nil
}
