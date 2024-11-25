package repositories

import (
	"errors"
	"github.com/joho/godotenv"
	"log"
	"mncTest/internal/app/models"
	"mncTest/internal/utils"
	"os"
)

var filePathTransaction string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	filePathTransaction = os.Getenv("PATH_FILE") + "transaction.json"
}
func GetAllTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	utils.ReadJson(filePathTransaction, &transactions)
	if transactions == nil {
		return nil, errors.New("no transactions found")
	}
	return transactions, nil
}

func GetAllTransactionByCustomerId(customerId string) ([]models.Transaction, error) {
	var customerTransactions []models.Transaction
	transactions, err := GetAllTransactions()
	if err != nil {
		return nil, err
	}
	for _, transaction := range transactions {
		if transaction.CustomerID == customerId {
			customerTransactions = append(customerTransactions, transaction)
		}
	}
	return customerTransactions, nil
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
