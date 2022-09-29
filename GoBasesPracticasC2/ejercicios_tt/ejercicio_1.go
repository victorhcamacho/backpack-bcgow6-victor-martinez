package ejercicios_tt

import "fmt"

type Student struct {
	Name     string
	LastName string
	Dni      uint
	Date     string
}

func (s Student) Detail() string {
	return fmt.Sprintf("\nDNI:\t%v\nAlumno:\t%v %v\nFecha de ingreso:\t%v\n", s.Dni, s.Name, s.LastName, s.Date)
}
