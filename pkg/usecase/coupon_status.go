package usecase

import (
	"context"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
)

type CouponStatus interface {
	Get(context.Context, *CouponStatusGetInput) (*CouponStatusGetOutput, error)
	Create(context.Context, *CouponStatusCreateInput) (*CouponStatusCreateOutput, error)
	Use(context.Context, *CouponStatusUseInput) (*CouponStatusUseOutput, error)
	ListByUserID(context.Context, *CouponStatusListByUserInput) (*CouponStatusListByUserOutput, error)
}

type CouponStatusGetInput struct{}
type CouponStatusGetOutput struct {
	CouponStatus *model.CouponStatus
}

type CouponStatusCreateInput struct{}
type CouponStatusCreateOutput struct {
	CouponStatus *model.CouponStatus
}

type CouponStatusUseInput struct{}
type CouponStatusUseOutput struct {
	CouponStatus *model.CouponStatus
}

type CouponStatusListByUserInput struct{}
type CouponStatusListByUserOutput struct {
	CouponStatuses []*model.CouponStatus
}
