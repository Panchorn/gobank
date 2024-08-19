package service

import (
	"database/sql"
	"gobank/errs"
	"gobank/logs"
	"gobank/repository"
)

type customerService struct {
	repository repository.CustomerRepository
}

func NewCustomerService(repository repository.CustomerRepository) CustomerService {
	return customerService{repository: repository}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.repository.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	customerResponses := []CustomerResponse{}
	for _, customer := range customers {
		customerResponse := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		customerResponses = append(customerResponses, customerResponse)
	}
	return customerResponses, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.repository.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	customerResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}
	return &customerResponse, nil
}
