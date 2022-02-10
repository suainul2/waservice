package databases

import (
	"context"
	"fmt"
	"waservice/configs"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()
var Rdb *redis.Client

func RedisRun() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     configs.Redis.Address,
		Password: configs.Redis.Password, // no password set
		DB:       configs.Redis.DB,       // use default DB
	})
	if Rdb == nil {
		Rdb = rdb
		fmt.Println("redis Ready")
	}
}
