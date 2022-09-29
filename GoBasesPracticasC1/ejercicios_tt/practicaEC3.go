package ejercicios_tt

import "fmt"

func GetMonthLabel(number int) {

	fmt.Println("Ejercicio 3 - A qué mes corresponde")

	var label string

	switch number {
	case 1:
		label = "Enero"
	case 2:
		label = "Febrero"
	case 3:
		label = "Marzo"
	case 4:
		label = "Abril"
	case 5:
		label = "Mayo"
	case 6:
		label = "Junio"
	case 7:
		label = "Julio"
	case 8:
		label = "Agosto"
	case 9:
		label = "Septiembre"
	case 10:
		label = "Octubre"
	case 11:
		label = "Noviembre"
	case 12:
		label = "Diciembre"
	default:
		label = "ninguno de los meses que contempla un año"
	}

	fmt.Printf("El mes %v corresponde a %v \n\n", number, label)

}
