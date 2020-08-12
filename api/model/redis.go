package model

import (
	"github.com/go-redis/redis/v8"
)

type redisModel struct {
	redisClient *redis.Client
}
