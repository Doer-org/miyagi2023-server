package model

import (
	"time"

	"github.com/google/uuid"
)

type CouponStatus struct {
	ID        uuid.UUID
	UsedFlg   bool
	CreatedAt time.Time
	UpdatedAt time.Time
	Coupon    *Coupon
	User      *User
}
