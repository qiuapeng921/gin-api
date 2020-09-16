package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"strconv"
	"time"
)

type MongoClient struct {
	client *mongo.Client
}

var MoClient MongoClient

//初始化
func InitMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	host := os.Getenv("MONGO_HOST")
	user := os.Getenv("MONGO_USERNAME")
	pass := os.Getenv("MONGO_PASSWORD")
	database := os.Getenv("MONGO_DATABASE")
	uri := fmt.Sprintf("mongodb://%s:%s@%s/%s", user, pass, host, database)

	maxPool, _ := strconv.ParseUint(os.Getenv("MONGO_MAX_POOL"), 10, 64)

	var err error
	// 连接数据库
	MoClient.client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri).SetMaxPoolSize(maxPool))
	if err != nil {
		panic(fmt.Sprintf("Mongo初始化失败 错误信息：%s", err.Error()))
	}
	// 判断服务是不是可用
	err = MoClient.client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err.Error())
	}
	MoClient.client.Database(database)

	fmt.Println("mongodb连接成功")
}