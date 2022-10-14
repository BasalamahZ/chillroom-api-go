package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v9"
)

type ApiCache interface {
	Set(key string, value []interface{})
	Get(key string) []interface{}
}

type redisCache struct {
	host string
	db   int
	exp  time.Duration
}

func (c *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.host,
		Password: "",
		DB:       c.db,
	})
}

// Set implements ApiCache
func (c *redisCache) Set(key string, value []interface{}) {
	client := c.getClient()
	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	client.Set(context.Background(), key, json, c.exp*time.Second)
}

// Get implements ApiCache
func (c *redisCache) Get(key string) []interface{} {
	client := c.getClient()
	val, err := client.Get(context.Background(), key).Result()
	if err != nil {
		return nil
	}

	var data []interface{}
	err = json.Unmarshal([]byte(val), &data)
	if err != nil {
		panic(err)
	}
	return data
}

func NewRedisCache(host string, db int, exp time.Duration) ApiCache {
	return &redisCache{
		host: host,
		db:   db,
		exp:  exp,
	}
}
