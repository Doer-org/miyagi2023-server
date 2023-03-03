package dto

import (
	"time"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
	"github.com/google/uuid"
)

type Coupon struct {
	ID             string    `db:"id"`
	Name           string    `db:"name"`
	ExpirationDate uint      `db:"expiration_date"`
	DiscountRate   uint      `db:"discount_rate"`
	CreatedAt      time.Time `db:"created_at"`
	SpotID         string    `db:"spot_id"`
}

func NewCouponDotFromModel(m *model.Coupon) *Coupon {
	return &Coupon{
		ID:             m.ID.String(),
		Name:           m.Name,
		ExpirationDate: m.ExpirationDate,
		DiscountRate:   m.DiscountRate,
		CreatedAt:      m.CreatedAt,
		SpotID:         m.Spot.ID.String(),
	}
}

func (d *Coupon) ToModel(dtoSpot *Spot) (*model.Coupon, error) {
	id, err := uuid.Parse(d.ID)
	if err != nil {
		return nil, err
	}

	spot, err := dtoSpot.ToModel()
	if err != nil {
		return nil, err
	}

	return &model.Coupon{
		ID:             id,
		Name:           d.Name,
		ExpirationDate: d.ExpirationDate,
		DiscountRate:   d.DiscountRate,
		CreatedAt:      d.CreatedAt,
		Spot:           spot,
	}, nil
}
