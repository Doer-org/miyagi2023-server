package repository

import (
	"context"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
	"github.com/google/uuid"
)

type Coupon interface {
	Get(context.Context, uuid.UUID) (*model.Coupon, error)
	Create(context.Context, uuid.UUID) (*model.Coupon, error)
	List(context.Context) ([]*model.Coupon, error)
}
