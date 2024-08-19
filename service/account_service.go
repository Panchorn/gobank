package service

import (
	"gobank/errs"
	"gobank/logs"
	"gobank/repository"
	"strings"
	"time"
)

type accountService struct {
	repository repository.AccountRepository
}

func NewAccountService(repository repository.AccountRepository) AccountService {
	return accountService{repository: repository}
}

func (s accountService) NewAccount(customerID int, request NewAccountRequest) (*AccountResponse, error) {
	// Validate
	if request.Amount < 5000 {
		return nil, errs.NewValidationError("amount at least 5,000")
	}
	if strings.ToLower(request.AccountType) != "saving" && strings.ToLower(request.AccountType) != "checking" {
		return nil, errs.NewValidationError("account type must be 'saving' or 'checking'")
	}

	account := repository.Account{
		CustomerID:  customerID,
		OpeningDate: time.Now().Format("2006-1-2 15:04:05"),
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      1,
	}
	newAccount, err := s.repository.Create(account)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	response := AccountResponse{
		AccountID:   newAccount.AccountID,
		OpeningDate: newAccount.OpeningDate,
		AccountType: newAccount.AccountType,
		Amount:      newAccount.Amount,
		Status:      newAccount.Status,
	}
	return &response, nil
}

func (s accountService) GetAccounts(customerID int) ([]AccountResponse, error) {
	accounts, err := s.repository.GetAll(customerID)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	responses := []AccountResponse{}
	for _, account := range accounts {
		response := AccountResponse{
			AccountID:   account.AccountID,
			OpeningDate: account.OpeningDate,
			AccountType: account.AccountType,
			Amount:      account.Amount,
			Status:      account.Status,
		}
		responses = append(responses, response)
	}
	return responses, nil
}
