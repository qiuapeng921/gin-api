package db

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"os"
)

var client *elastic.Client

func InitElastic() {
	var err error
	host := os.Getenv("ELASTIC_HOST")
	client, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(host))
	if err != nil {
		panic(fmt.Sprintf("elastic初始化失败 错误信息：%s",err.Error()))
	}
	ctx := context.Background()
	_, _, err = client.Ping(host).Do(ctx)
	if err != nil {
		panic(fmt.Sprintf("elastic连接检测 错误原因：%s",err.Error()))
	}
	fmt.Println("elastic连接成功")
}

func ElasticClient() *elastic.Client {
	return client
}
