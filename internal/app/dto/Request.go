package dto

type LoginRequest struct {
	Username string `json:"username" required:"true"`
	Password string `json:"password" required:"true"`
}

type TransactionRequest struct {
	MerchantId string  `json:"merchant_id" required:"true"`
	Amount     float64 `json:"amount" required:"true"`
}
