package application_test

import (
	"testing"

	"github.com/mateusdeitos/golang-hexagonal/application"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 0

	err := product.Enable()

	require.Nil(t, err)
}