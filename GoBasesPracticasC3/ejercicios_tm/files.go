package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	ID        int
	Name      string
	UnitPrice float64
	Quantity  int
}

func CreateFile() {

	product_1 := Product{
		ID:        1100,
		Name:      "Laptop",
		UnitPrice: 2555.9,
		Quantity:  2,
	}

	product_2 := Product{
		ID:        1111,
		Name:      "Smartphone",
		UnitPrice: 5449.9,
		Quantity:  3,
	}

	line_1 := fmt.Sprintf("%d;%s;%.2f;%d\n", product_1.ID, product_1.Name, product_1.UnitPrice, product_1.Quantity)
	line_2 := fmt.Sprintf("%d;%s;%.2f;%d\n", product_2.ID, product_2.Name, product_2.UnitPrice, product_2.Quantity)

	err := os.WriteFile("./order.csv", []byte(line_1+line_2), 0644)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("El archivo se genero correctamente")
}

func ReadFile() {

	data, err := os.ReadFile("./order.csv")

	if err != nil {
		panic(err.Error())
	}

	PrintTable(string(data))
}

func PrintTable(content string) {

	fmt.Println("ID\t\t\tNombre\t\tPrecio\t\tCantidad")

	lines := strings.Fields(content)

	var total float64
	for _, line := range lines {
		product := strings.Split(line, ";")
		fmt.Printf("%s\t\t\t%s\t\t%s\t\t%s\n", product[0], product[1], product[2], product[3])
		r, _ := strconv.ParseFloat(product[2], 64)
		total += r
	}

	fmt.Printf("\t\t\t\tTotal de la compra: %.2f\n", total)
}

func main() {
	// CreateFile()
	ReadFile()
}
