package model

import (
	"time"

	"github.com/google/uuid"
)

type Coupon struct {
	ID             uuid.UUID
	Name           string
	ExpirationDate uint
	DiscountRate   uint
	CreatedAt      time.Time
	Spot           *Spot
}
