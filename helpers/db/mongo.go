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

var mongoClient *mongo.Client

//初始化
func InitMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	host := os.Getenv("MONGO_HOST")
	user := os.Getenv("MONGO_USERNAME")
	pass := os.Getenv("MONGO_PASSWORD")
	uri := fmt.Sprintf("mongodb://%s:%s@%s/admin", user, pass, host)

	maxPool, _ := strconv.ParseUint(os.Getenv("MONGO_MAX_POOL"), 10, 64)
	var err error
	// 连接数据库
	mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(uri).SetMaxPoolSize(maxPool))
	if err != nil {
		panic(fmt.Sprintf("Mongo初始化失败 错误信息：%s",err.Error()))
	}
	// 判断服务是不是可用
	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		panic(err.Error())
	}
	database := os.Getenv("MONGO_DATABASE")
	mongoClient.Database(database)
	fmt.Println("mongodb连接成功")
}

func MongoClient() *mongo.Client {
	return mongoClient
}