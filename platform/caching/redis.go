package caching

import "github.com/go-redis/redis/v9"

func RedisConnection() (*redis.Client, error) {
	client := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})

	return client, nil
}
