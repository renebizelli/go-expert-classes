package database

import (
	"fmt"
	"math/rand"
	"renebizelli/api/internal/entities"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {

	db, e := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if e != nil {
		t.Error(e)
	}

	db.AutoMigrate(&entities.Product{})

	name := "Guitarra"
	price := 99.98

	product, _ := entities.NewProduct(name, price)
	productDb := NewProduct(db)

	eProductDb := productDb.Create(product)
	assert.Nil(t, eProductDb)

	var productFound entities.Product

	err := db.First(&productFound, " id = ?", product.Id).Error
	assert.Nil(t, err)

	assert.NotEmpty(t, productFound.Id)
	assert.Equal(t, name, productFound.Name)
	assert.Equal(t, price, productFound.Price)
	assert.NotEmpty(t, productFound.CreatedAt)
}

func TestFindAll(t *testing.T) {

	db, e := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if e != nil {
		t.Error(e)
	}

	db.AutoMigrate(&entities.Product{})

	productDb := NewProduct(db)

	for i := 0; i < 20; i++ {

		name := fmt.Sprintf("Guitarra %d", i)

		product, _ := entities.NewProduct(name, rand.Float64()*100)

		if e := productDb.Create(product); e != nil {
			assert.NoError(t, e)
		}
	}

	products, e := productDb.FindAll(1, 8, "")
	assert.NoError(t, e)

	assert.Len(t, products, 8)

	products, e = productDb.FindAll(3, 8, "")
	assert.NoError(t, e)

	assert.Len(t, products, 4)
}

func TestFindById(t *testing.T) {

	db, e := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if e != nil {
		t.Error(e)
	}

	db.AutoMigrate(&entities.Product{})

	var products []entities.Product

	productDb := NewProduct(db)

	for i := 0; i < 20; i++ {

		name := fmt.Sprintf("Guitarra %d", i)

		product, _ := entities.NewProduct(name, rand.Float64()*100)

		products = append(products, *product)

		if e := productDb.Create(product); e != nil {
			assert.NoError(t, e)
		}
	}

	for _, p := range products {
		f, e := productDb.FindById(p.Id.String())
		assert.NoError(t, e)

		assert.Equal(t, p.Name, f.Name)
		assert.Equal(t, p.Price, f.Price)
		assert.Equal(t, p.CreatedAt.Local(), f.CreatedAt.Local())
	}
}

func TestUpdateProduct(t *testing.T) {

	db, e := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if e != nil {
		t.Error(e)
	}

	db.AutoMigrate(&entities.Product{})

	productDb := NewProduct(db)

	name := "Guitarra"
	price := 99.98

	product, _ := entities.NewProduct(name, price)
	productDb.Create(product)

	product.Name = "Violão"
	product.Price = 199.98

	e = productDb.Update(product)
	assert.Nil(t, e)

	var productFound entities.Product

	err := db.First(&productFound, " id = ?", product.Id).Error
	assert.Nil(t, err)

	assert.NotEmpty(t, productFound.Id)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
	assert.NotEmpty(t, productFound.CreatedAt)
}

func TestDeleteProduct(t *testing.T) {

	db, e := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if e != nil {
		t.Error(e)
	}

	db.AutoMigrate(&entities.Product{})

	productDb := NewProduct(db)

	name := "Guitarra"
	price := 99.98

	product, _ := entities.NewProduct(name, price)
	productDb.Create(product)

	e = productDb.Delete(product.Id.String())
	assert.Nil(t, e)

	var productFound entities.Product

	err := db.First(&productFound, " id = ?", product.Id).Error
	assert.NotNil(t, err)
}
