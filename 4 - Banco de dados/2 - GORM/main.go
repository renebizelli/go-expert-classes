package main

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {

	// parseTime = true -> nao da erro para inserir datetime
	dsn := "root:root@tcp(localhost:3306)/goexpert?parseTime=true&charset=utf8mb4"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	panicIfError(err)

	db.AutoMigrate(&Product{})

	// db.Create(&Product{
	// 	Name:  "Guitarra Fender",
	// 	Price: 2100,
	// })

	// products := []Product{{Name: "Violão", Price: 2000}, {Name: "Baixo", Price: 5000}}

	// db.Create(&products)

	// var first Product
	// db.First(&first, 2)
	// fmt.Println(first)

	// var query Product
	// db.First(&query, "Name = ?", "Baixo")
	// fmt.Println(query)

	// products = []Product{}

	// db.Where("Name LIKE ?", "%FENDER%").Find(&products)
	// for i, p := range products {

	// 	p.Name = fmt.Sprintf("%v %d", p.Name, i)
	// 	fmt.Println(p)
	// 	db.Save(&p)
	// }

	// fmt.Println("***************************************")

	// products = []Product{}

	// db.Limit(10).Offset(2).Find(&products, "Name LIKE ?", "%FENDER%")
	// for _, p := range products {
	// 	fmt.Println(p)
	// }

	// fmt.Println("***************************************")

	// products = []Product{}

	// db.Find(&products, "Name LIKE ?", "%vio%")
	// for _, p := range products {
	// 	db.Delete(&p)
	// }

	// fmt.Println("***************************************")

	// products = []Product{}

	// db.Limit(10).Offset(2).Find(&products, "Name LIKE ?", "%vio%")
	// for _, p := range products {
	// 	fmt.Println(p)
	// }

	db.Create(&Product{
		Name:  "Guitarra Fender",
		Price: 2100})

	db.Create(&Product{
		Name:  "Violão Fender",
		Price: 2100})

	products := []Product{}

	db.Find(&products)
	for i, p := range products {
		p.Name = fmt.Sprintf("%v - %d", p.Name, i)
		db.Save(&p)
	}

	fmt.Println("***************************************")

	products = []Product{}

	db.Find(&products)
	for _, p := range products {
		fmt.Println(p)
	}

	fmt.Println("***************************************")

	products = []Product{}

	db.Find(&products, "Name LIKE ?", "%vio%")
	for _, p := range products {
		db.Delete(&p)
	}

	fmt.Println("***************************************")

	products = []Product{}

	db.Find(&products, "Name LIKE ?", "%vio%")
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
