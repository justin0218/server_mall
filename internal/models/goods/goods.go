package goods

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Goods struct {
	GoodsId   string    `json:"goods_id"`
	Uid       int       `json:"uid"`
	Name      string    `json:"name"`
	Details   string    `json:"details"`
	Price     int       `json:"price"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
