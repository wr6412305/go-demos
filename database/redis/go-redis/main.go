package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

// 定义一组常量
const (
	redisIP   = "127.0.0.1"
	redisPort = "6379"
	redisPwd  = ""
	redisDB   = 0
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     redisIP + ":" + redisPort, // ip:port
		Password: redisPwd,                  // redis连接密码
		DB:       redisDB,                   // 选择的redis库
		PoolSize: 20,                        // 设置连接数,默认是10个连接
	})
}

func main() {
	defer client.Close()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	// base()
	// strings()
	// hash()
	// list()
	// set()
	// zset()
	// subscribe()
}
