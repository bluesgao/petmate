package model

import (
	"github.com/go-redis/redis"
	"log"
	"os"
)

func InitRedisClient(addr string, passwd string, poolsize int) (*redis.Client, error) {
	log.Print("redis client init start...")

	option := redis.Options{
		Addr:     addr,     //ip:port
		Password: passwd,   //密码
		DB:       0,        //数据库
		PoolSize: poolsize, //连接池
	}

	redisCli := redis.NewClient(&option)

	//测试reids server能否联通
	if ret, err := redisCli.Ping().Result(); err != nil {
		//配置文件错误，退出进程
		log.Print("redis client init error ...")
		os.Exit(2)
		//return nil, err
	}else{
		log.Print("redis client init end..."+ret)
	}
	return redisCli, nil
}
