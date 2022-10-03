package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type product struct {
	Id           int
	Name         string
	UnitPrice    float64
	Stock        int
	Color        string
	UnitCode     string
	Published    bool
	CreationDate string
}

func ReadJsonFile() []product {

	content, err := os.ReadFile("../ejercicio1/products.json")

	if err != nil {
		panic(err)
	}

	var products []product

	if err := json.Unmarshal(content, &products); err != nil {
		panic(err)
	}

	// fmt.Printf("Contenido del archivo json: \n %+v \n", products)

	return products
}

func GetAll(ctx *gin.Context) {
	products := ReadJsonFile()
	ctx.JSON(http.StatusOK, gin.H{
		"message":  "success",
		"products": products,
	})
}

func main() {
	router := gin.Default()

	router.GET("/products/", GetAll)

	router.Run(":8081")
}
