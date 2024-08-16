package repository

import "errors"

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() customerRepositoryMock {
	customers := []Customer{
		{CustomerID: 1001, Name: "Ashish", City: "New Delhi", ZipCode: "110011", DateOfBirth: "2000-01-01", Status: 1},
		{CustomerID: 1002, Name: "Rob", City: "New Delhi", ZipCode: "110011", DateOfBirth: "2000-01-01", Status: 0},
	}
	return customerRepositoryMock{customers: customers}
}

func (m customerRepositoryMock) GetAll() ([]Customer, error) {
	return m.customers, nil
}

func (m customerRepositoryMock) GetById(id int) (*Customer, error) {
	for _, customer := range m.customers {
		if customer.CustomerID == id {
			return &customer, nil
		}
	}
	return nil, errors.New("Customer not found")
}
