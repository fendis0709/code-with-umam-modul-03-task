package repository

import (
	"context"
	"database/sql"
	"fendi/modul-02-task/model"
	"fmt"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAllCategory(ctx context.Context) ([]model.Category, error) {
	query := "SELECT id, uuid, name, description FROM categories WHERE deleted_at IS NULL ORDER BY id ASC"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return []model.Category{}, nil
		}
		fmt.Println("repository.category.GetAllCategory() Query Error: ", err.Error())
		return nil, err
	}
	defer rows.Close()

	categories := make([]model.Category, 0)
	for rows.Next() {
		var c model.Category
		err := rows.Scan(&c.ID, &c.UUID, &c.Name, &c.Description)
		if err != nil {
			fmt.Println("repository.category.GetAllCategory() Scan Error: ", err.Error())
			return nil, err
		}
		categories = append(categories, c)
	}

	return categories, nil
}

func (r *CategoryRepository) GetCategoryByUUID(ctx context.Context, uuid string) (model.Category, error) {
	query := "SELECT id, uuid, name, description FROM categories WHERE uuid = $1 AND deleted_at IS NULL"
	row := r.db.QueryRowContext(ctx, query, uuid)

	var c model.Category
	err := row.Scan(&c.ID, &c.UUID, &c.Name, &c.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Category{}, nil
		}
		fmt.Println("repository.category.GetCategoryByUUID() Scan Error: ", err.Error())
		return model.Category{}, err
	}

	return c, nil
}

func (r *CategoryRepository) CreateCategory(ctx context.Context, c model.Category) error {
	query := "INSERT INTO categories (uuid, name, description) VALUES ($1, $2, $3)"
	_, err := r.db.ExecContext(ctx, query, c.UUID, c.Name, c.Description)
	if err != nil {
		fmt.Println("repository.category.CreateCategory() Exec Error: ", err.Error())
	}

	return err
}

func (r *CategoryRepository) UpdateCategory(ctx context.Context, c model.Category) error {
	query := "UPDATE categories SET name = $1, description = $2, updated_at = NOW() WHERE uuid = $3"
	_, err := r.db.ExecContext(ctx, query, c.Name, c.Description, c.UUID)
	if err != nil {
		fmt.Println("repository.category.UpdateCategory() Exec Error: ", err.Error())
	}

	return err
}

func (r *CategoryRepository) DeleteCategory(ctx context.Context, uuid string) error {
	query := "UPDATE categories SET deleted_at = NOW() WHERE uuid = $1"
	_, err := r.db.ExecContext(ctx, query, uuid)
	if err != nil {
		fmt.Println("repository.category.DeleteCategory() Exec Error: ", err.Error())
	}

	return err
}
