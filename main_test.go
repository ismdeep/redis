package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"testing"
)

var testClient *Client

func TestMain(m *testing.M) {
	testClient = New(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	s := testClient.Client.Ping(context.Background())
	if err := s.Err(); err != nil {
		panic(err)
	}
	m.Run()
}
