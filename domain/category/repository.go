package category

import (
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Migration 生成商品分类数据库表
func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&Repository{})
	if err != nil {
		log.Print(err)
	}
}

// InsertSampleData 创建商品分类测试数据
func (r *Repository) InsertSampleData() {
	categories := []Category{
		{Name: "CAT1", Desc: "Category 1"},
		{Name: "CAT2", Desc: "Category 2"},
	}
	for _, category := range categories {
		r.db.Where(Category{Name: category.Name}).Attrs(Category{Name: category.Name}).FirstOrCreate(&category)
	}
}

// Create 创建商品分类
func (r *Repository) Create(category *Category) error {
	tx := r.db.Create(category)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// GetByName 根据名称查询商品分类数据
func (r *Repository) GetByName(name string) []Category {
	var categories []Category
	err := r.db.Where("Name=?", name).Scan(&categories)
	if err != nil {
		return nil
	}
	return categories
}

// BulkCreate 批量创建商品分类数据
func (r *Repository) BulkCreate(categories []*Category) (int, error) {
	var count int64
	err := r.db.Create(&categories).Count(&count).Error
	return int(count), err
}

// GetAll 获取商品分类分页数据
func (r *Repository) GetAll(pageIndex, pageSize int) ([]Category, int) {
	var categories []Category
	var count int64
	r.db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&categories).Count(&count)
	return categories, int(count)
}
