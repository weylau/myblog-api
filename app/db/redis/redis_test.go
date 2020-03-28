package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"testing"
	"time"
)

func TestRedisCli(t *testing.T) {
	rediscli := new(RedisCli)
	rediscli.Pool = &redis.Pool{
		MaxIdle:     256,
		MaxActive:   0,
		IdleTimeout: time.Duration(120),
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				"172.16.57.110:6379",
				redis.DialReadTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialWriteTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialConnectTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialDatabase(0),
				//red.DialPassword(""),
			)
		},
	}
	con := rediscli.Pool.Get()
	if err := con.Err(); err != nil {
		t.Error(err)
	}
	defer con.Close()
	con.Do("set","hello","world")
	s, err := redis.String(con.Do("GET", "hello"))
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("get hello is %v", s)
}
