package transport

// CategoryRequest represents the payload for creating or updating a category.
type CategoryRequest struct {
	UUID        *string `json:"uuid"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
}

// ProductRequest represents the payload for creating or updating a product.
type ProductRequest struct {
	UUID       *string  `json:"uuid"`
	Name       string   `json:"name"`
	Stock      *int64   `json:"stock"`
	Price      *float64 `json:"price"`
	CategoryID string   `json:"category_id"`
}

// CheckoutRequest represents the payload for checking out products.
type CheckoutRequest struct {
	Items []CheckoutItem `json:"items"`
}

// CheckoutItem represents an item in the checkout request.
type CheckoutItem struct {
	SKU      string `json:"sku"`
	Quantity int64  `json:"quantity"`
}
