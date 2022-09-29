package main

import (
	"backpack-bcgow6-victor-martinez/GoBasesPracticasC2/ejercicios_tt"
	"fmt"
	"os"
	"time"
)

func ejercicio1() {

	student1 := ejercicios_tt.Student{
		Name:     "Víctor Hugo",
		LastName: "Martínez",
		Dni:      637483939,
		Date:     time.Now().UTC().String(),
	}

	student2 := ejercicios_tt.Student{
		Name:     "Alexis Giovana",
		LastName: "Montes de Oca",
		Dni:      4930483923,
		Date:     time.Now().UTC().String(),
	}

	fmt.Println(student1.Detail())
	fmt.Println(student2.Detail())
}

func ejercicio2() {

	matrix := ejercicios_tt.Matrix{
		Width:  4,
		Height: 2,
	}

	err := matrix.SetData(1.4, 7.4, 8.7, 3.4, 2.8, 5.3, 67.0, 87.4)

	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Representacion de la Matriz:")
	matrix.PrintData()
	fmt.Printf("La matriz es cuadratica?: %v\n", matrix.Quadratic())
	fmt.Printf("El valor maximo de los valores de la matriz es: %.1f\n", matrix.MaxValue)
	fmt.Println("")
}

func ejercicio3() {
	var price float64 = 1000
	var category string = ejercicios_tt.BIG
	fmt.Printf("Producto %v con un precio de $%.2f\n", category, price)
	shop := ejercicios_tt.NewProduct(category, price)
	fmt.Printf("Precio total del producto: %.2f\n", shop.GetTotal())
	fmt.Println("")
}

func main() {

	ejercicio1()
	ejercicio2()
	ejercicio3()
}
