package services

import (
	"server_mall/api"
	"server_mall/internal/models/freight_template"
	"server_mall/internal/models/goods"
	"server_mall/internal/models/sku"
	"server_mall/internal/models/sku_item"
	"server_mall/internal/models/sku_type"
)

type GoodsService struct {
}

func (s *GoodsService) Detail(goodsId string) (ret goods.GoodsFull, err error) {
	db := api.Mysql.Get()
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
		item.SkuItem, _ = sku_item.NewModel(db).GetOneById(v.ItemId)
		item.SkuType, _ = sku_type.NewModel(db).GetOneById(v.TypeId)
		ret.SkuInfo = append(ret.SkuInfo, item)
	}
	return
}
