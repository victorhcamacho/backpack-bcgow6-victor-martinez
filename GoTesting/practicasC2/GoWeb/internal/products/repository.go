package products

import (
	"fmt"

	"github.com/victorhcamacho/backpack-bcgow6-victor-martinez/GoTesting/practicasC2/GoWeb/internal/domain"
	"github.com/victorhcamacho/backpack-bcgow6-victor-martinez/GoTesting/practicasC2/GoWeb/pkg/store"
)

/*
La interfaz sera el punto de acceso a
todos los metodos disponibles en el repository
*/
type Repository interface {
	Delete(id int) error
	GetLastID() (int, error)
	GetAll() ([]domain.Product, error)
	GetByID(id int) (domain.Product, error)
	Save(id int, name string, price float64, stock int) (domain.Product, error)
	Update(id int, name string, price float64, stock int) (domain.Product, error)
	UpdateNameAndPrice(id int, name string, price float64) (domain.Product, error)
}

/*
Por ahora no hay otra capa debajo de repository,
por esa razon se declara como una estructura vacia
*/
type repository struct {
	db store.Store
}

/*
Esta funcion nos ayuda a obtener algo parecido a
una instancia del repository, para poder utilizar sus metodos
*/
func NewRepository(db store.Store) Repository {
	return &repository{db}
}

/*
Apartir de aca se construyen las implementaciones
de los metodos definidos por la interfaz
*/

func (r *repository) GetLastID() (result int, err error) {

	var products []domain.Product
	if err = r.db.Read(&products); err != nil {
		return
	}

	result = products[len(products)-1].ID

	return
}

func (r *repository) GetAll() ([]domain.Product, error) {

	var result []domain.Product
	if err := r.db.Read(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) GetByID(id int) (result domain.Product, err error) {

	var products []domain.Product
	if err = r.db.Read(&products); err != nil {
		return
	}

	result, _ = findEntityByID(id, products)

	if result.ID == 0 {
		err = fmt.Errorf("no fue posible encontrar el producto con id %d", id)
	}

	return
}

func (r *repository) Save(id int, name string, price float64, stock int) (result domain.Product, err error) {

	var products []domain.Product
	if err = r.db.Read(&products); err != nil {
		return
	}

	result = domain.Product{ID: id, Name: name, UnitPrice: price, StockQuantity: stock}
	products = append(products, result)

	if err := r.db.Write(products); err != nil {
		return domain.Product{}, err
	}

	return
}

func (r *repository) Update(id int, name string, price float64, stock int) (result domain.Product, err error) {

	var products []domain.Product
	if err = r.db.Read(&products); err != nil {
		return
	}

	p, i := findEntityByID(id, products)

	if p.ID == 0 {
		err = fmt.Errorf("no fue posible encontrar el producto con id %d", id)
		return
	}

	result = domain.Product{ID: id, Name: name, UnitPrice: price, StockQuantity: stock}

	products[i] = result

	if err := r.db.Write(products); err != nil {
		return domain.Product{}, err
	}

	return
}

func (r *repository) UpdateNameAndPrice(id int, name string, price float64) (result domain.Product, err error) {

	var products []domain.Product
	if err = r.db.Read(&products); err != nil {
		return
	}

	p, i := findEntityByID(id, products)

	if p.ID == 0 {
		err = fmt.Errorf("no fue posible encontrar el producto con id %d", id)
		return
	}

	p.Name = name
	p.UnitPrice = price

	products[i] = p

	result = p

	if err := r.db.Write(products); err != nil {
		return domain.Product{}, err
	}

	return
}

func (r *repository) Delete(id int) (err error) {

	var products []domain.Product
	if err = r.db.Read(&products); err != nil {
		return
	}

	p, i := findEntityByID(id, products)

	if p.ID == 0 {
		err = fmt.Errorf("no fue posible encontrar el producto con id %d", id)
		return
	}
	products = append(products[:i], products[i+1:]...)

	if err := r.db.Write(products); err != nil {
		return err
	}

	return
}

func findEntityByID(id int, products []domain.Product) (p domain.Product, index int) {

	for i, pr := range products {
		if pr.ID == id {
			p = pr
			index = i
			break
		}
	}

	return
}
