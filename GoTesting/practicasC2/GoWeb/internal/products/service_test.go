package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/victorhcamacho/backpack-bcgow6-victor-martinez/GoTesting/practicasC2/GoWeb/internal/domain"
	"github.com/victorhcamacho/backpack-bcgow6-victor-martinez/GoTesting/practicasC2/GoWeb/test/mocks"
)

func TestServiceIntegrationUpdate(t *testing.T) {

	// update successful
	expectedResult := domain.Product{
		ID:            1,
		Name:          "Laptop",
		UnitPrice:     1000,
		StockQuantity: 50,
	}

	initialData := []domain.Product{
		{
			ID:            1,
			Name:          "Keyboard",
			UnitPrice:     150,
			StockQuantity: 5,
		},
	}

	store := mocks.MockStore{
		DataMock: initialData,
	}

	repo := NewRepository(&store)
	service := NewService(repo)

	result, err := service.Update(1, "Laptop", 1000, 50)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result, "el resultado es diferente a lo esperado")
	assert.True(t, store.ReadWasCalled, "no paso por la funcion read de store")

	// update failed not found
	expectedErr := errors.New("no fue posible encontrar el producto con id 5")

	result, err = service.Update(5, "Laptop", 1000, 50)

	assert.Empty(t, result)
	assert.EqualError(t, err, expectedErr.Error())

	// update failed write function
	expectedErr = errors.New("unexpected fail at write file function")

	store.ErrOnWrite = errors.New("unexpected fail at write file function")

	repo = NewRepository(&store)
	service = NewService(repo)

	result, err = service.Update(1, "Laptop", 1000, 50)

	assert.Empty(t, result)
	assert.NotNil(t, err)
	assert.EqualError(t, expectedErr, err.Error(), "los errores no coinciden")
}

func TestServiceIntegrationDelete(t *testing.T) {

	// delete successful
	expectedResult := []domain.Product{
		{
			ID:            2,
			Name:          "Laptop",
			UnitPrice:     1500,
			StockQuantity: 10,
		},
	}

	initialData := []domain.Product{
		{
			ID:            1,
			Name:          "Keyboard",
			UnitPrice:     150,
			StockQuantity: 5,
		},
		{
			ID:            2,
			Name:          "Laptop",
			UnitPrice:     1500,
			StockQuantity: 10,
		},
	}

	store := mocks.MockStore{
		DataMock: initialData,
	}

	repo := NewRepository(&store)
	service := NewService(repo)

	err := service.Delete(1)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, store.DataMock, "el resultado es diferente a lo esperado")

	// delete failed not found
	expectedErr := errors.New("no fue posible encontrar el producto con id 5")

	err = service.Delete(5)

	assert.NotNil(t, err)
	assert.EqualError(t, err, expectedErr.Error())

	// update failed write function
	expectedErr = errors.New("unexpected fail at write file function")

	store.ErrOnWrite = errors.New("unexpected fail at write file function")

	repo = NewRepository(&store)
	service = NewService(repo)

	err = service.Delete(1)

	assert.NotNil(t, err)
	assert.EqualError(t, expectedErr, err.Error(), "los errores no coinciden")
}
