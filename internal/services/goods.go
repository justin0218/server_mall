package services

import (
	"server_mall/api"
	"server_mall/internal/models/goods"
)

type GoodsService struct {
}

func (s *GoodsService) Detail(goodsId string) (ret goods.Goods, err error) {
	db := api.Mysql.Get()
	return goods.NewModel(db).Detail(goodsId)
}
