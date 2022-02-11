package models

type User struct {
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	Email         string  `json:"email"`
	Contact       string  `json:"contact"`
	City          string  `json:"city"`
	WalletBalance float64 `json:"wallet_balance"`
	//ID            string  `json:"id"`
	//Password      string  `json:"password"`
	Credentials UserCred `json:"cred"`
}
