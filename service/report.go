package service

import (
	"context"
	"fendi/modul-03-task/repository"
	"fendi/modul-03-task/transport"
	"time"
)

type ReportService struct {
	repo *repository.ReportRepository
}

func NewReportService(repo *repository.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetTodayReport(ctx context.Context) (transport.ReportResponse, error) {
	dateStart := time.Now().Format("2006-01-02") + " 00:00:00"
	dateEnd := time.Now().Format("2006-01-02") + " 23:59:59"

	report, err := s.repo.FetchReport(dateStart, dateEnd)
	if err != nil {
		return transport.ReportResponse{}, err
	}

	response := transport.ReportResponse{
		TotalRevenue:     report.TotalRevenue,
		TotalTransaction: report.TotalTransaction,
		MostPurchasedItem: &transport.MostPurchasedItemResponse{
			ProductID:   report.MostPurchasedItem.ProductID,
			ProductName: report.MostPurchasedItem.ProductName,
			Quantity:    report.MostPurchasedItem.Quantity,
		},
	}

	return response, nil
}

func (s *ReportService) GetReportByDate(ctx context.Context, startDate, endDate string) (transport.ReportResponse, error) {
	dateStart := startDate + " 00:00:00"
	dateEnd := endDate + " 23:59:59"

	report, err := s.repo.FetchReport(dateStart, dateEnd)
	if err != nil {
		return transport.ReportResponse{}, err
	}

	response := transport.ReportResponse{
		TotalRevenue:     report.TotalRevenue,
		TotalTransaction: report.TotalTransaction,
		MostPurchasedItem: &transport.MostPurchasedItemResponse{
			ProductID:   report.MostPurchasedItem.ProductID,
			ProductName: report.MostPurchasedItem.ProductName,
			Quantity:    report.MostPurchasedItem.Quantity,
		},
	}

	return response, nil
}
