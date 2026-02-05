package transport

// StatusResponse represents a standard status response.
type StatusResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

// ProductItemResponse represents a product item in the response.
type ProductItemResponse struct {
	ID       string                `json:"id"`
	Name     string                `json:"name"`
	Stock    *int64                `json:"stock"`
	Price    *float64              `json:"price"`
	Category *CategoryItemResponse `json:"category"`
}

// CategoryItemResponse represents a category item in the response.
type CategoryItemResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}
