package services

import (
	"errors"
	"github.com/google/uuid"
	"mncTest/internal/app/models"
	"mncTest/internal/app/pkg/data"
	"mncTest/internal/app/repositories"
	"mncTest/internal/utils"
	"time"
)

func CreateTransaction(authHeader, merchantId string, amount float64) (*models.Transaction, error) {
	id, err := utils.AuthUtil(authHeader)
	if err != nil {
		return nil, err
	}
	customer := repositories.FindById(id)
	if customer == nil {
		return nil, errors.New("customer not found")
	}
	if !customer.IsLogin {
		return nil, errors.New("invalid authorization header")
	}
	err = data.BackupFile()
	if err != nil {
		return nil, err
	}
	if customer.Balance < amount {
		return nil, errors.New("insufficient balance")
	}
	merchant, err := repositories.FindMerchantById(merchantId)
	if err != nil {
		return nil, err
	}
	transaction := models.Transaction{
		ID:         uuid.New().String(),
		Timestamp:  time.Now(),
		CustomerID: customer.Id,
		MerchantID: merchant.ID,
		Amount:     amount,
		Status:     "Success",
	}
	_, err = repositories.AddTransaction(transaction)
	if err != nil {
		_ = data.RollbackFile()
		return nil, err
	}
	customer.Balance = customer.Balance - amount
	_, err = repositories.UpdateCustomer(*customer)
	if err != nil {
		_ = data.RollbackFile()
		return nil, err
	}
	return &transaction, nil
}

func GetAllTransactions(authHeader string) ([]models.Transaction, error) {
	id, err := utils.AuthUtil(authHeader)
	if err != nil {
		return nil, err
	}
	customer := repositories.FindById(id)
	if !customer.IsLogin {
		return nil, errors.New("invalid authorization header")
	}
	transactions, err := repositories.GetAllTransactionByCustomerId(id)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
