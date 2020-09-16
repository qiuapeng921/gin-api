package redis

import "time"

type Config struct {
	Host        string
	Port        int
	Auth        string
	PoolSize    int
	MaxRetries  int
	IdleTimeout time.Duration
}
