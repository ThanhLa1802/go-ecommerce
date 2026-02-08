package initialize

import (
	"context"
	"fmt"
	"go-ecommerce-backend-api/global"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: r.Password, // nop password
		DB:       r.Database, // default DB
		PoolSize: 10,
	})

	// test ping
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("Redis is connecting!")
	global.Rdb = rdb
}
