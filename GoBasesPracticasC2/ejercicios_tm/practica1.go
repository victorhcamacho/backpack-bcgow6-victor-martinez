package ejercicios_tm

func ImpuestoSalario(salario float64) (float64, error) {

	var resultado float64

	if salario <= 0 {

	} else if salario > 50000 && salario < 150000 {
		resultado = salario * 0.17
	} else if salario > 150000 {
		resultado = salario * 0.27
	}

	return resultado, nil
}
