package main

import (
	"fmt"
	"server_mall/internal/routers"
	"server_mall/pkg/job/bill"
	"server_mall/store"
	"time"
)

func init() {
	redis := new(store.Redis)
	err := redis.Get().Ping().Err()
	if err != nil {
		panic(err)
	}
	mysql := new(store.Mysql)
	mysql.Get()
	log := new(store.Log)
	bill.Run()
	log.Get().Debug("server started at %v", time.Now())
}

func main() {
	config := new(store.Config)
	err := routers.Init().Run(fmt.Sprintf(":%d", config.Get().Http.Port))
	if err != nil {
		panic(err)
	}
}
