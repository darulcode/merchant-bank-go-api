package services

import (
	"errors"
	"github.com/google/uuid"
	"mncTest/internal/app/models"
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
	repositories.AddTransaction(transaction)
	customer.Balance = customer.Balance - amount
	repositories.UpdateCustomer(*customer)
	return &transaction, nil
}

func GetAllTransactions(authHeader string) ([]models.Transaction, error) {
	_, err := utils.AuthUtil(authHeader)
	if err != nil {
		return nil, err
	}
	transactions, err := repositories.GetAllTransactions()
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
