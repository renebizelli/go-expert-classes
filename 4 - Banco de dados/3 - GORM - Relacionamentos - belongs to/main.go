package main

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryId int
	Category   Category
	gorm.Model
}

func main() {

	// parseTime = true -> nao da erro para inserir datetime
	dsn := "root:root@tcp(localhost:3306)/goexpert?parseTime=true&charset=utf8mb4"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	panicIfError(err)

	db.AutoMigrate(&Product{}, &Category{})

	//category := Category{Name: "Cordas"}

	// db.Create(&category)

	// db.Create(&Product{
	// 	Name:       "Guitarra Fender",
	// 	Price:      2100,
	// 	CategoryId: category.ID,
	// })

	fmt.Println("***************************************")

	products := []Product{}

	db.Preload("Category").Find(&products)
	for _, p := range products {
		fmt.Println(p)
	}

}

func panicIfError(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, "Ocorreu algum erro:")
		panic(err)
	}
}
