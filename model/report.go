package model

type ReportData struct {
	TotalTransaction  int64             `json:"total_transaksi"`
	TotalRevenue      float64           `json:"total_revenue"`
	MostPurchasedItem MostPurchasedItem `json:"produk_terlaris"`
}

type MostPurchasedItem struct {
	ProductID   string `json:"id"`
	ProductName string `json:"nama"`
	Quantity    int64  `json:"qty_terjual"`
}
