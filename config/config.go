package config

import (
	"gin-api/app/utility/mysql"
	"gin-api/app/utility/redis"
	"github.com/spf13/viper"
	"os"
	"strconv"
	"strings"
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
	mysql *mysql.Config
}

func (conf *Config) Redis() *redis.Config {
	if conf.redis == nil {
		conf.mu.Lock()
		defer conf.mu.Unlock()
		address := strings.Split(os.Getenv("REDIS_HOST"), ":")
		port, _ := strconv.Atoi(address[1])
		Idle, _ := strconv.Atoi(os.Getenv("MAX_IDLE"))
		Active, _ := strconv.Atoi(os.Getenv("MAX_ACTIVE"))
		conf.redis = &redis.Config{
			Host:        address[0],
			Port:        port,
			Auth:        os.Getenv("REDIS_PASSWORD"),
			PoolSize:    Idle,
			MaxRetries:  Active,
			IdleTimeout: time.Duration(10) * time.Second,
		}
	}
	return conf.redis
}

// mysql
func (conf *Config) Mysql() *mysql.Config {
	if conf.mysql == nil {
		conf.mu.Lock()
		defer conf.mu.Unlock()
		host := os.Getenv("DB_HOST")
		database := os.Getenv("DB_DATABASE")
		username := os.Getenv("DB_USERNAME")
		password := os.Getenv("DB_PASSWORD")
		charset := os.Getenv("DB_CHARSET")
		maxIdle, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE"))
		maxOpen, _ := strconv.Atoi(os.Getenv("DB_MAX_OPEN"))
		conf.mysql = &mysql.Config{
			Host:     host,
			Database: database,
			Username: username,
			Password: password,
			Charset:  charset,
			MaxIdle:  maxIdle,
			MaxOpen:  maxOpen,
		}
	}
	return conf.mysql
}