package mock

import (
	"time"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
	"github.com/Doer-org/miyagi2023-server/pkg/usecase"
	"github.com/google/uuid"
)

var coupon_status *model.CouponStatus = &model.CouponStatus{
	ID:        uuid.New(),
	UsedFlg:   false,
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
	Coupon:    coupon,
	User:      User,
}

var coupon_statuses []*model.CouponStatus = []*model.CouponStatus{
	{
		ID:        uuid.New(),
		UsedFlg:   false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Coupon:    coupon,
		User:      User,
	},
	{
		ID:        uuid.New(),
		UsedFlg:   true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Coupon:    coupon,
		User:      User,
	},
}

func NewCouponStatusGetOutput() *usecase.CouponStatusGetOutput {
	return &usecase.CouponStatusGetOutput{
		CouponStatus: coupon_status,
	}
}

func NewCouponStatusCreateOutput() *usecase.CouponStatusCreateOutput {
	return &usecase.CouponStatusCreateOutput{
		CouponStatus: coupon_status,
	}
}

func NewCouponStatusUseOutput() *usecase.CouponStatusUseOutput {
	return &usecase.CouponStatusUseOutput{
		CouponStatus: coupon_status,
	}
}

func NewCouponStatusListByUserOutput() *usecase.CouponStatusListByUserOutput {
	return &usecase.CouponStatusListByUserOutput{
		CouponStatuses: coupon_statuses,
	}
}
