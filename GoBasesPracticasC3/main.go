package main

import (
	"backpack-bcgow6-victor-martinez/GoBasesPracticasC3/ejercicios_tt"
	"fmt"
)

func main() {
	user := ejercicios_tt.User{}
	socialNetwork(&user)

	product := ejercicios_tt.NewProduct("Laptop", 1499.99)
	euser := ejercicios_tt.NewEUser("Víctor Hugo", "Martínez", "vmartinez@mercadolibre.com.mx")
	ecommerce(&euser, &product)

	var products []ejercicios_tt.Product
	products = append(products, ejercicios_tt.NewProductE("Laptop", 1499.99, 2))
	products = append(products, ejercicios_tt.NewProductE("Mouse", 349.99, 2))
	products = append(products, ejercicios_tt.NewProductE("Keyboard", 559.99, 2))

	var services []ejercicios_tt.Service
	services = append(services, ejercicios_tt.NewService("Agua potable", 35.55, 90))
	services = append(services, ejercicios_tt.NewService("Luz electrica", 70.25, 60))
	services = append(services, ejercicios_tt.NewService("Gas natural", 45.2, 20))

	var maintenances []ejercicios_tt.Maintenance
	maintenances = append(maintenances, ejercicios_tt.NewMaintenance("Pintura", 49.0))
	maintenances = append(maintenances, ejercicios_tt.NewMaintenance("Plomeria", 80.5))
	maintenances = append(maintenances, ejercicios_tt.NewMaintenance("Electricista", 249.99))

	price(&products, &services, &maintenances)
}

func socialNetwork(usr *ejercicios_tt.User) {

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

func ecommerce(user *ejercicios_tt.EUser, product *ejercicios_tt.Product) {

	fmt.Printf("Ecommerce user inicio: %+v\n", *user)

	ejercicios_tt.AddProduct(user, product, 5)

	fmt.Printf("Ecommerce user con productos: %+v\n", *user)

	ejercicios_tt.ClearProducts(user)

	fmt.Printf("Ecommerce user sin productos: %+v\n", *user)

}

func price(products *[]ejercicios_tt.Product, services *[]ejercicios_tt.Service, maintenances *[]ejercicios_tt.Maintenance) {

	var total float64

	channel := make(chan float64)

	go ejercicios_tt.GetProductsCost(products, channel)
	go ejercicios_tt.GetServicesCost(services, channel)
	go ejercicios_tt.GetMaintenancesCost(maintenances, channel)

	for i := 0; i < 3; i++ {
		total += <-channel
	}

	fmt.Printf("Costo total final: $%.4f\n", total)

}
