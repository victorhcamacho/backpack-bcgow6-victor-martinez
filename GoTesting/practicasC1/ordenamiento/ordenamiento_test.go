package ordenamiento

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenamientoOK(t *testing.T) {

	lista := []int{4, 8, 1, 7, 3, 6, 2, 5, 10, 9}

	resultadoEsperado := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	resultado, err := OrdenarSlice(lista)

	assert.Nil(t, err)
	assert.Equal(t, resultadoEsperado, resultado, "La funcion OrdenarSlice retorna: %s, pero el resultado esperado es: %s", resultado, resultadoEsperado)
}

func TestOrdenamientoBad(t *testing.T) {

	lista := []int{}

	errorEsperado := errors.New("no es posible ordenar una lista vacia")

	_, err := OrdenarSlice(lista)

	assert.NotNil(t, err, "el error deberia ser diferente de nil")
	assert.ErrorContains(t, err, errorEsperado.Error(), "el error no contiene: %s", errorEsperado.Error())
}
