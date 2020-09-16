package db

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"os"
	"strconv"
	"time"
)

var redisClient *redis.Client

func InitRedis() *redis.Client {
	Address := os.Getenv("REDIS_HOST")
	Password := os.Getenv("REDIS_PASSWORD")
	Idle, _ := strconv.Atoi(os.Getenv("MAX_IDLE"))
	Active, _ := strconv.Atoi(os.Getenv("MAX_ACTIVE"))
	redisClient = redis.NewClient(&redis.Options{
		Addr:        Address,
		Password:    Password,         // Redis账号
		DB:          0,                // Redis库
		PoolSize:    Active,           // Redis连接池大小
		MaxRetries:  Idle,             // 最大重试次数
		IdleTimeout: 10 * time.Second, // 空闲链接超时时间
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("redis连接成功")
	return redisClient
}

func RedisClient() *redis.Client {
	return redisClient
}