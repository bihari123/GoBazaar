package models

import "time"

type Product struct {
	ID           int
	Name         string
	Description  string
	Price        float64
	Stock        int
	DeliveryTime time.Duration //A Duration represents the elapsed time between two instants as an int64 nanosecond count. The representation limits the largest representable duration to approximately 290 years.
}
