package service

import (
	"context"
	"fendi/modul-02-task/model"
	"fendi/modul-02-task/repository"
	"fendi/modul-02-task/transport"
	"fmt"

	"github.com/google/uuid"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAllCategory(ctx context.Context) ([]transport.CategoryItemResponse, error) {
	categories, err := s.repo.GetAllCategory(ctx)
	if err != nil {
		fmt.Print("s.repo.GetAllCategory() Error: ", err.Error())
		return nil, err
	}
	if len(categories) == 0 {
		return []transport.CategoryItemResponse{}, nil
	}

	categoriesResponse := transformCategory(categories)

	return categoriesResponse, nil
}

func (s *CategoryService) GetCategoryByUUID(ctx context.Context, uuid string) (transport.CategoryItemResponse, error) {
	category, err := s.repo.GetCategoryByUUID(ctx, uuid)
	if err != nil {
		fmt.Print("s.repo.GetCategoryByUUID() Error: ", err.Error())
		return transport.CategoryItemResponse{}, err
	}

	categoryResponse := transport.CategoryItemResponse{
		ID:          category.UUID,
		Name:        category.Name,
		Description: category.Description,
	}

	return categoryResponse, nil
}

func transformCategory(c []model.Category) []transport.CategoryItemResponse {
	var categoriesResponse []transport.CategoryItemResponse
	for _, category := range c {
		categoryResponse := transport.CategoryItemResponse{
			ID:          category.UUID,
			Name:        category.Name,
			Description: category.Description,
		}
		categoriesResponse = append(categoriesResponse, categoryResponse)
	}

	return categoriesResponse
}

func (s *CategoryService) CreateCategory(ctx context.Context, req transport.CategoryRequest) (transport.CategoryItemResponse, error) {
	randomUUID := uuid.New().String()

	newCategory := model.Category{
		UUID:        randomUUID,
		Name:        req.Name,
		Description: &req.Description,
	}

	err := s.repo.CreateCategory(ctx, newCategory)
	if err != nil {
		fmt.Print("s.repo.CreateCategory() Error: ", err.Error())
		return transport.CategoryItemResponse{}, err
	}

	categoryResponse := transport.CategoryItemResponse{
		ID:          newCategory.UUID,
		Name:        newCategory.Name,
		Description: newCategory.Description,
	}

	return categoryResponse, nil
}

func (s *CategoryService) UpdateCategory(ctx context.Context, id string, req transport.CategoryRequest) (transport.CategoryItemResponse, error) {
	newCategory := model.Category{
		UUID:        id,
		Name:        req.Name,
		Description: &req.Description,
	}

	err := s.repo.UpdateCategory(ctx, newCategory)
	if err != nil {
		fmt.Print("s.repo.UpdateCategory() Error: ", err.Error())
		return transport.CategoryItemResponse{}, err
	}

	categoryResponse := transport.CategoryItemResponse{
		ID:          newCategory.UUID,
		Name:        newCategory.Name,
		Description: newCategory.Description,
	}

	return categoryResponse, nil
}

func (s *CategoryService) DeleteCategory(ctx context.Context, id string) error {
	err := s.repo.DeleteCategory(ctx, id)
	if err != nil {
		fmt.Print("s.repo.DeleteCategory() Error: ", err.Error())
		return err
	}

	return nil
}
