package models

import "time"

type Transaction struct {
	ID         string    `json:"id"`
	Timestamp  time.Time `json:"timestamp"`
	CustomerID string    `json:"customer_id"`
	MerchantID string    `json:"merchant_id"`
	Amount     float64   `json:"amount,omitempty"`
	Status     string    `json:"status,omitempty"`
}
