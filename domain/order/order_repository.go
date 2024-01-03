package order

import (
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&Order{})
	if err != nil {
		log.Print(err)
	}
}

// 根据订单id查找
func (r *Repository) FindByOrderID(oid uint) (*Order, error) {
	var currentOrder *Order
	if err := r.db.Where("IsCanceled = ?", false).Where("ID", oid).First(&currentOrder).Error; err != nil {
		return nil, err
	}
	return currentOrder, nil

}

// Update 更新订单
func (r *Repository) Update(order Order) error {
	tx := r.db.Save(&order)
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}

// Create 创建订单
func (r *Repository) Create(order *Order) error {
	tx := r.db.Create(order)
	if tx.Error != nil {
		return nil
	}
	return nil
}

// GetAll 获取所有订单
func (r *Repository) GetAll(pageIndex int, pageSize int, uid uint) ([]Order, int) {
	var orders []Order
	var count int64
	r.db.Where("IsCanceled = ? and UserID= ?", false, uid).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&orders).Count(&count)
	for i, order := range orders {
		r.db.Where("OrderID = ?", order.ID).Find(&orders[i].OrderedItems)
		for i2, item := range orders[i].OrderedItems {
			r.db.Where("ID = ?", item.ProductID).First(&orders[i].OrderedItems[i2].Product)
		}
	}
	return orders, int(count)
}
