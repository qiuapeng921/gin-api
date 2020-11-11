package app

import (
	"gin-api/app/utility/mysql"
	"gin-api/app/utility/redis"
	"gin-api/config"
	"net"
	"net/http"
	"time"
)

// newRedis
func newRedis(conf *config.Config) *redis.Client {
	return redis.New(conf.Redis())
}

// newDBManager
func newDBManager(conf *config.Config) *mysql.Manager {
	return mysql.NewManager(conf.Mysql())
}

// newHttpClient
func newHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{ // 配置连接池
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			IdleConnTimeout: time.Duration(10) * time.Second,
		},
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       time.Duration(10) * time.Second,
	}
}