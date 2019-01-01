package db

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

var DB redis.Conn

func initRedis() {
	port := viper.GetString("redis.port")
	fmt.Println(port)
	con, err := redis.Dial("tcp", fmt.Sprintf(":%s", "6377"))
	if err != nil {
		panic("can not connect to redis")
	}
	DB = con
}

func Start() {
	initRedis()
}
