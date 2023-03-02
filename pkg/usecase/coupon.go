package usecase

import (
	"context"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
)

type Coupon interface {
	Get(context.Context, *CouponGetInput) (*CouponGetOutput, error)
	Create(context.Context, *CouponCreateInput) (*CouponCreateOutput, error)
	List(context.Context, *CouponListInput) (*CouponListOutput, error)
}

type CouponGetInput struct{}
type CouponGetOutput struct {
	Coupon *model.Coupon
}

type CouponCreateInput struct{}
type CouponCreateOutput struct {
	Coupon *model.Coupon
}

type CouponListInput struct{}
type CouponListOutput struct {
	Coupons []*model.Coupon
}
