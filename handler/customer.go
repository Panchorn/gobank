package handler

import (
	"encoding/json"
	"fmt"
	"gobank/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type customerHandler struct {
	service service.CustomerService
}

func NewCustomerHandler(service service.CustomerService) customerHandler {
	return customerHandler{service: service}
}

func (h customerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.service.GetCustomers()
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Fprintln(w, err)
        return
    }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (h customerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])
	customer, err := h.service.GetCustomer(customerID)
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Fprintln(w, err)
        return
    }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}
