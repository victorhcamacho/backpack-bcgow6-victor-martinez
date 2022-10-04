package products

type Service interface {
	GetAll() ([]Product, error)
	GetByID(id int) (Product, error)
	Save(name string, price float64, stock int) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll() (result []Product, err error) {
	list, e := s.repository.GetAll()
	if e != nil {
		err = e
	}
	result = list
	return
}

func (s *service) GetByID(id int) (result Product, err error) {
	product, e := s.repository.GetByID(id)
	if e != nil {
		err = e
	}
	result = product
	return
}

func (s *service) Save(name string, price float64, stock int) (result Product, err error) {

	currentID, e := s.repository.GetLastID()
	if e != nil {
		err = e
		return
	}

	currentID++

	product, e := s.repository.Save(currentID, name, price, stock)
	if e != nil {
		err = e
	}

	result = product

	return
}
