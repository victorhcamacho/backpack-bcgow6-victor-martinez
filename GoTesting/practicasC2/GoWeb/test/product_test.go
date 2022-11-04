package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/victorhcamacho/backpack-bcgow6-victor-martinez/GoTesting/practicasC2/GoWeb/cmd/server/handler"
	"github.com/victorhcamacho/backpack-bcgow6-victor-martinez/GoTesting/practicasC2/GoWeb/internal/domain"
	"github.com/victorhcamacho/backpack-bcgow6-victor-martinez/GoTesting/practicasC2/GoWeb/internal/products"
	"github.com/victorhcamacho/backpack-bcgow6-victor-martinez/GoTesting/practicasC2/GoWeb/test/mocks"
)

func createServer(storageMock mocks.MockStore) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	_ = os.Setenv("TOKEN_AUTH", "12345QWERTY")

	repository := products.NewRepository(&storageMock)
	service := products.NewService(repository)
	productHandler := handler.NewProduct(service)

	server := gin.Default()

	routes := server.Group("/products")

	routes.GET("/", productHandler.GetAll())

	return server
}

func createTestRequest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {

	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "12345QWERTY")

	return req, httptest.NewRecorder()
}

func TestFunctionalGetAll(t *testing.T) {

	storage := mocks.MockStore{
		DataMock: []domain.Product{
			{
				ID:            1,
				Name:          "Laptop",
				UnitPrice:     1000,
				StockQuantity: 10,
			},
			{
				ID:            2,
				Name:          "Keyboard",
				UnitPrice:     499,
				StockQuantity: 100,
			},
		},
	}

	server := createServer(storage)

	request, response := createTestRequest(http.MethodGet, "/products/", "")

	server.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)

	var result []domain.Product
	// fmt.Printf("bytes: %v", response.Body.Bytes())
	errM := json.Unmarshal(response.Body.Bytes(), &result)

	assert.Nil(t, errM)
	// assert.Equal(t, storage.DataMock, result)
}
