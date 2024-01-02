package cart

import (
	"gorm.io/gorm"
	"shopping_go/domain/product"
	"shopping_go/domain/user"
)

// Cart 购物车结构体
type Cart struct {
	gorm.Model
	UserID uint
	User   user.User `gorm:"foreignKey:ID;references:UserID"`
}

func NewCart(uid uint) *Cart {
	return &Cart{
		UserID: uid,
	}
}

// Item Item结构体 购物车中的商品
type Item struct {
	gorm.Model
	Product   product.Product `gorm:"foreignKey:ProductID"`
	ProductID uint
	Count     int
	CartID    uint
	Cart      Cart `gorm:"foreignKey:CartID" json:"-"`
}

// NewCartItem 创建Item
func NewCartItem(productId uint, cartId uint, count int) *Item {
	return &Item{
		ProductID: productId,
		Count:     count,
		CartID:    cartId,
	}
}
