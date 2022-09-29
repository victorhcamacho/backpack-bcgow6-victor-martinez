package ejercicios_tm

import "fmt"

func GetVars() {

	fmt.Println("Ejercicio 3 - Declaración de variables")

	result := "var nombre string\n var apellido string\n var edad int\n apellido := \"Martinez\"\n var licencia_de_conducir boolean = true\n var estatura_de_la_persona int\n cantidadDeHijos := 2\n"

	fmt.Printf("Correct variables:\n%v\n\n", result)
}
