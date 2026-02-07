package repository

import (
	"context"
	"database/sql"
	"fendi/modul-03-task/helper"
	"fendi/modul-03-task/model"
	"fendi/modul-03-task/transport"
	"fmt"
	"time"
)

type CheckoutRepository struct {
	db *sql.DB
}

func NewCheckoutRepository(db *sql.DB) *CheckoutRepository {
	return &CheckoutRepository{db: db}
}

func (r *CheckoutRepository) CreateCheckoutTransaction(ctx context.Context, req transport.CheckoutRequest) (*model.Transaction, error) {
	tx, err := r.db.Begin()
	if err != nil {
		fmt.Print("Failed to begin transaction: ", err)
		return nil, err
	}
	defer tx.Rollback()

	var productUUIDs []string
	var products []model.Product
	for _, item := range req.Items {
		productUUIDs = append(productUUIDs, item.ID)
	}

	if len(productUUIDs) == 0 {
		return nil, fmt.Errorf("no items provided for checkout")
	}

	// Build placeholders for IN clause (PostgreSQL style: $1, $2, $3...)
	placeholders := ""
	args := make([]interface{}, len(productUUIDs))
	for i, uuid := range productUUIDs {
		if i > 0 {
			placeholders += ", "
		}
		placeholders += fmt.Sprintf("$%d", i+1)
		args[i] = uuid
	}

	var query string
	query = fmt.Sprintf("SELECT id, uuid, name, stock, price FROM products WHERE uuid IN (%s)", placeholders)
	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no products found for the given UUIDs")
		}
		fmt.Print("Failed to query products: ", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p model.Product
		if err := rows.Scan(&p.ID, &p.UUID, &p.Name, &p.Stock, &p.Price); err != nil {
			fmt.Print("Failed to scan product row: ", err)
			return nil, err
		}
		products = append(products, p)
	}

	var totalAmount float64
	var transactionDetails []model.TransactionDetail

	for _, item := range req.Items {
		var product model.Product
		for _, p := range products {
			if p.UUID == item.ID {
				product = p
				break
			}
		}

		if product.UUID == "" {
			fmt.Printf("Product with UUID %s not found\n", item.ID)
			continue
		}

		itemQty := item.Quantity
		if product.Stock != nil {
			if *product.Stock <= 0 {
				fmt.Printf("Product with UUID %s is out of stock\n", item.ID)
				continue
			}

			if itemQty > *product.Stock {
				itemQty = *product.Stock
			}
			newStock := *product.Stock - itemQty

			query = "UPDATE products SET stock = $1 WHERE uuid = $2"
			_, err := tx.ExecContext(ctx, query, newStock, product.UUID)
			if err != nil {
				fmt.Print("Failed to update product stock: ", err)
				return nil, err
			}
		}

		var price, subTotal float64
		price = 0
		if product.Price != nil {
			price = *product.Price
		}
		subTotal = float64(itemQty) * price
		totalAmount += subTotal

		transactionDetails = append(transactionDetails, model.TransactionDetail{
			ProductID:   product.ID,
			ProductUUID: product.UUID,
			Price:       price,
			Quantity:    itemQty,
			SubTotal:    subTotal,
		})
	}

	var transactionID int64
	var transactionUUID string
	var currentTime time.Time
	transactionUUID = helper.GenerateUUID()
	currentTime = time.Now()
	query = "INSERT INTO transactions (uuid, total_amount, transaction_at) VALUES ($1, $2, $3) RETURNING id"
	err = tx.QueryRowContext(ctx, query, transactionUUID, totalAmount, currentTime).Scan(&transactionID)
	if err != nil {
		fmt.Print("Failed to insert transaction: ", err)
		return nil, err
	}

	for i, detail := range transactionDetails {
		var trxDetailID int64
		query = "INSERT INTO transaction_details (transaction_id, product_id, price, quantity, sub_total) VALUES ($1, $2, $3, $4, $5) RETURNING id"
		err := tx.QueryRowContext(ctx, query, transactionID, detail.ProductID, detail.Price, detail.Quantity, detail.SubTotal).Scan(&trxDetailID)
		if err != nil {
			fmt.Print("Failed to insert transaction detail: ", err)
			return nil, err
		}

		transactionDetails[i].ID = trxDetailID
	}

	err = tx.Commit()
	if err != nil {
		fmt.Print("Failed to commit transaction: ", err)
		return nil, err
	}

	var transaction model.Transaction
	transaction = model.Transaction{
		ID:            transactionID,
		UUID:          transactionUUID,
		TotalAmount:   totalAmount,
		TransactionAt: currentTime,
		Details:       transactionDetails,
	}

	return &transaction, nil
}
