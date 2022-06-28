package redis

import (
	"context"
	"github.com/go-redis/redis/v9"
	"mall/global/config"
	"mall/global/log"
	"sync"
	"time"
)

type RedisHelper struct {
	*redis.Client
}

var redisHelper *RedisHelper
var ctx = context.Background()

var redisOnce sync.Once

func init() {
	rdb := redis.NewClient(&redis.Options{
		Addr:         "",
		Password:     "",
		DB:           0,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})

	redisConfig := config.GetConfig().Redis

	rdb.Options().Addr = redisConfig.Addr
	rdb.Options().Password = redisConfig.Password
	rdb.Options().DB = redisConfig.DB

	if redisConfig.DialTimeout != 0 {
		rdb.Options().DialTimeout = time.Second * time.Duration(redisConfig.DialTimeout)
	}
	if redisConfig.WriteTimeout != 0 {
		rdb.Options().WriteTimeout = time.Second * time.Duration(redisConfig.WriteTimeout)
	}
	if redisConfig.ReadTimeout != 0 {
		rdb.Options().ReadTimeout = time.Second * time.Duration(redisConfig.ReadTimeout)
	}
	if redisConfig.PoolSize != 0 {
		rdb.Options().PoolSize = redisConfig.PoolSize
	}
	if redisConfig.PoolTimeout != 0 {
		rdb.Options().PoolTimeout = time.Second * time.Duration(redisConfig.PoolTimeout)
	}

	redisOnce.Do(func() {
		rdh := new(RedisHelper)
		rdh.Client = rdb
		redisHelper = rdh
	})
	ctx := context.Background()

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.Logger.Fatal("fail to connect redis, err: " + err.Error())
		return
	}
	log.Logger.Info("success to connect redis")
}
func GetRedis() (*RedisHelper, context.Context) {
	return redisHelper, ctx
}
