package sku_type

import (
	"github.com/jinzhu/gorm"
	"time"
)

type SkuType struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
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
		Name: "sku_types",
	}
}

func (s *Model) GetOneById(id int) (ret SkuType, err error) {
	err = s.Db.Table(s.Name).Where("id = ?", id).First(&ret).Error
	if err != nil {
		return
	}
	return
}
