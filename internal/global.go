package internal

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"net/http"
)

var DB *gorm.DB

var Redis *redis.Client

var RedisCtx context.Context

var HttpClient *http.Client
