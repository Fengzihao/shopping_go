package user

import (
	"gorm.io/gorm"
	"log"
)

// Repository 结构体
type Repository struct {
	db *gorm.DB
}

// NewUserRepository 实例化
func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Migration 生成表
func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&User{})
	if err != nil {
		log.Print(err)
	}
}

// Create 创建用户数据
func (r *Repository) Create(u *User) error {
	tx := r.db.Create(u)
	return tx.Error
}

// GetByName 根据名称查找用户
func (r *Repository) GetByName(name string) (User, error) {
	var user User
	tx := r.db.Where(" Username=? and IsDeleted = ?", name, 0).Scan(&user)
	if tx.Error != nil {
		return User{}, tx.Error
	}
	return user, nil
}

// InsertSampleData 默认添加测试数据
func (r *Repository) InsertSampleData() {
	user := NewUser("admin", "admin", "admin")
	user.IsAdmin = true
	r.db.Where(User{Username: user.Username}).Attrs(
		User{
			Username: user.Username,
			Password: user.Password,
		}).FirstOrCreate(&user)
	user = NewUser("user", "user", "user")
	r.db.Where(User{Username: user.Username}).Attrs(
		User{
			Username: user.Username,
			Password: user.Password,
		}).FirstOrCreate(&user)
}

// Update 更新数据
func (r *Repository) Update(user *User) error {
	return r.db.Save(&user).Error
}
