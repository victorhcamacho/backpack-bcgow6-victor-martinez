package products

import "fmt"

type Product struct {
	ID            int     `json:"id"`
	Name          string  `json:"nombre"`
	UnitPrice     float64 `json:"precioUnitario"`
	StockQuantity int     `json:"cantidadExistencias"`
}

var products []Product

/*
La interfaz sera el punto de acceso a
todos los metodos disponibles en el repository
*/
type Repository interface {
	GetAll() ([]Product, error)
	GetByID(id int) (Product, error)
	Save(id int, name string, price float64, stock int) (Product, error)
	GetLastID() (int, error)
}

/*
Por ahora no hay otra capa debajo de repository,
por esa razon se declara como una estructura vacia
*/
type repository struct{}

/*
Esta funcion nos ayuda a obtener algo parecido a
una instancia del repository, para poder utilizar sus metodos
*/
func NewRepository() Repository {
	return &repository{}
}

/*
Apartir de aca se construyen las implementaciones
de los metodos definidos por la interfaz
*/

func (r *repository) GetLastID() (result int, err error) {
	result = len(products)
	return
}

func (r *repository) GetAll() (result []Product, err error) {
	result = products
	return
}

func (r *repository) GetByID(id int) (result Product, err error) {
	for _, p := range products {
		if p.ID == id {
			result = p
			break
		}
	}

	if result.ID == 0 {
		err = fmt.Errorf("no fue posible encontrar el producto con id %d", id)
	}

	return
}

func (r *repository) Save(id int, name string, price float64, stock int) (result Product, err error) {
	result = Product{id, name, price, stock}
	products = append(products, result)
	return
}
