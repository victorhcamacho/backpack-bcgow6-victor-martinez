package main

import (
	"fmt"
	"os"
)

const FileName string = ""

func readFile() (string, error) {

	content, err := os.ReadFile(FileName)

	if err != nil {
		return "", fmt.Errorf("el archivo indicado no fue encontrado o está dañado")
	}

	return string(content), nil
}

func main() {

	defer func() {
		fmt.Println("Ejecución finalizada")
	}()

	fmt.Println("Inicio de la ejecución")

	content, err := readFile()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Contenido del archivo:\n%v", content)
}
