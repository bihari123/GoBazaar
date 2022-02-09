package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID           uuid.UUID     `json:"id"`
	MerchantID   uuid.UUID     `json:"merchant_id"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Price        float64       `json:"price"`
	Stock        int           `json:"stock"`
	DeliveryTime time.Duration `json:"deliveryTime"` //A Duration represents the elapsed time between two instants as an int64 nanosecond count. The representation limits the largest representable duration to approximately 290 years.
}
