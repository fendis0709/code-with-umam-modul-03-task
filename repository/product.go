package repository

import (
	"context"
	"database/sql"
	"fendi/modul-02-task/model"
	"fmt"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAllProduct(ctx context.Context) ([]model.Product, error) {
	query := `
		SELECT 
			p.id, p.uuid, p.name, p.stock, p.price, p.category_id,
			c.id, c.uuid, c.name, c.description
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.id AND c.deleted_at IS NULL
		WHERE p.deleted_at IS NULL 
		ORDER BY p.id ASC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return []model.Product{}, nil
		}
		fmt.Println("repository.product.GetAllProduct() Query Error: ", err.Error())
		return nil, err
	}
	defer rows.Close()

	products := make([]model.Product, 0)
	for rows.Next() {
		var p model.Product
		var categoryID, categoryDBID sql.NullInt64
		var categoryUUID, categoryName sql.NullString
		var categoryDesc sql.NullString

		err := rows.Scan(
			&p.ID, &p.UUID, &p.Name, &p.Stock, &p.Price, &categoryID,
			&categoryDBID, &categoryUUID, &categoryName, &categoryDesc,
		)
		if err != nil {
			fmt.Println("repository.product.GetAllProduct() Scan Error: ", err.Error())
			return nil, err
		}

		if categoryID.Valid {
			categoryIDValue := categoryID.Int64
			p.CategoryID = &categoryIDValue
		}

		if categoryUUID.Valid && categoryName.Valid {
			category := &model.Category{
				ID:   categoryDBID.Int64,
				UUID: categoryUUID.String,
				Name: categoryName.String,
			}
			if categoryDesc.Valid {
				category.Description = &categoryDesc.String
			}
			p.Category = category
		}

		products = append(products, p)
	}

	return products, nil
}

func (r *ProductRepository) GetProductByUUID(ctx context.Context, uuid string) (model.Product, error) {
	query := `
		SELECT 
			p.id, p.uuid, p.name, p.stock, p.price, p.category_id,
			c.id, c.uuid, c.name, c.description
		FROM products p
		LEFT JOIN categories c ON p.category_id = c.id AND c.deleted_at IS NULL
		WHERE p.uuid = $1 AND p.deleted_at IS NULL
	`
	row := r.db.QueryRowContext(ctx, query, uuid)

	var p model.Product
	var categoryID, categoryDBID sql.NullInt64
	var categoryUUID, categoryName sql.NullString
	var categoryDesc sql.NullString

	err := row.Scan(
		&p.ID, &p.UUID, &p.Name, &p.Stock, &p.Price, &categoryID,
		&categoryDBID, &categoryUUID, &categoryName, &categoryDesc,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Product{}, nil
		}
		fmt.Println("repository.product.GetProductByUUID() Scan Error: ", err.Error())
		return model.Product{}, err
	}

	if categoryID.Valid {
		categoryIDValue := categoryID.Int64
		p.CategoryID = &categoryIDValue
	}

	if categoryUUID.Valid && categoryName.Valid {
		category := &model.Category{
			ID:   categoryDBID.Int64,
			UUID: categoryUUID.String,
			Name: categoryName.String,
		}
		if categoryDesc.Valid {
			category.Description = &categoryDesc.String
		}
		p.Category = category
	}

	return p, nil
}

func (r *ProductRepository) CreateProduct(ctx context.Context, p model.Product) error {
	query := "INSERT INTO products (uuid, name, stock, price, category_id) VALUES ($1, $2, $3, $4, $5)"
	_, err := r.db.ExecContext(ctx, query, p.UUID, p.Name, p.Stock, p.Price, p.CategoryID)
	if err != nil {
		fmt.Println("repository.product.CreateProduct() Exec Error: ", err.Error())
	}

	return err
}

func (r *ProductRepository) UpdateProduct(ctx context.Context, p model.Product) error {
	query := "UPDATE products SET name = $1, stock = $2, price = $3, category_id = $4, updated_at = NOW() WHERE uuid = $5"
	_, err := r.db.ExecContext(ctx, query, p.Name, p.Stock, p.Price, p.CategoryID, p.UUID)
	if err != nil {
		fmt.Println("repository.product.UpdateProduct() Exec Error: ", err.Error())
	}

	return err
}

func (r *ProductRepository) DeleteProduct(ctx context.Context, uuid string) error {
	query := "UPDATE products SET deleted_at = NOW() WHERE uuid = $1"
	_, err := r.db.ExecContext(ctx, query, uuid)
	if err != nil {
		fmt.Println("repository.product.DeleteProduct() Exec Error: ", err.Error())
	}

	return err
}
