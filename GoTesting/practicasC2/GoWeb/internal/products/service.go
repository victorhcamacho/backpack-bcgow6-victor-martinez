package products

import "github.com/victorhcamacho/backpack-bcgow6-victor-martinez/GoTesting/practicasC2/GoWeb/internal/domain"

type Service interface {
	Delete(id int) error
	GetAll() ([]domain.Product, error)
	GetByID(id int) (domain.Product, error)
	Save(name string, price float64, stock int) (domain.Product, error)
	Update(id int, name string, price float64, stock int) (domain.Product, error)
	UpdateNameAndPrice(id int, name string, price float64) (domain.Product, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll() ([]domain.Product, error) {
	return s.repository.GetAll()
}

func (s *service) GetByID(id int) (domain.Product, error) {
	return s.repository.GetByID(id)
}

func (s *service) Save(name string, price float64, stock int) (domain.Product, error) {

	currentID, e := s.repository.GetLastID()

	if e != nil {
		return domain.Product{}, e
	}

	currentID++

	return s.repository.Save(currentID, name, price, stock)
}

func (s *service) Update(id int, name string, price float64, stock int) (domain.Product, error) {
	return s.repository.Update(id, name, price, stock)
}

func (s *service) UpdateNameAndPrice(id int, name string, price float64) (domain.Product, error) {
	return s.repository.UpdateNameAndPrice(id, name, price)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
