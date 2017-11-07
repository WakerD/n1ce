package cache

import (
	"github.com/garyburd/redigo/redis"
)

var RedisCli redis.Conn

func GetRedis() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		c.Close()
	}
	RedisCli = c
}
