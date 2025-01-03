package entities

import (
	"errors"
	pkg_entities "renebizelli/api/pkg/entities"
	"time"

	"github.com/google/uuid"
)

var (
	ErrorIDIsRequired        = errors.New("id is required")
	ErrorInvalidID           = errors.New("invalid id")
	ErrorNameIsRequired      = errors.New("name is required")
	ErrorPriceIsRequired     = errors.New("price is required")
	ErrorinvalidPrice        = errors.New("invalid price")
	ErrorCreatedAtIsRequired = errors.New("created at is required")
)

type Product struct {
	Id        pkg_entities.ID `json:"id"`
	Name      string          `json:"name"`
	Price     float64         `json:"price"`
	CreatedAt time.Time       `json:"createdAt"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		Id:        pkg_entities.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	if e := product.Validate(); e != nil {
		return nil, e
	}

	return product, nil
}

func (p *Product) Validate() error {

	if p.Id.String() == "" {
		return ErrorIDIsRequired
	}

	if _, e := uuid.Parse(p.Id.String()); e != nil {
		return ErrorInvalidID
	}

	if p.Name == "" {
		return ErrorNameIsRequired
	}

	if p.Price == 0 {
		return ErrorPriceIsRequired
	}

	if p.Price < 0 {
		return ErrorinvalidPrice
	}

	// if p.CreatedAt == time. {
	// 	return ErrorCreatedAtIsRequired
	// }

	return nil

}
