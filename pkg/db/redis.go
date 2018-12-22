package db

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

var DB redis.Conn

func InitRedis() {
	port := viper.Get("redis.port")
	con, err := redis.Dial("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic("can not connect to redis")
	}
	DB = con
}
