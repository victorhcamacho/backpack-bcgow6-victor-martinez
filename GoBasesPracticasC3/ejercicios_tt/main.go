package main

import (
	"fmt"
)

func main() {
	user := User{}
	socialNetwork(&user)

	product := NewProduct("Laptop", 1499.99)
	euser := NewEUser("Víctor Hugo", "Martínez", "vmartinez@mercadolibre.com.mx")
	ecommerce(&euser, &product)

	var products []Product
	products = append(products, NewProductE("Laptop", 1499.99, 2))
	products = append(products, NewProductE("Mouse", 349.99, 2))
	products = append(products, NewProductE("Keyboard", 559.99, 2))

	var services []Service
	services = append(services, NewService("Agua potable", 35.55, 90))
	services = append(services, NewService("Luz electrica", 70.25, 60))
	services = append(services, NewService("Gas natural", 45.2, 20))

	var maintenances []Maintenance
	maintenances = append(maintenances, NewMaintenance("Pintura", 49.0))
	maintenances = append(maintenances, NewMaintenance("Plomeria", 80.5))
	maintenances = append(maintenances, NewMaintenance("Electricista", 249.99))

	price(&products, &services, &maintenances)
}

func socialNetwork(usr *User) {

	usr.SetAge(28)
	usr.SetFullName("Víctor Hugo", "Martínez")
	usr.SetEmail("vmartinez@mercadolibre.com.mx")
	usr.SetPassword("1234567890")

	fmt.Printf("Usuario nuevo: %+v\n", usr)

	usr.SetAge(30)
	usr.SetFullName("Maria Fernanda", "Perez")
	usr.SetEmail("mfer@mercadolibre.com.mx")
	usr.SetPassword("ksl_00_*!")

	fmt.Printf("Usuario modificado: %+v\n", usr)
}

func ecommerce(user *EUser, product *Product) {

	fmt.Printf("Ecommerce user inicio: %+v\n", *user)

	AddProduct(user, product, 5)

	fmt.Printf("Ecommerce user con productos: %+v\n", *user)

	ClearProducts(user)

	fmt.Printf("Ecommerce user sin productos: %+v\n", *user)

}

func price(products *[]Product, services *[]Service, maintenances *[]Maintenance) {

	var total float64

	channel := make(chan float64)

	go GetProductsCost(products, channel)
	go GetServicesCost(services, channel)
	go GetMaintenancesCost(maintenances, channel)

	for i := 0; i < 3; i++ {
		total += <-channel
	}

	fmt.Printf("Costo total final: $%.4f\n", total)

}
