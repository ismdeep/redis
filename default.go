package redis

import (
	"github.com/go-redis/redis/v8"
)

// Client struct
type Client struct {
	Client *redis.Client
}

// New a client
func New(conf *redis.Options) *Client {
	c := redis.NewClient(conf)
	return &Client{Client: c}
}
