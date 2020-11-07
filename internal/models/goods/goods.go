package goods

import (
	"github.com/jinzhu/gorm"
	"server_mall/internal/models/freight_template"
	"server_mall/internal/models/sku"
	"time"
)

type Goods struct {
	Id                int       `json:"id"`
	Banners           string    `json:"banners"`
	GoodsId           string    `json:"goods_id"`
	FreightTemplateId int       `json:"freight_template_id"`
	Uid               int       `json:"uid"`
	Name              string    `json:"name"`
	Details           string    `json:"details"`
	Price             int       `json:"price"`
	Status            int       `json:"status"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type GoodsFull struct {
	Goods
	FreightTemplateInfo freight_template.FreightTemplate `json:"freight_template_info"`
	SkuInfo             []sku.SkuFull                    `json:"sku_info"`
}

type Model struct {
	Db   *gorm.DB
	Name string
}

func NewModel(db *gorm.DB) *Model {
	return &Model{
		Db:   db,
		Name: "goods",
	}
}

func (s *Model) Detail(goodsId string) (ret Goods, err error) {
	err = s.Db.Table(s.Name).Where("goods_id = ?", goodsId).First(&ret).Error
	if err != nil {
		return
	}
	return
}
