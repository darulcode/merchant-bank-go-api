package services

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"log"
	"mncTest/internal/app/models"
	"mncTest/internal/app/repositories"
	"mncTest/internal/utils"
	"os"
)

var filePathCustomer string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	filePathCustomer = os.Getenv("PATH_FILE") + "customer.json"
}

func RegisterService(username, password string) (*models.Customer, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	usernameResult := repositories.FindByUsername(username)
	if usernameResult != nil {
		return nil, errors.New(fmt.Sprintf("User with username %s already exists", username))
	}

	customer := models.Customer{
		Id:       uuid.New().String(),
		Username: username,
		Password: string(hashedPassword),
		Balance:  0,
		IsLogin:  false,
	}
	repositories.AddCustomer(customer)
	return &customer, nil
}

func LoginService(username, password string) (*models.Customer, error) {
	var customers []models.Customer
	err := utils.ReadJson(filePathCustomer, &customers)
	if err != nil {
		return nil, errors.New("Bad file json")
	}

	customerResult := repositories.FindByUsername(username)
	err = bcrypt.CompareHashAndPassword([]byte(customerResult.Password), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, errors.New("invalid Credential")
		} else {
			return nil, err
		}
	} else {
		fmt.Println("Login successful!")
	}

	customerResult.IsLogin = true
	updateCustomer, err := repositories.UpdateCustomer(*customerResult)
	if err != nil {
		return nil, errors.New("something went wrong")
	}

	return &updateCustomer, nil
}

func LogoutService(authHeader string) (*models.Customer, error) {
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
