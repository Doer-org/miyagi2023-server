package interactor

import (
	"context"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/repository"
	"github.com/Doer-org/miyagi2023-server/pkg/interactor/mock"
	"github.com/Doer-org/miyagi2023-server/pkg/usecase"
)

type CouponStatus struct {
	repository repository.CouponStatus
}

func NewCouponStatus(repository repository.CouponStatus) *CouponStatus {
	return &CouponStatus{repository: repository}
}

func (uc *CouponStatus) Get(ctx context.Context, in *usecase.CouponStatusGetInput) (*usecase.CouponStatusGetOutput, error) {
	return mock.NewCouponStatusGetOutput(), nil
}

func (uc *CouponStatus) Create(ctx context.Context, in *usecase.CouponStatusCreateInput) (*usecase.CouponStatusCreateOutput, error) {
	return mock.NewCouponStatusCreateOutput(), nil
}

func (uc *CouponStatus) Use(ctx context.Context, in *usecase.CouponStatusUseInput) (*usecase.CouponStatusUseOutput, error) {
	return mock.NewCouponStatusUseOutput(), nil
}

func (uc *CouponStatus) ListByUserID(ctx context.Context, in *usecase.CouponStatusListByUserInput) (*usecase.CouponStatusListByUserOutput, error) {
	return mock.NewCouponStatusListByUserOutput(), nil
}
