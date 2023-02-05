package config

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/thanh-vt/splash-inventory-service/internal"
	"os"
	"time"
)

type RedisCache struct {
}

func ConnectRedis() {
	internal.Redis = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       0,                           // use default DB
	})
	internal.RedisCtx = context.Background()
}

func (r RedisCache) Get(k string) (interface{}, bool) {
	val, err := internal.Redis.Get(internal.RedisCtx, k).Result()
	if err != nil {
		if err == redis.Nil {
			return val, false
		}
		panic(err)
	}
	return val, false
}

func (r RedisCache) GetWithExpiration(k string) (interface{}, time.Time, bool) {
	val, exists := r.Get(k)
	return val, time.Time{}, exists
}

func (r RedisCache) Set(k string, x interface{}) {
	//val := reflect.ValueOf(x).Elem()
	//jwk := val.FieldByName("jwk").Interface().(*jose.JSONWebKey)
	_, err := internal.Redis.Set(internal.RedisCtx, k, fmt.Sprintf("%v", x), 0).Result()
	if err != nil {
		if err != redis.Nil {
			panic(err)
		}
	}
}
