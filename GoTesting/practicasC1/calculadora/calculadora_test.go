package calculadora

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumarOK(t *testing.T) {

	primerNumero := 5
	segundoNumero := 5
	resultadoEsperado := 10

	resultado := Sumar(primerNumero, segundoNumero)

	if resultado != resultadoEsperado {
		t.Errorf("La funcion Sumar retorna: %d, pero el resultado esperado es: %d", resultado, resultadoEsperado)
	}
}

func TestRestarOK(t *testing.T) {

	primerNumero := 15
	segundoNumero := 5
	resultadoEsperado := 10

	resultado := Restar(primerNumero, segundoNumero)

	if resultado != resultadoEsperado {
		t.Errorf("La funcion Restar retorna: %d, pero el resultado esperado es: %d", resultado, resultadoEsperado)
	}
}

func TestDividirOK(t *testing.T) {

	primerNumero := 15
	segundoNumero := 3
	resultadoEsperado := 5

	resultado, err := Dividir(primerNumero, segundoNumero)

	assert.Nil(t, err)
	assert.Equal(t, resultadoEsperado, resultado, "La funcion Dividir retorna: %d, pero el resultado esperado es: %d", resultado, resultadoEsperado)
}

func TestDividirBad(t *testing.T) {

	primerNumero := 15
	segundoNumero := 0
	errorEsperado := errors.New("el denominador no puede ser 0")

	_, err := Dividir(primerNumero, segundoNumero)

	assert.NotNil(t, err, "el error deberia ser diferente de nil")
	assert.ErrorContains(t, err, errorEsperado.Error(), "el error no contiene: %s", errorEsperado.Error())
}
