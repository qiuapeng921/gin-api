package config

import (
	"gin-api/helpers/redis"
	"github.com/spf13/viper"
	"os"
	"sync"
	"time"
)

// 全局配置结构体实例
func New() *Config {
	conf := new(Config)
	conf.Viper = viper.New()
	conf.mu = new(sync.Mutex)
	return conf
}

type Config struct {
	mu *sync.Mutex
	*viper.Viper
	redis *redis.Config
}

func (conf *Config) Redis() *redis.Config {
	if conf.redis == nil {
		conf.mu.Lock()
		defer conf.mu.Unlock()
		conf.redis = &redis.Config{
			Host:        conf.GetString(os.Getenv("REDIS_HOST")),
			Port:        conf.GetInt(os.Getenv("REDIS_PORT")),
			Auth:        conf.GetString(os.Getenv("REDIS_PASSWORD")),
			PoolSize:    conf.GetInt(os.Getenv("REDIS_MAX_IDLE")),
			MaxRetries:  conf.GetInt(os.Getenv("REDIS_MAX_IDLE=50\nREDIS_MAX_ACTIVE")),
			IdleTimeout: time.Duration(10) * time.Second,
		}
	}
	return conf.redis
}
