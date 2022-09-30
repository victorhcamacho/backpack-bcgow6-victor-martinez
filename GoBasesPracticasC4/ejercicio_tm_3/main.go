package main

import (
	"fmt"
	"math/rand"
	"os"
)

type Customer struct {
	Documents   uint
	Name        string
	LastName    string
	PhoneNumber string
	Address     string
}

const FileName string = ""

func readFile() (data []byte, err error) {

	data, err = os.ReadFile(FileName)

	if err != nil {
		return nil, fmt.Errorf("error: el archivo indicado no fue encontrado o está dañado")
	}

	return data, nil
}

func getId(min int, max int) (id int) {
	if min > 0 {
		id = rand.Intn(max-min) + min
	}
	return
}

func checkCustomerExist(customer Customer) (bool, error) {

	data, err := readFile()

	if err != nil {
		return false, err
	}

	fmt.Println(string(data))

	return true, nil
}

func checkCustomerStruct(customer Customer) (err error) {

	if customer.Name == "" || customer.LastName == "" {
		err = fmt.Errorf("el nombre y apellidos del cliente no pueden ir vacio")
	} else if customer.PhoneNumber == "" {
		err = fmt.Errorf("el numero de telefono del cliente no puede ir vacio")
	} else if customer.Address == "" {
		err = fmt.Errorf("la direccion del cliente no puede ir vacia")
	}

	return
}

func main() {

	fmt.Println("Inicio de la ejecucion")

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Fin de la ejecucion")
			fmt.Println("Se detectaron varios errores en tiempo de ejecucion")
			fmt.Println("No han quedado archivos abiertos")
		}
	}()

	id := getId(10, 100)

	if id == 0 {
		panic("no se obtuvo un legajo valido")
	}

	fmt.Printf("Legajo generado: %d\n", id)

	customer := Customer{
		Name:        "Víctor Hugo",
		LastName:    "Martínez",
		PhoneNumber: "5521280252",
		Address:     "Estafetas 21, Col. Postal, Del. Benito Juarez, CDMX",
	}

	fmt.Println("Inicia la revision del cliente...")

	exist, err1 := checkCustomerExist(customer)

	fmt.Println("Termina la revision del cliente...")

	if err1 != nil {
		panic(err1)
	}

	if exist {
		fmt.Println("El cliente ya existe en el archivo")
	}

	fmt.Println("Inicia validacion de los datos del cliente...")

	err2 := checkCustomerStruct(customer)

	fmt.Println("Termina validacion de los datos del cliente...")

	if err2 != nil {
		panic(err2)
	}

	fmt.Println("El clientes es valido")

}
