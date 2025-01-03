package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {

	name := "Guitarra"
	price := 99.0

	product, e := NewProduct(name, price)

	assert.Nil(t, e)
	assert.NotNil(t, product)
	assert.Equal(t, name, product.Name)
	assert.Equal(t, price, product.Price)
	assert.NotEmpty(t, product.Id)
}

func TestProductWhenNameIsRequired(t *testing.T) {

	product, e := NewProduct("", 99.7)

	assert.Nil(t, product)
	assert.Equal(t, ErrorNameIsRequired, e)
}

func TestProductWhenPriceIsRequired(t *testing.T) {

	product, e := NewProduct("X", 0.0)

	assert.Nil(t, product)
	assert.Equal(t, ErrorPriceIsRequired, e)
}

func TestProductWhenPriceInvalid(t *testing.T) {

	product, e := NewProduct("X", -20.8)

	assert.Nil(t, product)
	assert.Equal(t, ErrorinvalidPrice, e)
}

func TestProductValidate(t *testing.T) {

	product, e := NewProduct("X", 1.6)

	assert.Nil(t, e)
	assert.NotNil(t, product)
	assert.Nil(t, product.Validate())
}
