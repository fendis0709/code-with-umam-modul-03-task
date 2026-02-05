package handler

import (
	"encoding/json"
	"fendi/modul-02-task/service"
	"fendi/modul-02-task/transport"
	"fmt"
	"net/http"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) HandleProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.GetAllProduct(w, r)
		return
	}
	if r.Method == http.MethodPost {
		h.CreateProduct(w, r)
		return
	}

	http.NotFound(w, r)
}

func (h *ProductHandler) GetAllProduct(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("search")

	res, err := h.service.GetAllProduct(r.Context(), keyword)
	if err != nil {
		fmt.Print("handler.product.GetAllProduct() Error: ", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productReq transport.ProductRequest
	err := json.NewDecoder(r.Body).Decode(&productReq)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	res, err := h.service.CreateProduct(r.Context(), productReq)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (h *ProductHandler) HandleProductItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.GetProductByUUID(w, r)
		return
	}
	if r.Method == http.MethodPut {
		h.UpdateProduct(w, r)
		return
	}
	if r.Method == http.MethodDelete {
		h.DeleteProduct(w, r)
		return
	}

	http.NotFound(w, r)
}

func (h *ProductHandler) GetProductByUUID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/products/"):]
	if idStr == "" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	res, err := h.service.GetProductByUUID(r.Context(), idStr)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if res.ID == "" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/products/"):]
	if idStr == "" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	var productReq transport.ProductRequest
	err := json.NewDecoder(r.Body).Decode(&productReq)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	res, err := h.service.UpdateProduct(r.Context(), idStr, productReq)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/products/"):]
	if idStr == "" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	err := h.service.DeleteProduct(r.Context(), idStr)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transport.StatusResponse{
		Code:   http.StatusOK,
		Status: "OK",
	})
}
