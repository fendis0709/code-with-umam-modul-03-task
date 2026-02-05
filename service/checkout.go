package service

import (
	"context"
	"fendi/modul-03-task/repository"
	"fendi/modul-03-task/transport"
)

type CheckoutService struct {
	repo        *repository.CheckoutRepository
	productRepo *repository.ProductRepository
}

func NewCheckoutService(repo *repository.CheckoutRepository, productRepo *repository.ProductRepository) *CheckoutService {
	return &CheckoutService{repo: repo, productRepo: productRepo}
}

func (s *CheckoutService) CreateCheckout(ctx context.Context, req transport.CheckoutRequest) (transport.CheckoutResponse, error) {
	var SKUs []string
	for _, item := range req.Items {
		SKUs = append(SKUs, item.SKU)
	}

	products, err := s.productRepo.GetProductBySKUs(ctx, SKUs)
	if err != nil {
		return transport.CheckoutResponse{}, err
	}
	if len(products) == 0 {
		return transport.CheckoutResponse{}, nil
	}

	var totalAmount float64
	for _, item := range req.Items {
		for _, product := range products {
			if item.SKU == product.SKU {
				price := product.Price
				if price == nil {
					price = new(float64)
				}
				totalAmount += *price * float64(item.Quantity)
			}
		}
	}

	checkout := transport.CheckoutResponse{
		TotalAmount: totalAmount,
	}

	s.repo.CreateCheckoutTransaction(ctx, req)

	return checkout, nil
}
