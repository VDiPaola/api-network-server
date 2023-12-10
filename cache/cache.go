package caching

import (
	"encoding/json"
	"time"

	"github.com/VDiPaola/api-network-server/models"
	"github.com/go-redis/redis/v7"
)

type RedisCache struct {
	host    string
	db      int
	url     string
	expires time.Duration
	client  *redis.Client
}

func NewRedisCache(host string, db int, expires time.Duration) ResponseCache {
	return &RedisCache{
		host:    host,
		db:      db,
		expires: expires,
	}
}

func NewRedisCacheURL(url string, expires time.Duration) ResponseCache {
	return &RedisCache{
		url:     url,
		expires: expires,
	}
}

func (cache *RedisCache) getClient() *redis.Client {
	//set client if not exist
	if cache.client == nil {
		if len(cache.url) > 1 {
			opt, err := redis.ParseURL(cache.url)
			if err != nil {
				panic(err)
			}
			cache.client = redis.NewClient(opt)
		} else {
			cache.client = redis.NewClient(&redis.Options{
				Addr:     cache.host,
				Password: "",
				DB:       cache.db,
			})
		}
	}
	return cache.client
}

func (cache *RedisCache) Set(key string, value *models.Node) {
	client := cache.getClient()

	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	client.Set(key, json, cache.expires)
}

func (cache *RedisCache) Get(key string) *models.Node {
	client := cache.getClient()

	value, err := client.Get(key).Result()
	if err != nil {
		panic(err)
	}

	node := models.Node{}
	err = json.Unmarshal([]byte(value), &node)

	if err != nil {
		panic(err)
	}

	return &node

}
