package product

import (
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	db *gorm.DB
}

// NewProductRepository 实例化
func NewProductRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Migration 生成表
func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&Product{}) //生成表结构
	if err != nil {
		log.Print(err)
	}
}

// Update 更新
func (r *Repository) Update(updateProduct Product) error {
	sku, err := r.FindBySKU(updateProduct.SKU)
	if err != nil {
		return err
	}
	tx := r.db.Model(&sku).Updates(&updateProduct)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// SearchByString 搜索并且返回分页结果
func (r *Repository) SearchByString(str string, pageIndex, pageSize int) ([]Product, int) {
	var products []Product
	convertedStr := "%" + str + "%"
	var count int64
	r.db.Where("IsDeleted = ? and (Name like ? or SKU like ?)", false, convertedStr, convertedStr).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&products).Count(&count)
	return products, int(count)
}

// FindBySKU 根据sku查找
func (r *Repository) FindBySKU(sku string) (*Product, error) {
	var product *Product
	tx := r.db.Where("IsDeleted=?", 0).Where(Product{SKU: sku}).First(&product)
	if tx.Error != nil {
		return nil, ErrProductNotFound
	}
	return product, nil
}

// Create 创建
func (r *Repository) Create(product *Product) error {
	tx := r.db.Create(product)
	return tx.Error
}

// GetAll 查询所有商品
func (r *Repository) GetAll(pageIndex, pageSize int) ([]Product, int) {
	var products []Product
	var count int64
	r.db.Where("IsDeleted = ?", 0).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&products).Count(&count)
	return products, int(count)
}

// Delete 根据sku删除
func (r *Repository) Delete(sku string) error {
	bySKU, err := r.FindBySKU(sku)
	if err != nil {
		return ErrProductNotFound
	}
	bySKU.IsDeleted = true
	err = r.db.Save(bySKU).Error
	return err
}
