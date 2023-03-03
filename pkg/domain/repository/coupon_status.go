package repository

import (
	"context"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
	"github.com/google/uuid"
)

type CouponStatus interface {
	Get(context.Context, uuid.UUID) (*model.CouponStatus, error)
	Create(ctx context.Context, couponStatus *model.CouponStatus) (*model.CouponStatus, error)
	Use(ctx context.Context, couponID ,userID uuid.UUID) (*model.CouponStatus, error)
	ListByUserID(ctx context.Context, couponID uuid.UUID, userID uuid.UUID) (*model.CouponStatus, error)
}
