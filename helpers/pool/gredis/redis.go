package gredis

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"os"
	"strconv"
	"time"
)

var client *redis.Client

func SetupRedis() *redis.Client {
	Address := os.Getenv("REDIS_HOST")
	Password := os.Getenv("REDIS_PASSWORD")
	Idle, _ := strconv.Atoi(os.Getenv("MAX_IDLE"))
	Active, _ := strconv.Atoi(os.Getenv("MAX_ACTIVE"))
	client = redis.NewClient(&redis.Options{
		Addr:        Address,
		Password:    Password,         // Redis账号
		DB:          0,                // Redis库
		PoolSize:    Active,           // Redis连接池大小
		MaxRetries:  Idle,             // 最大重试次数
		IdleTimeout: 10 * time.Second, // 空闲链接超时时间
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("redis连接成功")
	return client
}

func GetRedis() *redis.Client {
	return client
}