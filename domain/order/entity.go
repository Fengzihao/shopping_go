package order

import (
	"gorm.io/gorm"
	"shopping_go/domain/product"
	"shopping_go/domain/user"
)

// Order 订单结构体
type Order struct {
	gorm.Model
	UserID       uint
	User         user.User     `gorm:"foreignKey:ID;references:UserID" json:"-"`
	OrderedItems []OrderedItem `gorm:"foreignKey:OrderID"`
	TotalPrice   float32
	IsCanceled   bool
}

// OrderedItem 订单项结构体
type OrderedItem struct {
	gorm.Model
	Product    product.Product `gorm:"foreignKey:ProductID"`
	ProductID  uint
	Count      int
	OrderID    uint
	IsCanceled bool
}

func NewOrder(uid uint, items []OrderedItem) *Order {
	var totalPrice float32 = 0.0
	for _, item := range items {
		totalPrice += item.Product.Price
	}
	return &Order{
		UserID:       uid,
		OrderedItems: items,
		TotalPrice:   totalPrice,
		IsCanceled:   false,
	}
}

func NewOrderedItem(count int, pid uint) *OrderedItem {
	return &OrderedItem{
		Count:      count,
		ProductID:  pid,
		IsCanceled: false,
	}
}
