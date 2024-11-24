package unit

import (
	"mncTest/internal/app/models"
	"mncTest/internal/utils"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadJson(t *testing.T) {
	filePath := filepath.Join("..", "testdata", "customer_sample.json")

	var customers []models.Customer

	err := utils.ReadJson(filePath, &customers)

	assert.NoError(t, err, "ReadJson returned an unexpected error")
	assert.Len(t, customers, 2, "Expected two users in the result")
	assert.Equal(t, "darul", customers[0].Username, "First user's name does not match")
	assert.Equal(t, "liza", customers[1].Username, "Second user's name does not match")
}

func TestWriteJson(t *testing.T) {
	filePath := filepath.Join("..", "testdata", "output_sample.json")

	customers := []models.Customer{
		{
			Id:       1,
			Username: "darul",
			Password: "123",
			Balance:  50000,
			IsLogin:  true,
		},
		{
			Id:       2,
			Username: "liza",
			Password: "123",
			Balance:  0,
			IsLogin:  true,
		},
	}

	err := utils.WriteJson(filePath, customers)

	assert.NoError(t, err, "WriteJson returned an unexpected error")

	_, err = os.Stat(filePath)
	assert.NoError(t, err, "Output file does not exist")

	var readCustomers []models.Customer
	err = utils.ReadJson(filePath, &readCustomers)
	assert.NoError(t, err, "ReadJson returned an unexpected error")
	assert.Len(t, readCustomers, 2, "Expected two users in the result")
	assert.Equal(t, "darul", readCustomers[0].Username, "First user's name does not match")
	assert.Equal(t, "liza", readCustomers[1].Username, "Second user's name does not match")

	err = os.Remove(filePath)
	assert.NoError(t, err, "Failed to remove output file")
}
