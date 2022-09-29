package ejercicios_tt

import "fmt"

type Matrix struct {
	Values   []float64
	Width    int
	Height   int
	MaxValue float64
}

func (m *Matrix) SetData(numbers ...float64) (err error) {

	m.Values = numbers

	if len(m.Values) != m.Width*m.Height {
		err = fmt.Errorf("la cantidad de valores no coincide con las dimensiones de la matriz")
		return
	}

	// set max value from array values
	for _, val := range numbers {
		if val > m.MaxValue {
			m.MaxValue = val
		}
	}

	return
}

func (m Matrix) PrintData() {
	for row := 0; row < m.Height; row++ {
		fmt.Printf("\t%.1f\n", m.Values[row*m.Width:row*m.Width+m.Width])
	}
}

func (m Matrix) Quadratic() (result bool) {

	if (m.Height == m.Width) && m.Height != 0 {
		result = true
	}

	return
}
