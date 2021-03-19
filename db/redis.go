package db

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
	"time"
)

var (
	RedisDB *redis.Client
	ctx     context.Context
)

func init() {
	RedisDB = redis.NewClient(&redis.Options{
		//todo 修改端口
		//Addr:         config.SuRedis.Host + ":" + config.SuRedis.Port,
		//Password:     config.SuRedis.PassWord,
		//DB:           config.SuRedis.DatabasesName,
		//PoolSize:     config.SuRedis.POOL_SIZE,
		//MinIdleConns: config.SuRedis.MIN_IDLE_CONNS,
		Addr:     "127.0.0.1:6379", // 要连接的redis IP:port
		Password: "root",           // redis 密码
		DB:       0,                // 要连接的redis 库
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//ctx := context.Background()
	defer cancel()
	result, err := RedisDB.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println(result)
}
