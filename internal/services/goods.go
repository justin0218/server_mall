package services

import (
	"server_mall/internal/models/freight_template"
	"server_mall/internal/models/goods"
	"server_mall/internal/models/sku"
)

type GoodsService struct {
	baseService
}

func (s *GoodsService) Detail(goodsId string) (ret goods.GoodsFull, err error) {
	db := s.Mysql.Get()
	ret.Goods, err = goods.NewModel(db).Detail(goodsId)
	if err != nil {
		return
	}
	ret.FreightTemplateInfo, err = freight_template.NewModel(db).GetOne(ret.Goods.FreightTemplateId)
	if err != nil {
		return
	}
	skus, e := sku.NewModel(db).GetByGoodsId(ret.Goods.GoodsId)
	if e != nil {
		err = e
		return
	}
	for _, v := range skus {
		item := sku.SkuFull{}
		item.Sku = v
		ret.SkuInfo = append(ret.SkuInfo, item)
	}
	return
}
