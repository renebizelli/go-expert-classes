package database

import "renebizelli/api/internal/entities"

type UserInterface interface {
	Create(user *entities.User) error
	FindByEmail(email string) (*entities.User, error)
}

type ProductInterface interface {
	Create(product *entities.Product) error
	FindAll(page, limit int, sort string) ([]entities.Product, error)
	FindById(id string) (*entities.Product, error)
	Update(product *entities.Product) error
	Delete(id string) error
}
