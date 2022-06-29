package queries

import (
	"context"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v9"
)

type CachingQueries struct {
	Redis *redis.Client
}

func (q *CachingQueries) Get(key string) (string, error) {
	value, err := q.Redis.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return "", errors.New("key doesn't exists")
	}
	if err != nil {
		return "", err
	}
	//	q.Redis.LRange().Result()

	return value, nil
}

func (q *CachingQueries) Set(key, value string) error {
	minuteCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))

	// q.Redis.RPush()

	if err := q.Redis.Set(context.Background(), key, value, time.Minute*time.Duration(minuteCount)).Err(); err != nil {
		return err
	}

	return nil
}
