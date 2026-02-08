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
	transaction, err := s.repo.CreateCheckoutTransaction(ctx, req)
	if err != nil {
		return transport.CheckoutResponse{}, err
	}

	var itemDetails []transport.CheckoutItemResponse
	for _, detail := range transaction.Details {
		itemResp := transport.CheckoutItemResponse{
			ProductID:   detail.ProductUUID,
			ProductName: detail.ProductName,
			Quantity:    detail.Quantity,
			UnitPrice:   detail.Price,
			TotalPrice:  detail.SubTotal,
		}
		itemDetails = append(itemDetails, itemResp)
	}

	checkout := transport.CheckoutResponse{
		ID:          transaction.UUID,
		Date:        transaction.PurchasedAt.Format("2006-01-02"),
		TotalAmount: transaction.TotalAmount,
		Items:       itemDetails,
	}

	return checkout, nil
}
