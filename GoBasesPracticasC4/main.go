package main

import (
	"errors"
	"fmt"
	"os"
)

const MIN_SALARY int = 150000

type CustomError struct {
	code    int
	message string
}

func (err *CustomError) Error() string {
	return fmt.Sprintf("error: %v (%d)", err.message, err.code)
}

func main() {

	salary := 150000

	err := CheckSalaryStruct(&salary)
	// err := CheckSalaryNew(&salary)
	// err := CheckSalaryFmt(&salary)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Debe pagar impuesto")
}

func CheckSalaryStruct(salary *int) error {

	if *salary < MIN_SALARY {
		return &CustomError{
			code:    500,
			message: "el salario ingresado no alcanza el mínimo imponible",
		}
	}

	return nil
}

func CheckSalaryNew(salary *int) error {

	if *salary < MIN_SALARY {
		return errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	}

	return nil
}

func CheckSalaryFmt(salary *int) error {

	if *salary < MIN_SALARY {
		return fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", *salary)
	}

	return nil
}
