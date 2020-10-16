package freight_template

import (
	"github.com/jinzhu/gorm"
	"time"
)

type FreightTemplate struct {
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
		Name: "freight_templates",
	}
}

func (s *Model) GetOne(id int) (ret FreightTemplate, err error) {
	err = s.Db.Table(s.Name).Where("id = ?", id).First(&ret).Error
	if err != nil {
		return
	}
	return
}
