package models

import "github.com/google/uuid"

type Merchant struct {
	CompanyName     string    `json:"company_name"`
	Email           string    `json:"email"`
	Address         string    `json:"address"`
	DiscountOffered float64   `json:"discount_offered"`
	Products        []Product `json:"products"`
	ID              uuid.UUID `json:"id"`
}
