package ejercicios_tt

type EUser struct {
	name     string
	lastName string
	email    string
	products []Product
}

func NewEUser(name string, lastName string, email string) EUser {
	return EUser{
		name:     name,
		lastName: lastName,
		email:    email,
	}
}

func AddProduct(euser *EUser, product *Product, quantity int) {
	product.quantity = quantity
	euser.products = append(euser.products, *product)
}

func ClearProducts(euser *EUser) {
	euser.products = []Product{}
}
