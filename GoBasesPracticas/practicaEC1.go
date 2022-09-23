package GoBasesPracticas

import "fmt"

func SpellWord() {

	fmt.Println("Ejercicio 1 - Letras de una palabra")

	var word string = "teletransportacion"

	var lengthWord int = len(word)

	fmt.Printf("La palabra %v tiene %v letras\n\n",word,lengthWord)

	for i, l := range word {
		fmt.Printf("%v - %c\n",i,l)
	}

	fmt.Print("\n\n")

}