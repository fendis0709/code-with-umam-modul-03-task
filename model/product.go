package model

// Product represents a product entity.
type Product struct {
	ID         int64     `json:"id"`
	UUID       string    `json:"uuid"`
	Name       string    `json:"name"`
	Stock      *int64    `json:"stock"`
	Price      *float64  `json:"price"`
	CategoryID *int64    `json:"category_id"`
	Category   *Category `json:"category"`
}
