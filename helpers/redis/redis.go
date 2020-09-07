package redis

import (
	"fmt"
	"github.com/go-redis/redis/v7"
)

// Client
type Client struct {
	conf *Config
	*redis.Client
}

func New(conf *Config) *Client {
	return &Client{conf:conf}
}

func (client *Client) Connect() error {
	connection := redis.NewClient(client.conf.Options())
	_, err := connection.Ping().Result()
	if err != nil {
		return err
	}
	fmt.Println("Connect redis success:", client.conf.Host)
	client.Client = connection
	return nil
}