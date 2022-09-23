package main

import (
	"fmt"
	"GoBasesPracticas"
)

func main() {

	fmt.Printf("Starts program...\n\n")

	GoBasesPracticas.GetBasicData()
	GoBasesPracticas.GetWeather()
	GoBasesPracticas.GetVars()
	GoBasesPracticas.GetTypes()

	GoBasesPracticas.SpellWord()
	GoBasesPracticas.GetMonthLabel(5)
	GoBasesPracticas.GetLoan()
	GoBasesPracticas.GetAgeEmployee()

	fmt.Printf("Finished program...\n\n")
}