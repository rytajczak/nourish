package main

import (
	"time"

	"github.com/charmbracelet/log"
	"github.com/go-redis/redis"
)

type Cache struct {
	client *redis.Client
}

func NewCache(addr string) *Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	if err := client.Ping().Err(); err != nil {
		log.Fatal("Couldn't connect to cache", "err", err)
	}

	return &Cache{
		client: client,
	}
}

func (c *Cache) Set(key string, value interface{}, exp time.Duration) error {
	return c.client.Set(key, value, exp).Err()
}

func (c *Cache) Get(key string) (string, error) {
	return c.client.Get(key).Result()
}
