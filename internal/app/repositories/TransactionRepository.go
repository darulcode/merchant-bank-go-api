package repositories

import (
	"errors"
	"mncTest/internal/app/models"
	"mncTest/internal/utils"
)

const filePathTransaction = "/home/enigma/GolandProjects/mnc_test/data/transaction.json"

func GetAllTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	utils.ReadJson(filePathTransaction, &transactions)
	if transactions == nil {
		return nil, errors.New("No transactions found")
	}
	return transactions, nil
}

func AddTransaction(Transaction models.Transaction) (*models.Transaction, error) {
	getAllTransactions, err := GetAllTransactions()
	if err != nil {
		return nil, err
	}
	getAllTransactions = append(getAllTransactions, Transaction)
	utils.WriteJson(filePathTransaction, getAllTransactions)
	return &Transaction, nil
}

func UpdateTransaction(transaction models.Transaction) (*models.Transaction, error) {
	transactions, err := GetAllTransactions()
	if err != nil {
		return nil, err
	}
	for i := range transactions {
		if transactions[i] == transaction {
			transactions[i] = transaction
			utils.WriteJson(filePathTransaction, transactions)
			return &transactions[i], nil
		}
	}
	return nil, errors.New("Transaction not found")
}
