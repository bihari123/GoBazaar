package models

import "github.com/google/uuid"

type Cred struct {
	UserID   uuid.UUID `json:"userIs"`
	UserPass string    `json:"userPass"`
}
