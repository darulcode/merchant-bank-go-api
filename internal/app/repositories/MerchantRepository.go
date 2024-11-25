package repositories

import (
	"errors"
	"github.com/joho/godotenv"
	"log"
	"mncTest/internal/app/models"
	"mncTest/internal/utils"
	"os"
)

var filePathMerchant string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	filePathMerchant = os.Getenv("PATH_FILE") + "merchant.json"
}

func GetAllMerchants() ([]models.Merchant, error) {
	var merchants []models.Merchant
	err := utils.ReadJson(filePathMerchant, &merchants)
	if err != nil {
		return nil, err
	}
	return merchants, nil
}

func FindMerchantById(merchantId string) (*models.Merchant, error) {
	merchants, err := GetAllMerchants()
	if err != nil {
		return nil, err
	}
	for _, merchant := range merchants {
		if merchant.ID == merchantId {
			return &merchant, nil
		}
	}
	return nil, errors.New("Merchant not found")
}
