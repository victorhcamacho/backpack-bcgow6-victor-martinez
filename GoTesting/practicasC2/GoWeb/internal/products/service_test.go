package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceIntegrationUpdate(t *testing.T) {

	expectedResult := Product{
		ID:            1,
		Name:          "Laptop",
		UnitPrice:     1000,
		StockQuantity: 50,
	}

	initialData := []Product{
		{
			ID:            1,
			Name:          "Keyboard",
			UnitPrice:     150,
			StockQuantity: 5,
		},
	}

	store := MockStore{
		mockedData:    initialData,
		readWasCalled: false,
	}

	repo := NewRepository(&store)
	service := NewService(repo)

	result, err := service.Update(1, "Laptop", 1000, 50)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result, "el resultado es diferente a lo esperado")
	assert.True(t, store.readWasCalled, "no paso por la funcion read de store")
}

func TestServiceIntegrationUpdateFail(t *testing.T) {

	expectedErr := errors.New("no fue posible encontrar el producto con id 5")

	initialData := []Product{
		{
			ID:            1,
			Name:          "Keyboard",
			UnitPrice:     150,
			StockQuantity: 5,
		},
	}

	store := MockStore{
		mockedData: initialData,
	}

	repo := NewRepository(&store)
	service := NewService(repo)

	result, err := service.Update(5, "Laptop", 1000, 50)

	assert.Empty(t, result)
	assert.EqualError(t, err, expectedErr.Error())
}

func TestServiceIntegrationDelete(t *testing.T) {

	expectedResult := []Product{
		{
			ID:            2,
			Name:          "Laptop",
			UnitPrice:     1500,
			StockQuantity: 10,
		},
	}

	initialData := []Product{
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

	store := MockStore{
		mockedData: initialData,
	}

	repo := NewRepository(&store)
	service := NewService(repo)

	err := service.Delete(1)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, store.mockedData, "el resultado es diferente a lo esperado")
}

func TestServiceIntegrationDeleteFail(t *testing.T) {

	expectedErr := errors.New("no fue posible encontrar el producto con id 5")

	initialData := []Product{
		{
			ID:            1,
			Name:          "Keyboard",
			UnitPrice:     150,
			StockQuantity: 5,
		},
	}

	store := MockStore{
		mockedData: initialData,
	}

	repo := NewRepository(&store)
	service := NewService(repo)

	err := service.Delete(5)

	assert.NotNil(t, err)
	assert.EqualError(t, err, expectedErr.Error())
}
