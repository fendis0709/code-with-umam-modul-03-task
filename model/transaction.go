package model

import "time"

// Transaction represents a checkout transaction entity.
type Transaction struct {
	ID          int64               `json:"id"`
	UUID        string              `json:"uuid"`
	TotalAmount float64             `json:"total_amount"`
	PurchasedAt time.Time           `json:"purchased_at"`
	Details     []TransactionDetail `json:"details"`
}

// TransactionDetail represents the details of a transaction.
type TransactionDetail struct {
	ID            int64   `json:"id"`
	TransactionID int64   `json:"transaction_id"`
	ProductID     int64   `json:"product_id"`
	ProductUUID   string  `json:"product_uuid"`
	ProductName   string  `json:"product_name"`
	Price         float64 `json:"price"`
	Quantity      int64   `json:"quantity"`
	SubTotal      float64 `json:"sub_total"`
}
