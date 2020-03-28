package redis

import (
	"github.com/gomodule/redigo/redis"
	"myblog-api/app/config"
	"time"
)

type RedisCli struct {
	Pool *redis.Pool
}

var RedisClient *RedisCli

func Default() () {
	rediscli := new(RedisCli)
	rediscli.Pool = &redis.Pool{
		MaxIdle:     256,
		MaxActive:   0,
		IdleTimeout: time.Duration(120),
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				config.Configs.RedisHost,
				redis.DialReadTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialWriteTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialConnectTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialDatabase(config.Configs.RedisDb),
				redis.DialPassword(config.Configs.RedisPwd),
			)
		},
	}
	RedisClient = rediscli
}
