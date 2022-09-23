package GoBasesPracticas

import "fmt"

func GetAgeEmployee() {

	fmt.Println("Ejercicio 4 - Qué edad tiene...")

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Printf("Empleados: %v\n",employees)

	fmt.Printf("Edad de Benjamin: %v\n",employees["Benjamin"])

	count := 0

	for _,v := range employees {
		if v > 21 {
			count++
		}
	}

	fmt.Printf("%v empleados mayores de 21 años\n",count)

	employees["Federico"] = 25

	fmt.Printf("Se agrego a Federico: %v\n",employees)

	delete(employees,"Pedro")

	fmt.Printf("Se elimino a Pedro: %v\n\n",employees)
}