package dao

import (
	"context"
	"github.com/go-redis/redis/v9"
	"mall/global/config"
	"mall/global/log"
	"time"
)

var redisDB *redis.Client

func init() {
	redisConfig := config.GetConfig().Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:         redisConfig.Host,
		Password:     redisConfig.Password,
		DB:           redisConfig.DB,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})

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

	redisDB = rdb
	ctx := context.Background()

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.Logger.Fatal("fail to connect redis, err: " + err.Error())
		return
	}
	log.Logger.Info("success to connect redis")
}
func GetRedis() *redis.Client {
	return redisDB
}
