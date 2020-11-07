package services

import "server_mall/store"

type baseService struct {
	Mysql store.Mysql
	Redis store.Redis
}
