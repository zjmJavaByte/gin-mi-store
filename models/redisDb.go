package models

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gopkg.in/ini.v1"
	"os"
)


var (
	RedisDb *redis.Client
)

func init() {

	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	ip := cfg.Section("redis").Key("ip").String()
	port := cfg.Section("redis").Key("port").String()
	password := cfg.Section("redis").Key("password").String()
	database,_ := cfg.Section("redis").Key("database").Int()
	var ctx = context.Background()
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v",ip,port),
		Password: password, // no password set
		DB:       database,  // use default DB
	})
	_, err = RedisDb.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
	}


}
