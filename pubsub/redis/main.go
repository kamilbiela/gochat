package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

var redisPool *redis.Pool

func init() {
	redisPool = initPool(":6379", "")
	redisPool.Get()
}

func initPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			fmt.Println("Dialing redis")
			c, err := redis.Dial("tcp", server)
			if err != nil {
				log.Fatal(err)
				return nil, err
			}
			// if _, err := c.Do("AUTH", password); err != nil {
			// 	c.Close()
			// 	return nil, err
			// }
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
