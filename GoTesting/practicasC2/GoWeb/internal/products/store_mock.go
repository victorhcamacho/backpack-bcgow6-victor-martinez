package products

type MockStore struct {
	mockedData    []Product
	errOnWrite    error
	errOnRead     error
	readWasCalled bool
}

func (mock *MockStore) Read(data interface{}) error {

	if mock.errOnRead != nil {
		return mock.errOnRead
	}

	products := data.(*[]Product)
	*products = mock.mockedData

	mock.readWasCalled = true

	return nil
}

func (mock *MockStore) Write(data interface{}) error {

	if mock.errOnWrite != nil {
		return mock.errOnWrite
	}

	products := data.([]Product)
	mock.mockedData = products
	// mock.mockedData = append(mock.mockedData, products...)

	return nil
}
