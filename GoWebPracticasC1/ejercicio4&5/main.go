package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type product struct {
	Id           int64
	Name         string
	UnitPrice    float64
	Stock        int64
	Color        string
	UnitCode     string
	Published    bool
	CreationDate string
}

func main() {

	server := gin.Default()

	productsRouter := server.Group("/products")
	{
		productsRouter.GET("/", getByQueryParams)
		productsRouter.GET("/:id/detail", getById)
	}

	server.Run(":8081")

}

func readJsonFile() (products []product) {

	content, err := os.ReadFile("../ejercicio1/products.json")

	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(content, &products); err != nil {
		panic(err)
	}

	return
}

func getByQueryParams(ctx *gin.Context) {

	idQuery := ctx.Query("id")
	nameQuery := ctx.Query("name")
	priceQuery := ctx.Query("price")
	stockQuery := ctx.Query("stock")
	colorQuery := ctx.Query("color")
	codeQuery := ctx.Query("code")
	publishedQuery := ctx.Query("published")
	creationQuery := ctx.Query("creation")

	var filteredProducts []product

	if idQuery != "" {

		filteredProducts = filterByIdProduct(idQuery)

		if len(filteredProducts) == 0 {
			ctx.String(http.StatusNotFound, "El producto con id %s no fue encontrado", idQuery)
			return
		}

	} else if nameQuery != "" {

		filteredProducts = checkProductValueByKey("name", nameQuery)

		if len(filteredProducts) == 0 {
			ctx.String(http.StatusNotFound, "El producto con nombre '%s' no fue encontrado", nameQuery)
			return
		}

	} else if priceQuery != "" {

		filteredProducts = filterByPriceProduct(priceQuery)

		if len(filteredProducts) == 0 {
			ctx.String(http.StatusNotFound, "El producto con precio de %s no fue encontrado", priceQuery)
			return
		}

	} else if stockQuery != "" {

		filteredProducts = filterByStockProduct(stockQuery)

		if len(filteredProducts) == 0 {
			ctx.String(http.StatusNotFound, "El producto con stock de %s no fue encontrado", stockQuery)
			return
		}

	} else if colorQuery != "" {

		filteredProducts = checkProductValueByKey("color", colorQuery)

		if len(filteredProducts) == 0 {
			ctx.String(http.StatusNotFound, "El producto de color '%s' no fue encontrado", colorQuery)
			return
		}

	} else if codeQuery != "" {

		filteredProducts = checkProductValueByKey("code", codeQuery)

		if len(filteredProducts) == 0 {
			ctx.String(http.StatusNotFound, "El producto con codigo '%s' no fue encontrado", codeQuery)
			return
		}

	} else if publishedQuery != "" {

		filteredProducts = filterByStatusProduct(publishedQuery)

		if len(filteredProducts) == 0 {
			ctx.String(http.StatusNotFound, "El producto con estatus '%s' no fue encontrado", publishedQuery)
			return
		}

	} else if creationQuery != "" {

		filteredProducts = checkProductValueByKey("date", creationQuery)

		if len(filteredProducts) == 0 {
			ctx.String(http.StatusNotFound, "El producto con fecha de creacion '%s' no fue encontrado", creationQuery)
			return
		}

	} else {
		allProducts := readJsonFile()
		ctx.JSON(http.StatusOK, allProducts)
		return
	}

	ctx.JSON(http.StatusOK, filteredProducts)
}

func getById(ctx *gin.Context) {

	idParam := ctx.Param("id")

	filtered := filterByIdProduct(idParam)

	if len(filtered) == 0 {
		ctx.String(http.StatusNotFound, "El producto con id %s no fue encontrado", idParam)
		return
	}

	ctx.JSON(http.StatusOK, filtered)
}

func filterByIdProduct(id string) (result []product) {

	idProduct, errParse := strconv.ParseInt(id, 10, 64)

	if errParse != nil {
		panic(errParse)
	}

	result = checkProductValueByKey("id", idProduct)

	return
}

func filterByPriceProduct(price string) (result []product) {

	priceProduct, errParse := strconv.ParseFloat(price, 64)

	if errParse != nil {
		panic(errParse)
	}

	result = checkProductValueByKey("price", priceProduct)

	return
}

func filterByStockProduct(stock string) (result []product) {

	stockProduct, errParse := strconv.ParseInt(stock, 10, 64)

	if errParse != nil {
		panic(errParse)
	}

	result = checkProductValueByKey("stock", stockProduct)

	return
}

func filterByStatusProduct(status string) (result []product) {

	statusProduct, errParse := strconv.ParseBool(status)

	if errParse != nil {
		panic(errParse)
	}

	result = checkProductValueByKey("published", statusProduct)

	return
}

func checkProductValueByKey(key string, value interface{}) (result []product) {

	var data interface{}

	products := readJsonFile()

	for _, p := range products {

		switch key {
		case "id":
			data = p.Id
		case "name":
			data = p.Name
		case "price":
			data = p.UnitPrice
		case "stock":
			data = p.Stock
		case "color":
			data = p.Color
		case "code":
			data = p.UnitCode
		case "published":
			data = p.Published
		case "date":
			data = p.CreationDate
		default:
			data = nil
		}

		if data == value {
			result = append(result, p)
		}
	}

	return
}
