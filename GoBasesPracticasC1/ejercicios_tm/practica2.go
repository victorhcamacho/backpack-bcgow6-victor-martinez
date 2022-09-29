package ejercicios_tm

import "fmt"

func GetWeather() {

	fmt.Println("Ejercicio 2 - Clima")

	var (
		temp  int     = 24
		hum   string  = "66%"
		press float32 = 1015.2
	)

	fmt.Printf("Weather: Temp %v° - Humidity %v - Pressure %v mb\n\n", temp, hum, press)
}
