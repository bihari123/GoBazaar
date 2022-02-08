package models

type Merchant struct {
	CompanyName     string
	ID              int
	Email           string
	Address         string
	DiscountOffered float64
	Products        []Product
}
