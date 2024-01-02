package product

import (
	"gorm.io/gorm"
	"shopping_go/domain/category"
)

type Product struct {
	gorm.Model
	Name       string            //商品名称
	SKU        string            //唯一编码
	Desc       string            //商品描述
	StockCount int               //库存
	Price      float32           //价格
	CategoryID uint              //商品分类id
	Category   category.Category `json:"-"` //商品分类
	IsDeleted  bool
}

// NewProduct 商品结构体实例
func NewProduct(name string, desc string, stockCount int, price float32, cid uint) *Product {
	return &Product{
		Name:       name,
		Desc:       desc,
		StockCount: stockCount,
		Price:      price,
		CategoryID: cid,
		IsDeleted:  false,
	}
}
