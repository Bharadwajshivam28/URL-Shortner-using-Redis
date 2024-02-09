package database

import (
	"context"
	"os"
	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func CreateClient(dbNo int) *redis.Client{
	rbd := redis.NewClient(&redis.Options{
		Addr: os.Get("DB_ADDR"),
		Password: os.Get("DB_PASS"),
		DB: dbNo,
	})
	return rbd
}