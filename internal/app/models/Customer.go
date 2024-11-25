package models

type Customer struct {
	Id       string  `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
	IsLogin  bool    `json:"is_login"`
}
