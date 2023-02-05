package config

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/thanh-vt/splash-inventory-service/internal"
	"os"
)

func ConnectRedis() error {
	internal.Redis = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       0,                           // use default DB
	})

	internal.RedisCtx = context.Background()
	_, err := internal.Redis.Ping(internal.RedisCtx).Result()
	if err != nil {
		if err == redis.Nil {
			return nil
		}
		return err
	}
	return nil
}
