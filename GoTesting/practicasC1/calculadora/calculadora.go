package calculadora

import "errors"

func Sumar(a int, b int) int {
	return a + b
}

func Restar(a int, b int) int {
	return a - b
}

func Dividir(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("el denominador no puede ser 0")
		// return 0, errors.New("no se puede divir entre 0")
	}
	return a / b, nil
}
