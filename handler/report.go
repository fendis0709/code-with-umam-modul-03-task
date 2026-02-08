package handler

import (
	"encoding/json"
	"fendi/modul-03-task/service"
	"fmt"
	"net/http"
)

type ReportHandler struct {
	service *service.ReportService
}

func NewReportHandler(service *service.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (h *ReportHandler) HandleTodayReport(w http.ResponseWriter, r *http.Request) {
	res, err := h.service.GetTodayReport(r.Context())
	if err != nil {
		fmt.Printf("handler.report.HandleTodayReport() Error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *ReportHandler) HandleReportByDate(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	res, err := h.service.GetReportByDate(r.Context(), startDate, endDate)
	if err != nil {
		fmt.Printf("handler.report.HandleReportByDate() Error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
