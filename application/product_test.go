package application_test

import (
	"testing"

	"github.com/mateusdeitos/golang-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{
		Name: "Hello",
		Price: 10,
		Status: application.DISABLED,
	}

	err := product.Enable()

	require.Nil(t, err)
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{
		Name: "Hello",
		Price: 0,
		Status: application.ENABLED,
	}

	err := product.Disable()

	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be 0 to disable the product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{
		ID: uuid.NewV4().String(),
		Name: "Hello",
		Price: 10,
		Status: application.DISABLED,
	}

	isValid, err := product.IsValid()

	require.Nil(t, err)
	require.True(t, isValid)
	
	product.Status = "INVALID"
	isValid, err = product.IsValid()
	require.False(t, isValid)
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	isValid, err = product.IsValid()
	require.True(t, isValid)
	require.Nil(t, err)

	product.Price = -10
	isValid, err = product.IsValid()
	require.False(t, isValid)
	require.Equal(t, "the price must be greater or equal than 0", err.Error())
}