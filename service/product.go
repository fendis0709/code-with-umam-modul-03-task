package service

import (
	"context"
	"fendi/modul-02-task/model"
	"fendi/modul-02-task/repository"
	"fendi/modul-02-task/transport"
	"fmt"

	"github.com/google/uuid"
)

type ProductService struct {
	repo         *repository.ProductRepository
	categoryRepo *repository.CategoryRepository
}

func NewProductService(repo *repository.ProductRepository, categoryRepo *repository.CategoryRepository) *ProductService {
	return &ProductService{
		repo:         repo,
		categoryRepo: categoryRepo,
	}
}

func (s *ProductService) GetAllProduct(ctx context.Context) ([]transport.ProductItemResponse, error) {
	products, err := s.repo.GetAllProduct(ctx)
	if err != nil {
		fmt.Print("s.repo.GetAllProduct() Error: ", err.Error())
		return nil, err
	}
	if len(products) == 0 {
		return []transport.ProductItemResponse{}, nil
	}

	productsResponse := transformProduct(products)

	return productsResponse, nil
}

func (s *ProductService) GetProductByUUID(ctx context.Context, uuid string) (transport.ProductItemResponse, error) {
	product, err := s.repo.GetProductByUUID(ctx, uuid)
	if err != nil {
		fmt.Print("s.repo.GetProductByUUID() Error: ", err.Error())
		return transport.ProductItemResponse{}, err
	}

	var categoryResponse *transport.CategoryItemResponse
	if product.Category != nil {
		categoryResponse = &transport.CategoryItemResponse{
			ID:          product.Category.UUID,
			Name:        product.Category.Name,
			Description: product.Category.Description,
		}
	}

	productResponse := transport.ProductItemResponse{
		ID:       product.UUID,
		Name:     product.Name,
		Stock:    product.Stock,
		Price:    product.Price,
		Category: categoryResponse,
	}

	return productResponse, nil
}

func transformProduct(p []model.Product) []transport.ProductItemResponse {
	var productsResponse []transport.ProductItemResponse
	for _, product := range p {
		var categoryResponse *transport.CategoryItemResponse
		if product.Category != nil {
			categoryResponse = &transport.CategoryItemResponse{
				ID:          product.Category.UUID,
				Name:        product.Category.Name,
				Description: product.Category.Description,
			}
		}

		productResponse := transport.ProductItemResponse{
			ID:       product.UUID,
			Name:     product.Name,
			Stock:    product.Stock,
			Price:    product.Price,
			Category: categoryResponse,
		}
		productsResponse = append(productsResponse, productResponse)
	}

	return productsResponse
}

func (s *ProductService) CreateProduct(ctx context.Context, req transport.ProductRequest) (transport.ProductItemResponse, error) {
	randomUUID := uuid.New().String()

	var categoryID *int64
	if req.CategoryID != "" {
		// Fetch category to get the integer ID
		category, err := s.categoryRepo.GetCategoryByUUID(ctx, req.CategoryID)
		if err != nil {
			fmt.Print("s.categoryRepo.GetCategoryByUUID() Error: ", err.Error())
			return transport.ProductItemResponse{}, err
		}
		if category.UUID != "" {
			categoryID = &category.ID
		}
	}

	newProduct := model.Product{
		UUID:       randomUUID,
		Name:       req.Name,
		Stock:      req.Stock,
		Price:      req.Price,
		CategoryID: categoryID,
	}

	err := s.repo.CreateProduct(ctx, newProduct)
	if err != nil {
		fmt.Print("s.repo.CreateProduct() Error: ", err.Error())
		return transport.ProductItemResponse{}, err
	}

	// Fetch the created product to get category info
	createdProduct, err := s.repo.GetProductByUUID(ctx, randomUUID)
	if err != nil {
		fmt.Print("s.repo.GetProductByUUID() Error: ", err.Error())
		return transport.ProductItemResponse{}, err
	}

	var categoryResponse *transport.CategoryItemResponse
	if createdProduct.Category != nil {
		categoryResponse = &transport.CategoryItemResponse{
			ID:          createdProduct.Category.UUID,
			Name:        createdProduct.Category.Name,
			Description: createdProduct.Category.Description,
		}
	}

	productResponse := transport.ProductItemResponse{
		ID:       createdProduct.UUID,
		Name:     createdProduct.Name,
		Stock:    createdProduct.Stock,
		Price:    createdProduct.Price,
		Category: categoryResponse,
	}

	return productResponse, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, id string, req transport.ProductRequest) (transport.ProductItemResponse, error) {
	var categoryID *int64
	if req.CategoryID != "" {
		// Fetch category to get the integer ID
		category, err := s.categoryRepo.GetCategoryByUUID(ctx, req.CategoryID)
		if err != nil {
			fmt.Print("s.categoryRepo.GetCategoryByUUID() Error: ", err.Error())
			return transport.ProductItemResponse{}, err
		}
		if category.UUID != "" {
			categoryID = &category.ID
		}
	}

	newProduct := model.Product{
		UUID:       id,
		Name:       req.Name,
		Stock:      req.Stock,
		Price:      req.Price,
		CategoryID: categoryID,
	}

	err := s.repo.UpdateProduct(ctx, newProduct)
	if err != nil {
		fmt.Print("s.repo.UpdateProduct() Error: ", err.Error())
		return transport.ProductItemResponse{}, err
	}

	// Fetch the updated product to get category info
	updatedProduct, err := s.repo.GetProductByUUID(ctx, id)
	if err != nil {
		fmt.Print("s.repo.GetProductByUUID() Error: ", err.Error())
		return transport.ProductItemResponse{}, err
	}

	var categoryResponse *transport.CategoryItemResponse
	if updatedProduct.Category != nil {
		categoryResponse = &transport.CategoryItemResponse{
			ID:          updatedProduct.Category.UUID,
			Name:        updatedProduct.Category.Name,
			Description: updatedProduct.Category.Description,
		}
	}

	productResponse := transport.ProductItemResponse{
		ID:       updatedProduct.UUID,
		Name:     updatedProduct.Name,
		Stock:    updatedProduct.Stock,
		Price:    updatedProduct.Price,
		Category: categoryResponse,
	}

	return productResponse, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, id string) error {
	err := s.repo.DeleteProduct(ctx, id)
	if err != nil {
		fmt.Print("s.repo.DeleteProduct() Error: ", err.Error())
		return err
	}

	return nil
}
