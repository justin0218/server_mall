package sku

import (
	"github.com/jinzhu/gorm"
	"server_mall/internal/models/sku_item"
	"server_mall/internal/models/sku_type"
	"time"
)

type Sku struct {
	Id        int       `json:"id"`
	Cover     string    `json:"cover"`
	SkuId     string    `json:"sku_id"`
	GoodsId   string    `json:"goods_id"`
	ItemId    int       `json:"item_id"`
	TypeId    int       `json:"type_id"`
	Price     int       `json:"price"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SkuFull struct {
	Sku
	SkuItem sku_item.SkuItem `json:"sku_item"`
	SkuType sku_type.SkuType `json:"sku_type"`
}

type Model struct {
	Db   *gorm.DB
	Name string
}

func NewModel(db *gorm.DB) *Model {
	return &Model{
		Db:   db,
		Name: "skus",
	}
}

func (s *Model) GetBySkuId(skuId string) (ret Sku, err error) {
	err = s.Db.Table(s.Name).Where("sku_id = ?", skuId).First(&ret).Error
	if err != nil {
		return
	}
	return
}

func (s *Model) GetByGoodsId(goodsId string) (ret []Sku, err error) {
	err = s.Db.Table(s.Name).Where("goods_id = ?", goodsId).Find(&ret).Error
	if err != nil {
		return
	}
	return
}
