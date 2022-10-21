package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllProducts(t *testing.T) {

	initialData := []Product{
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

	store := MockStore{
		mockedData: initialData,
	}

	repo := NewRepository(&store)

	results, err := repo.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, initialData, results)
}

func TestUpdateNameProduct(t *testing.T) {

	expectedResult := Product{
		ID:            1,
		Name:          "Keyboard After Update",
		UnitPrice:     200,
		StockQuantity: 10,
	}

	initialData := []Product{
		{
			ID:            1,
			Name:          "Keyboard Before Update",
			UnitPrice:     100,
			StockQuantity: 10,
		},
	}

	store := MockStore{
		mockedData:    initialData,
		readWasCalled: false,
	}

	repo := NewRepository(&store)

	result, err := repo.UpdateNameAndPrice(1, "Keyboard After Update", 200)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result, "el resultado es diferente a lo esperado")
	assert.True(t, store.readWasCalled, "no paso por la funcion read de store")

}
