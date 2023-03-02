package mock

import (
	"time"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
	"github.com/Doer-org/miyagi2023-server/pkg/usecase"
	"github.com/google/uuid"
)

var coupon *model.Coupon = &model.Coupon{
	ID:             uuid.New(),
	Name:           "hoge mock",
	ExpirationDate: 100,
	DiscountRate:   10,
	CreatedAt:      time.Now(),
	Spot:           spot,
}

var coupons []*model.Coupon = []*model.Coupon{
	{
		ID:             uuid.New(),
		Name:           "hoge1 mock",
		ExpirationDate: 100,
		DiscountRate:   10,
		CreatedAt:      time.Now(),
		Spot:           spot,
	},
	{
		ID:             uuid.New(),
		Name:           "hoge2 mock",
		ExpirationDate: 200,
		DiscountRate:   20,
		CreatedAt:      time.Now(),
		Spot:           spot,
	},
	{
		ID:             uuid.New(),
		Name:           "hoge3 mock",
		ExpirationDate: 300,
		DiscountRate:   30,
		CreatedAt:      time.Now(),
		Spot:           spot,
	},
}

func NewCouponGetOutput() *usecase.CouponGetOutput {
	return &usecase.CouponGetOutput{
		Coupon: coupon,
	}
}

func NewCouponCreateOutput() *usecase.CouponCreateOutput {
	return &usecase.CouponCreateOutput{
		Coupon: coupon,
	}
}

func NewCouponListOutput() *usecase.CouponListOutput {
	return &usecase.CouponListOutput{
		Coupons: coupons,
	}
}
