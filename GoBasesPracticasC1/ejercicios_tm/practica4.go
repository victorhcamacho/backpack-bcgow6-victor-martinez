package ejercicios_tm

import "fmt"

func GetTypes() {

	fmt.Println("Ejercicio 4 - Tipos de datos")

	result := "var apellido string = \"Gomez\"\nvar edad int = 35\nbooleanValue := false\nvar sueldo float32 = 45857.90\nvar nombre string = \"Julián\""
	fmt.Printf("Correct types:\n%v\n\n", result)
}
