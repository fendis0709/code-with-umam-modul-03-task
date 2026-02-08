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

// CheckoutResponse represents the response for a checkout operation.
type CheckoutResponse struct {
	ID          string                 `json:"id"`
	Date        string                 `json:"date"`
	TotalAmount float64                `json:"total_amount"`
	Items       []CheckoutItemResponse `json:"items"`
}

// CheckoutItemResponse represents an item in the checkout response.
type CheckoutItemResponse struct {
	ProductID   string  `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    int64   `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	TotalPrice  float64 `json:"total_price"`
}

// ReportResponse represents the daily report response.
type ReportResponse struct {
	TotalRevenue      float64                    `json:"total_revenue"`
	TotalTransaction  int64                      `json:"total_transaksi"`
	MostPurchasedItem *MostPurchasedItemResponse `json:"produk_terlaris"`
}

// MostPurchasedItemResponse represents the most purchased item in the report.
type MostPurchasedItemResponse struct {
	ProductID   string `json:"id"`
	ProductName string `json:"nama"`
	Quantity    int64  `json:"qty_terjual"`
}
