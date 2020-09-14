package app

import (
	"gin-api/config"
	"gin-api/helpers/redis"
	"go.uber.org/dig"
)

var container *dig.Container

func init() {
	container = dig.New()
	// 注入配置文件
	_ = container.Provide(config.New)
	_ = container.Provide(newRedis)
}

func Redis() (client *redis.Client) {
	_ = container.Invoke(func(cli *redis.Client) {
		client = cli
	})
	return
}