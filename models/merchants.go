package models

type Merchant struct {
	CompanyName     string    `json:"company_name"`
	Email           string    `json:"email"`
	MerchantAddress string    `json:"merchant_address"`
	DiscountOffered float64   `json:"discount_offered"`
	Products        []Product `json:"products"`
	//ID              string    `json:"id"`
	Credentials MerchCred `json:"cred"`
}
