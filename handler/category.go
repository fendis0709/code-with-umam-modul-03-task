package handler

import (
	"encoding/json"
	"fendi/modul-03-task/service"
	"fendi/modul-03-task/transport"
	"fmt"
	"net/http"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) HandleCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.GetAllCategory(w, r)
		return
	}
	if r.Method == http.MethodPost {
		h.CreateCategory(w, r)
		return
	}

	http.NotFound(w, r)
}

func (h *CategoryHandler) GetAllCategory(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("search")

	res, err := h.service.GetAllCategory(r.Context(), keyword)
	if err != nil {
		fmt.Print("handler.category.GetAllCategory() Error: ", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var categoryReq transport.CategoryRequest
	err := json.NewDecoder(r.Body).Decode(&categoryReq)
	if err != nil {
		fmt.Print("handler.category.CreateCategory() Decode Error: ", err.Error())
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	res, err := h.service.CreateCategory(r.Context(), categoryReq)
	if err != nil {
		fmt.Print("handler.category.CreateCategory() Error: ", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (h *CategoryHandler) HandleCategoryItem(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.GetCategoryByUUID(w, r)
		return
	}
	if r.Method == http.MethodPut {
		h.UpdateCategory(w, r)
		return
	}
	if r.Method == http.MethodDelete {
		h.DeleteCategory(w, r)
		return
	}

	http.NotFound(w, r)
}

func (h *CategoryHandler) GetCategoryByUUID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/categories/"):]
	if idStr == "" {
		fmt.Print("handler.category.GetCategoryByUUID() Error: ID is empty")
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	res, err := h.service.GetCategoryByUUID(r.Context(), idStr)
	if err != nil {
		fmt.Print("handler.category.GetCategoryByUUID() Error: ", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if res.ID == "" {
		fmt.Print("handler.category.GetCategoryByUUID() Error: category not found")
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/categories/"):]
	if idStr == "" {
		fmt.Print("handler.category.UpdateCategory() Error: ID is empty")
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	var categoryReq transport.CategoryRequest
	err := json.NewDecoder(r.Body).Decode(&categoryReq)
	if err != nil {
		fmt.Print("handler.category.UpdateCategory() Decode Error: ", err.Error())
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	res, err := h.service.UpdateCategory(r.Context(), idStr, categoryReq)
	if err != nil {
		if err.Error() == "category not found" {
			fmt.Print("handler.category.UpdateCategory() Error: category not found")
			http.Error(w, "Bad Request: Category not found", http.StatusBadRequest)
			return
		}

		fmt.Print("handler.category.UpdateCategory() Error: ", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/categories/"):]
	if idStr == "" {
		fmt.Print("handler.category.DeleteCategory() Error: ID is empty")
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	err := h.service.DeleteCategory(r.Context(), idStr)
	if err != nil {
		fmt.Print("handler.category.DeleteCategory() Error: ", err.Error())
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
