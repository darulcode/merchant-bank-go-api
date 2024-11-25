package repositories

import (
	"errors"
	"mncTest/internal/app/models"
	"mncTest/internal/utils"
)

const filePathCustomer = "/home/enigma/GolandProjects/mnc_test/data/customer.json"

func GetAllCustomers() ([]models.Customer, error) {
	var customers []models.Customer
	err := utils.ReadJson(filePathCustomer, &customers)
	if err != nil {
		return nil, errors.New("failed to read customers from file")
	}
	return customers, nil
}

func AddCustomer(newCustomer models.Customer) []models.Customer {
	allCustomers, _ := GetAllCustomers()
	allCustomers = append(allCustomers, newCustomer)
	utils.WriteJson(filePathCustomer, allCustomers)
	return allCustomers
}

func FindById(id string) *models.Customer {
	customers, _ := GetAllCustomers()
	for _, customer := range customers {
		if customer.Id == id {
			return &customer
		}
	}
	return nil
}

func FindByUsername(username string) *models.Customer {
	customers, _ := GetAllCustomers()
	for _, customer := range customers {
		if customer.Username == username {
			return &customer
		}
	}
	return nil
}

func FindByUsernameAndPassword(username string, password string) (customer *models.Customer) {
	customerResult := FindByUsername(username)
	if customerResult == nil || customerResult.Password != password {
		return nil
	}
	return customerResult
}

func UpdateCustomer(customer models.Customer) (models.Customer, error) {
	customers, _ := GetAllCustomers()
	for i, cust := range customers {
		if cust.Id == customer.Id {
			customers[i] = customer
			utils.WriteJson(filePathCustomer, customers)
			return customer, nil
		}
	}
	return customer, errors.New("customer not found")
}

func DeleteCustomer(id string) ([]models.Customer, bool) {
	customers, _ := GetAllCustomers()
	for i, customer := range customers {
		if customer.Id == id {
			customers = append(customers[:i], customers[i+1:]...)
			utils.WriteJson(filePathCustomer, customers)
			return customers, true
		}
	}
	return customers, false
}
