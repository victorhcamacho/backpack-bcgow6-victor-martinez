package ejercicios_tt

import "fmt"

func GetLoan() {

	fmt.Println("Ejercicio 2 - Préstamo")

	var (
		nombre     string  = "Victor Martinez"
		edad       int     = 28
		sueldo     float32 = 50000.0
		antiguedad int     = 3
		esActivo   bool    = true
	)

	const EdadRequerida = 22
	const SueldoRequerido = 100000.0
	const AntiguedadRequerida = 1

	fmt.Printf("Criterios del prestamo de %v:\n", nombre)

	switch {
	case edad > EdadRequerida:
		fmt.Println("Es mayor de edad")
	case edad < EdadRequerida:
		fmt.Println("No es mayor de edad")
	}

	if !esActivo {
		fmt.Println("No se encuentra laborando actualmente")
	} else if esActivo && antiguedad > AntiguedadRequerida {

		fmt.Println("Se encuentra laborando y con mas de un año de antiguedad")

		if sueldo > SueldoRequerido {
			fmt.Println("No se cobraran intereses")
		} else {
			fmt.Println("Se cobraran los intereses generales")
		}

	} else {
		fmt.Println("No cuenta con la antiguedad necesaria")
	}

	fmt.Print("\n\n")

}
