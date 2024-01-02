package product

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

//生成商品唯一编码

func (p *Product) BeforeSave(tx *gorm.DB) (err error) {
	p.SKU = uuid.New().String()
	return
}
