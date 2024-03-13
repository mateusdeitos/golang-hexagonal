package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/mateusdeitos/golang-hexagonal/adapters/db"
	"github.com/mateusdeitos/golang-hexagonal/application"
	"github.com/stretchr/testify/require"
)

var DB *sql.DB

func setUp() {
	DB, _ = sql.Open("sqlite3", ":memory:")
	createTableProducts(DB)
	createProduct(DB)
}

func createTableProducts(db *sql.DB) {
	query := "CREATE TABLE products (id STRING, name STRING, price FLOAT, status STRING)"

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}
}

func createProduct(db *sql.DB) {
	insert := "INSERT INTO products (id, name, price, status) VALUES (?, ?, ?, ?)"

	stmt, err := db.Prepare(insert)

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec("1", "Product 1", 10.0, "enabled")
	if err != nil {
		log.Fatal(err)
	}
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer DB.Close()

	productDb := db.NewProductDb(DB)

	product, err := productDb.Get("1")
	require.Nil(t, err)

	require.Equal(t, "1", product.GetID())
	require.Equal(t, "Product 1", product.GetName())
	require.Equal(t, 10.0, product.GetPrice())
	require.Equal(t, "enabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer DB.Close()

	productDb := db.NewProductDb(DB)

	product := application.NewProduct()
	product.Name = "Product 2"
	product.Price = 25

	result, err := productDb.Save(product)
	require.Nil(t, err)

	require.Equal(t, "Product 2", result.GetName())
	require.Equal(t, 25.0, result.GetPrice())
	require.Equal(t, "disabled", result.GetStatus())

	product.Status = "enabled"
	result, err = productDb.Save(product)
	require.Nil(t, err)

	require.Equal(t, "Product 2", result.GetName())
	require.Equal(t, 25.0, result.GetPrice())
	require.Equal(t, "enabled", result.GetStatus())

}
