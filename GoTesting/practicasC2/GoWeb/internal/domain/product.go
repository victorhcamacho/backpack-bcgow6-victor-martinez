package domain

type Product struct {
	ID            int     `json:"id"`
	Name          string  `json:"nombre"`
	UnitPrice     float64 `json:"precioUnitario"`
	StockQuantity int     `json:"cantidadExistencias"`
}
