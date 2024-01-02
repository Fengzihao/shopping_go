package cart

import "gorm.io/gorm"

func (i *Item) AfterUpdate(db *gorm.DB) (err error) {
	if i.Count <= 0 {
		return db.Unscoped().Delete(&i).Error
	}
	return
}
