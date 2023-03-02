package interactor

import (
	"context"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/repository"
	"github.com/Doer-org/miyagi2023-server/pkg/interactor/mock"
	"github.com/Doer-org/miyagi2023-server/pkg/usecase"
)

type Coupon struct {
	repository repository.Coupon
}

func NewCoupon(repository repository.Coupon) *Coupon {
	return &Coupon{repository: repository}
}

func (uc *Coupon) Get(ctx context.Context, in *usecase.CouponGetInput) (*usecase.CouponGetOutput, error) {
	return mock.NewCouponGetOutput(), nil
}

func (uc *Coupon) Create(ctx context.Context, in *usecase.CouponCreateInput) (*usecase.CouponCreateOutput, error) {
	return mock.NewCouponCreateOutput(), nil
}

func (uc *Coupon) List(ctx context.Context, in *usecase.CouponListInput) (*usecase.CouponListOutput, error) {
	return mock.NewCouponListOutput(), nil
}
