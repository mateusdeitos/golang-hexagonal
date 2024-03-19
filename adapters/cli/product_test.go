package cli_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mateusdeitos/golang-hexagonal/adapters/cli"
	mock_application "github.com/mateusdeitos/golang-hexagonal/application/mocks"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product 1"
	productPrice := 10.0
	productStatus := "enabled"
	productId := "abc"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	serviceMock := mock_application.NewMockProductServiceInterface(ctrl)
	serviceMock.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Enable(productMock).Return(nil).AnyTimes()
	serviceMock.EXPECT().Disable(productMock).Return(nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID %s with name %s has been created with price %f", productId, productName, productPrice)
	result, err := cli.Run(serviceMock, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	result, err = cli.Run(serviceMock, "enable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("Product ID %s with name %s has been enabled", productId, productName), result)

	result, err = cli.Run(serviceMock, "disable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("Product ID %s with name %s has been disabled", productId, productName), result)

	result, err = cli.Run(serviceMock, "get", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("Product ID %s with name %s has price %f", productId, productName, productPrice), result)
}
