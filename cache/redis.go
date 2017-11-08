package cache

import (
	"github.com/garyburd/redigo/redis"
)

var RedisCli redis.Conn

func GetRedis() error {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		return err
	}
	RedisCli = c
	return nil
}
