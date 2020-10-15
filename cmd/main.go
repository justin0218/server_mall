package main

import (
	"fmt"
	"server_mall/api"
	"server_mall/configs"
	"server_mall/internal/routers"
	"server_mall/pkg/job/bill"
	//"server_mall/job"
)

func main() {
	api.Log.Get().Debug("starting...")
	err := api.Rds.Get().Ping().Err()
	if err != nil {
		panic(err)
	}
	api.Mysql.Get()
	bill.Run()
	fmt.Println("server started!!!")
	err = routers.Init().Run(fmt.Sprintf(":%d", configs.Dft.Get().Http.Port))
	if err != nil {
		panic(err)
	}
}
