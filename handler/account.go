package handler

import (
	"encoding/json"
	"gobank/errs"
	"gobank/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type accountHandler struct {
	service service.AccountService
}

func NewAccountHandler(service service.AccountService) accountHandler {
    return accountHandler{service: service}
}

func (h accountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])

	if r.Header.Get("Content-Type") != "application/json" {
		handleError(w, errs.NewValidationError("content type must be json"))
        return
	}

	request := service.NewAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handleError(w, errs.NewValidationError("request body incorrect format"))
		return
	}

	response, err := h.service.NewAccount(customerID, request)
	if err!= nil {
        handleError(w, err)
        return
    }

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h accountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])
	responses, err := h.service.GetAccounts(customerID)
	if err != nil {
        handleError(w, err)
        return
    }
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses)
}