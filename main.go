package main

import (
	"database/sql"

	"github.com/mateusdeitos/golang-hexagonal/adapters/db"
	"github.com/mateusdeitos/golang-hexagonal/application"
)

func main() {
	Db, _ := sql.Open("sqlite3", "db.sqlite")
	productDb := db.NewProductDb(Db)
	productService := application.NewProductService(productDb)

	product, err := productService.Create("Product 1", 10)
	if err != nil {
		panic(err)
	}

	productService.Enable(product)
}
