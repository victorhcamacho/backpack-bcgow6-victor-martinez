package main

import (
	"fmt"

	"github.com/victorhcamacho/backpack-bcgow6-victor-martinez/GoBasesPracticasC1/ejercicios_tm"
	"github.com/victorhcamacho/backpack-bcgow6-victor-martinez/GoBasesPracticasC1/ejercicios_tt"
)

func ejecutarEjerciciosMañana() {
	ejercicios_tm.GetBasicData()
	ejercicios_tm.GetWeather()
	ejercicios_tm.GetVars()
	ejercicios_tm.GetTypes()
}

func ejecutarEjerciciosTarde() {
	ejercicios_tt.SpellWord()
	ejercicios_tt.GetLoan()
	ejercicios_tt.GetMonthLabel(5)
	ejercicios_tt.GetAgeEmployee()
}

func main() {

	fmt.Println("Inicia Practicas Go Bases Clase 1")

	ejecutarEjerciciosMañana()
	ejecutarEjerciciosTarde()

	fmt.Println("Termina Practicas Go Bases Clase 1")
}
