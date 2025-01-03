package database

import (
	"renebizelli/api/internal/entities"

	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{DB: db}
}

func (p *Product) Create(product *entities.Product) error {
	return p.DB.Create(product).Error
}

func (p *Product) FindById(id string) (*entities.Product, error) {
	var product entities.Product

	e := p.DB.First(&product, " id = ? ", id).Error

	if e != nil {
		return nil, e
	}

	return &product, nil
}

func (p *Product) Update(product *entities.Product) error {

	if _, e := p.FindById(product.Id.String()); e != nil {
		return e
	}

	e := p.DB.Save(product).Error

	return e
}

func (p *Product) Delete(id string) error {

	product, e := p.FindById(id)

	if e != nil {
		return e
	}

	e = p.DB.Delete(product).Error

	return e
}

func (p *Product) FindAll(page, limit int, sort string) ([]entities.Product, error) {

	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	var products []entities.Product
	var e error
	if page != 0 && limit != 0 {
		e = p.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&products).Error
	} else {
		e = p.DB.Order("created_at " + sort).Find(&products).Error
	}

	if e != nil {
		return nil, e
	}

	return products, nil
}
