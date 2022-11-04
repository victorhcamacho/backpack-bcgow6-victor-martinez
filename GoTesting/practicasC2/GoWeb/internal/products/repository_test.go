package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/victorhcamacho/backpack-bcgow6-victor-martinez/GoTesting/practicasC2/GoWeb/internal/domain"
	"github.com/victorhcamacho/backpack-bcgow6-victor-martinez/GoTesting/practicasC2/GoWeb/test/mocks"
)

func TestGetLastProductID(t *testing.T) {

	// reading successful
	initialData := []domain.Product{
		{
			ID:            1,
			Name:          "Keyboard",
			UnitPrice:     100,
			StockQuantity: 10,
		},
		{
			ID:            2,
			Name:          "Screen 27 inches",
			UnitPrice:     500,
			StockQuantity: 5,
		},
	}

	store := mocks.MockStore{
		DataMock: initialData,
	}

	repo := NewRepository(&store)

	id, err := repo.GetLastID()

	assert.Nil(t, err)
	assert.Equal(t, 2, id)

	// reading failed
	expectedErr := errors.New("unexpected fail at read file function")

	store.ErrOnRead = errors.New("unexpected fail at read file function")

	repo = NewRepository(&store)

	id, err = repo.GetLastID()

	assert.Empty(t, id)
	assert.NotNil(t, err)
	assert.EqualError(t, expectedErr, err.Error(), "los errores no coinciden")
}

func TestGetProductByID(t *testing.T) {

	// reading successful
	expectedResult := domain.Product{
		ID:            2,
		Name:          "Screen 27 inches",
		UnitPrice:     500,
		StockQuantity: 5,
	}

	initialData := []domain.Product{
		{
			ID:            1,
			Name:          "Keyboard",
			UnitPrice:     100,
			StockQuantity: 10,
		},
		{
			ID:            2,
			Name:          "Screen 27 inches",
			UnitPrice:     500,
			StockQuantity: 5,
		},
	}

	store := mocks.MockStore{
		DataMock: initialData,
	}

	repo := NewRepository(&store)

	result, err := repo.GetByID(2)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)

	// reading failed
	expectedErr := errors.New("no fue posible encontrar el producto con id 5")

	result, err = repo.GetByID(5)

	assert.Empty(t, result)
	assert.NotNil(t, err)
	assert.EqualError(t, expectedErr, err.Error(), "los errores no coinciden")
}

func TestGetAllProducts(t *testing.T) {

	// reading successful
	initialData := []domain.Product{
		{
			ID:            1,
			Name:          "Keyboard",
			UnitPrice:     100,
			StockQuantity: 10,
		},
		{
			ID:            2,
			Name:          "Screen 27 inches",
			UnitPrice:     500,
			StockQuantity: 5,
		},
	}

	store := mocks.MockStore{
		DataMock: initialData,
	}

	repo := NewRepository(&store)

	results, err := repo.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, initialData, results)

	// reading failed
	expectedErr := errors.New("unexpected fail at read file function")

	store.ErrOnRead = errors.New("unexpected fail at read file function")

	repo = NewRepository(&store)

	results, err = repo.GetAll()

	assert.Empty(t, results)
	assert.NotNil(t, err)
	assert.EqualError(t, expectedErr, err.Error(), "los errores no coinciden")
}

func TestSaveProduct(t *testing.T) {

	// create successful
	expectedResult := domain.Product{
		ID:            3,
		Name:          "Laptop",
		UnitPrice:     5000,
		StockQuantity: 100,
	}

	initialData := []domain.Product{
		{
			ID:            1,
			Name:          "Keyboard",
			UnitPrice:     100,
			StockQuantity: 10,
		},
		{
			ID:            2,
			Name:          "Screen 27 inches",
			UnitPrice:     500,
			StockQuantity: 5,
		},
	}

	store := mocks.MockStore{
		DataMock: initialData,
	}

	repo := NewRepository(&store)

	result, err := repo.Save(3, "Laptop", 5000, 100)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result, "el resultado es diferente a lo esperado")

	// create falied
	expectedErr := errors.New("unexpected fail at write file function")

	store.ErrOnWrite = errors.New("unexpected fail at write file function")

	repo = NewRepository(&store)

	result, err = repo.Save(3, "Laptop", 5000, 100)

	assert.Empty(t, result)
	assert.NotNil(t, err)
	assert.EqualError(t, expectedErr, err.Error(), "los errores no coinciden")
}

func TestUpdateNameProduct(t *testing.T) {

	// updating successful
	expectedResult := domain.Product{
		ID:            1,
		Name:          "Keyboard After Update",
		UnitPrice:     200,
		StockQuantity: 10,
	}

	initialData := []domain.Product{
		{
			ID:            1,
			Name:          "Keyboard Before Update",
			UnitPrice:     100,
			StockQuantity: 10,
		},
	}

	store := mocks.MockStore{
		DataMock: initialData,
	}

	repo := NewRepository(&store)

	result, err := repo.UpdateNameAndPrice(1, "Keyboard After Update", 200)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result, "el resultado es diferente a lo esperado")
	assert.True(t, store.ReadWasCalled, "no paso por la funcion read de store")

	// updating failed
	expectedErr := errors.New("no fue posible encontrar el producto con id 5")

	result, err = repo.UpdateNameAndPrice(5, "Keyboard After Update", 200)

	assert.Empty(t, result)
	assert.NotNil(t, err)
	assert.EqualError(t, expectedErr, err.Error(), "los errores no coinciden")
}
