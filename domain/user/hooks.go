package user

import (
	"gorm.io/gorm"
	"shopping_go/utils/hash"
)

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.Salt == "" {
		salt := hash.CreateSalt()
		password, err := hash.HashPassword(u.Password + salt)
		if err != nil {
			return nil
		}
		u.Password = password
		u.Salt = salt
	}
	return
}
