package models

import (
	"time"
)

type Product struct {
	ProductID          int           `json:"product_id"`
	MerchantID         int           `json:"merchant_id"`
	Name               string        `json:"name"`
	ProductDescription string        `json:"product_description"`
	Price              float64       `json:"price"`
	Stock              int           `json:"stock"`
	DeliveryTime       time.Duration `json:"deliveryTime"` //A Duration represents the elapsed time between two instants as an int64 nanosecond count. The representation limits the largest representable duration to approximately 290 years.
}
