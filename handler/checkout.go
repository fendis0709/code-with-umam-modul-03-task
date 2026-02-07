package handler

import (
	"encoding/json"
	"fendi/modul-03-task/service"
	"fendi/modul-03-task/transport"
	"fmt"
	"net/http"
)

type CheckoutHandler struct {
	service *service.CheckoutService
}

func NewCheckoutHandler(service *service.CheckoutService) *CheckoutHandler {
	return &CheckoutHandler{service: service}
}

func (h *CheckoutHandler) HandleCheckout(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		h.Checkout(w, r)
		return
	}

	http.NotFound(w, r)
}

func (h *CheckoutHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	var checkoutReq transport.CheckoutRequest
	err := json.NewDecoder(r.Body).Decode(&checkoutReq)
	if err != nil {
		fmt.Print("handler.checkout.Checkout() Decode Error: ", err.Error())
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	res, err := h.service.CreateCheckout(r.Context(), checkoutReq)
	if err != nil {
		fmt.Print("handler.checkout.Checkout() Error: ", err.Error())
		if err.Error() == "no products found for the given UUIDs" {
			http.Error(w, "No Products Found", http.StatusBadRequest)
			return
		}

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
