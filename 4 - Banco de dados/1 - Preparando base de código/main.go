package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.NewString(),
		Name:  name,
		Price: price,
	}
}

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	panicIfError(err)

	defer db.Close()

	product := NewProduct("Guitarra", 15000.45)

	err = insertProduct(db, *product)
	panicIfError(err)

	product.Price = 100

	err = updateProduct(db, *product)

	panicIfError(err)

	productQuery, err := selectOneProduct(db, product.ID)
	panicIfError(err)
	fmt.Print(productQuery)

	err = deleteOneProduct(db, product.ID)
	fmt.Println(err)

	products, err := selectManyProduct(db, "Guitarra")
	panicIfError(err)

	for _, p := range products {
		fmt.Printf("%v custa R$ %.2f \n", p.Name, p.Price)
	}

	_, err = selectOneProduct(db, "x")
	fmt.Print(err)

}

func insertProduct(db *sql.DB, product Product) error {

	stmt, err := db.Prepare("INSERT INTO products(Id, Name, Price) VALUES (?, ?, ?)")
	panicIfError(err)

	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price)

	return err
}

func updateProduct(db *sql.DB, product Product) error {

	stmt, err := db.Prepare("UPDATE products SET Name = ?, Price = ? WHERE Id = ?")
	panicIfError(err)

	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.ID)

	return err
}

func selectOneProduct(db *sql.DB, id string) (*Product, error) {

	stmt, err := db.Prepare("SELECT Id, Name, Price FROM products WHERE Id = ?")
	panicIfError(err)

	defer stmt.Close()

	var product Product

	result := stmt.QueryRow(id)

	err = result.Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func selectManyProduct(db *sql.DB, name string) ([]Product, error) {

	stmt, err := db.Prepare("SELECT Id, Name, Price FROM products WHERE Name = ?")
	panicIfError(err)

	defer stmt.Close()

	var products []Product

	result, err := stmt.Query(name)

	if err != nil {
		return nil, err
	}

	defer result.Close()

	for result.Next() {

		var product Product

		err = result.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func deleteOneProduct(db *sql.DB, id string) error {

	stmt, err := db.Prepare("DELETE FROM products WHERE Id = ?")
	panicIfError(err)

	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

func panicIfError(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, "Ocorreu algum erro:")
		panic(err)
	}
}
