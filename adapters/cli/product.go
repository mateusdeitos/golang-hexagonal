package cli

import (
	"fmt"

	"github.com/mateusdeitos/golang-hexagonal/application"
)

func Run(
	service application.ProductServiceInterface,
	action string,
	productId string,
	productName string,
	price float64,
) (string, error) {
	var result = ""
	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return "", err
		}

		result = fmt.Sprintf("Product ID %s with name %s has been created with price %f", product.GetID(), product.GetName(), product.GetPrice())
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return "", err
		}

		err = service.Enable(product)
		if err != nil {
			return "", err
		}

		result = fmt.Sprintf("Product ID %s with name %s has been enabled", product.GetID(), product.GetName())

	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return "", err
		}

		err = service.Disable(product)
		if err != nil {
			return "", err
		}

		result = fmt.Sprintf("Product ID %s with name %s has been disabled", product.GetID(), product.GetName())

	default:
		product, err := service.Get(productId)
		if err != nil {
			return "", err
		}

		result = fmt.Sprintf("Product ID %s with name %s has price %f", product.GetID(), product.GetName(), product.GetPrice())

	}

	return result, nil
}
