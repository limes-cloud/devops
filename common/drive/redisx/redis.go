package redisx

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type Config struct {
	Host string
	Pass string
}

var Client *redis.Client

func NewClient(conf Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Host,
		Password: conf.Pass,
	})

	if err := client.Ping(context.TODO()).Err(); err != nil {
		panic(err)
	}
	Client = client
	return client
}
