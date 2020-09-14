package app

import (
	"gin-api/config"
	"gin-api/helpers/redis"
)

// newRedis
func newRedis(conf *config.Config) *redis.Client {
	return redis.New(conf.Redis())
}
