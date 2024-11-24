package services

import (
	"errors"
	"mncTest/internal/app/models"
	"mncTest/internal/app/repositories"
	"mncTest/internal/utils"
)

func Login(username, password string) (*models.Customer, error) {
	filePath := "/home/enigma/GolandProjects/mnc_test/data/customer.json"
	var customers []models.Customer
	err := utils.ReadJson(filePath, &customers)
	if err != nil {
		return nil, errors.New("Bad file json")
	}

	customer := repositories.FindByUsernameAndPassword(username, password)
	if customer == nil {
		return nil, errors.New("username or password error")
	}

	customer.IsLogin = true
	updateCustomer, err := repositories.UpdateCustomer(*customer)
	if err != nil {
		return nil, errors.New("Failed to write to JSON file")
	}

	return &updateCustomer, nil
}

func Logout(authHeader string) (*models.Customer, error) {
	id, err := utils.AuthUtil(authHeader)
	if err != nil {
		return nil, err
	}

	customer := repositories.FindById(id)
	if customer == nil {
		return nil, errors.New("customer not found")
	}

	customer.IsLogin = false
	updatedCustomer, err := repositories.UpdateCustomer(*customer)
	if err != nil {
		return nil, errors.New("failed to update customer data")
	}

	return &updatedCustomer, nil
}
