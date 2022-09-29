package ejercicios_tt

import "fmt"

type Product struct {
	name      string
	unitPrice float64
	quantity  int
}

type Service struct {
	name      string
	unitPrice float64
	minutes   int
}

type Maintenance struct {
	name      string
	unitPrice float64
}

func NewProduct(name string, price float64) Product {
	return Product{name, price, 0}
}

func NewProductE(name string, price float64, quantity int) Product {
	return Product{name, price, quantity}
}

func NewService(name string, price float64, minutes int) Service {
	return Service{name, price, minutes}
}

func NewMaintenance(name string, price float64) Maintenance {
	return Maintenance{name, price}
}

func GetProductsCost(products *[]Product, channel chan float64) {
	var result float64
	for _, p := range *products {
		result += p.unitPrice * float64(p.quantity)
	}
	fmt.Printf("Costo total de los productos: $%.2f\n", result)
	channel <- result
}

func GetServicesCost(services *[]Service, channel chan float64) {
	var result float64
	for _, s := range *services {

		mid := s.minutes / 30

		if mid < 1 {
			mid = 1
		}

		result += s.unitPrice * float64(mid)
	}
	fmt.Printf("Costo total de los servicios: $%.2f\n", result)
	channel <- result
}

func GetMaintenancesCost(maintenances *[]Maintenance, channel chan float64) {
	var result float64
	for _, m := range *maintenances {
		result += m.unitPrice
	}
	fmt.Printf("Costo total de los mantenimientos: $%.2f\n", result)
	channel <- result
}
