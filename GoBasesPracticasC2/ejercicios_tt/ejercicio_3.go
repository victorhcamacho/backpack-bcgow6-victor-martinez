package ejercicios_tt

const (
	SMALL  string = "PEQUEÑO"
	MEDIUM string = "MEDIANO"
	BIG    string = "GRANDE"
)

type Product interface {
	GetTotal() float64
}

type Shop struct {
	price    float64
	category string
}

func NewProduct(category string, price float64) Product {
	return Shop{price: price, category: category}
}

func (s Shop) GetTotal() float64 {
	switch s.category {
	case SMALL:
		return s.price
	case MEDIUM:
		mantencion := (s.price * 3) / 100
		return s.price + mantencion
	case BIG:
		mantencion := (s.price * 6) / 100
		return s.price + mantencion + 2500
	default:
		return 0
	}
}
