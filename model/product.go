package model

// Product represents a product entity.
type Product struct {
	ID       int64     `json:"id"`
	UUID     string    `json:"uuid"`
	SKU      string    `json:"sku"`
	Name     string    `json:"name"`
	Stock    *int64    `json:"stock"`
	Price    *float64  `json:"price"`
	Category *Category `json:"category"`
}
