package persistence

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
)

type CouponStatus struct {
	db *sqlx.DB
}

func NewCouponStatus(db *sqlx.DB) *CouponStatus {
	return &CouponStatus{db: db}
}

func (d *CouponStatus) Get(context.Context, uuid.UUID) (*model.CouponStatus, error) {
	return nil, nil
}

func (d *CouponStatus) Create(context.Context, uuid.UUID) (*model.CouponStatus, error) {
	return nil, nil
}

func (d *CouponStatus) Use(ctx context.Context, couponID uuid.UUID, userID uuid.UUID) (*model.CouponStatus, error) {
	return nil, nil
}

func (d *CouponStatus) ListByUserID(ctx context.Context, couponID uuid.UUID, userID uuid.UUID) (*model.CouponStatus, error) {
	return nil, nil
}
