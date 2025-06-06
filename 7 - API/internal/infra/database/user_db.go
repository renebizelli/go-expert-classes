package database

import (
	"renebizelli/api/internal/entities"

	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{DB: db}
}

func (u *User) Create(user *entities.User) error {
	return u.DB.Create(&user).Error
}

func (u *User) FindByEmail(email string) (*entities.User, error) {
	var user entities.User
	if e := u.DB.Where(" email = ?", email).First(&user).Error; e != nil {
		return nil, e
	}

	return &user, nil
}
