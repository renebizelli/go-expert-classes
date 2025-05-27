// go:build wireinject
// +build wireinject
package main

func NewProductUseCase(db *sql.DB) *product.ProductUseCase {

	wire.Build(
		product.NewProductRepository,
		product.NewProductUseCase
	)

	return &product.ProductUseCase

}