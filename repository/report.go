package repository

import (
	"database/sql"
	"fendi/modul-03-task/model"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) FetchReport(dateStart, dateEnd string) (model.ReportData, error) {
	var report model.ReportData

	query := `
		SELECT 
			COALESCE(SUM(t.total_amount), 0) AS total_revenue,
			COUNT(t.id) AS total_transaction,
			COALESCE(p.uuid::text, '') AS most_purchased_product_id,
			COALESCE(p.name, '') AS most_purchased_product_name,
			COALESCE(SUM(td.quantity), 0) AS most_purchased_quantity
		FROM 
			transactions t
		LEFT JOIN 
			transaction_details td ON t.id = td.transaction_id
		LEFT JOIN 
			products p ON td.product_id = p.id
		WHERE 
			t.purchased_at BETWEEN $1 AND $2
		GROUP BY 
			p.id
		ORDER BY 
			most_purchased_quantity DESC
		LIMIT 1;
	`

	row := r.db.QueryRow(query, dateStart, dateEnd)
	err := row.Scan(
		&report.TotalRevenue,
		&report.TotalTransaction,
		&report.MostPurchasedItem.ProductID,
		&report.MostPurchasedItem.ProductName,
		&report.MostPurchasedItem.Quantity,
	)
	if err != nil {
		return model.ReportData{}, err
	}

	return report, nil
}
