package main

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryId   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductId int
}

func main() {

	// parseTime = true -> nao da erro para inserir datetime
	dsn := "root:root@tcp(localhost:3306)/goexpert?parseTime=true&charset=utf8mb4"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	panicIfError(err)

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	product := Product{
		Name:       "Guitarra Fender III",
		Price:      2100,
		CategoryId: 1,
	}

	db.Create(&product)

	db.Create(&SerialNumber{
		Number:    "123",
		ProductId: product.ID,
	})

	categories := []Category{}

	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	panicIfError(err)

	for _, c := range categories {
		fmt.Println(c.Name)
		for _, p := range c.Products {
			fmt.Println("-", p.Name, p.SerialNumber.Number)
		}
	}

}

func panicIfError(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, "Ocorreu algum erro:")
		panic(err)
	}
}
