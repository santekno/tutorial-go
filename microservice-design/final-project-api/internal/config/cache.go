package config

import "github.com/go-redis/redis/v8"

func OpenCache(address string, password string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password, // no password set
		DB:       0,        // use default DB
	})

	return rdb
}
