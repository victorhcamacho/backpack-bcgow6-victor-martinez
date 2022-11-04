package mocks

import "github.com/victorhcamacho/backpack-bcgow6-victor-martinez/GoTesting/practicasC2/GoWeb/internal/domain"

type MockStore struct {
	DataMock      []domain.Product
	ErrOnWrite    error
	ErrOnRead     error
	ReadWasCalled bool
}

func (mock *MockStore) Read(data interface{}) error {

	if mock.ErrOnRead != nil {
		return mock.ErrOnRead
	}

	products := data.(*[]domain.Product)
	*products = mock.DataMock

	mock.ReadWasCalled = true

	return nil
}

func (mock *MockStore) Write(data interface{}) error {

	if mock.ErrOnWrite != nil {
		return mock.ErrOnWrite
	}

	products := data.([]domain.Product)
	mock.DataMock = products
	// mock.mockedData = append(mock.mockedData, products...)

	return nil
}
