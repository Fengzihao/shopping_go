package cart

import (
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	db *gorm.DB
}

// NewCartRepository 实例化
func NewCartRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Migration 创建表
func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&Cart{})
	if err != nil {
		log.Print(err)
	}
}

// Update 更新
func (r *Repository) Update(cart Cart) error {
	tx := r.db.Save(cart)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// FindOrCreateByUserID 根据用户id查找或者创建购物车
func (r *Repository) FindOrCreateByUserID(userId uint) (*Cart, error) {
	var cart *Cart
	err := r.db.Where(&Cart{UserID: userId}).Attrs(NewCart(userId)).FirstOrCreate(&cart).Error
	if err != nil {
		return nil, err
	}
	return cart, nil
}

// FindByUserID 根据用户id查找购物车
func (r *Repository) FindByUserID(userId uint) (*Cart, error) {
	var cart *Cart
	err := r.db.Where(&Cart{UserID: userId}).Attrs(NewCart(userId)).First(&cart).Error
	if err != nil {
		return nil, err
	}
	return cart, nil
}

// ItemRepository 购物车中的商品结构体
type ItemRepository struct {
	db *gorm.DB
}

// NewCartItemRepository 实例化
func NewCartItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{
		db: db,
	}
}

// Migration 生成item表
func (i *ItemRepository) Migration() {
	err := i.db.AutoMigrate(&Item{})
	if err != nil {
		log.Print(err)
	}
}

// Update Item 更新购物车中的商品
func (i *ItemRepository) Update(item Item) error {
	tx := i.db.Save(&item)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// FindByID 根据商品id和购物车id查找item
func (i *ItemRepository) FindByID(pid uint, cid uint) (*Item, error) {
	var item *Item
	err := i.db.Where(&Item{ProductID: pid, CartID: cid}).First(&item).Error
	if err != nil {
		return nil, err
	}
	return item, nil
}

// Create 创建Item
func (i *ItemRepository) Create(item *Item) error {
	tx := i.db.Create(item)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// GetItems 查找购物车中所有的商品 item
func (i *ItemRepository) GetItems(cartId uint) ([]Item, error) {
	var cartItems []Item
	tx := i.db.Where(&Item{CartID: cartId}).Find(&cartItems)
	if tx.Error != nil {
		return nil, tx.Error
	}
	for i2, item := range cartItems {
		err := i.db.Model(item).Association("Product").Find(&cartItems[i2].Product)
		if err != nil {
			return nil, err
		}
	}
	return cartItems, nil
}
