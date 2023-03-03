package interactor

import (
	"context"
	"fmt"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
	"github.com/Doer-org/miyagi2023-server/pkg/domain/repository"
	"github.com/Doer-org/miyagi2023-server/pkg/interactor/mock"
	"github.com/Doer-org/miyagi2023-server/pkg/usecase"
	"github.com/google/uuid"
)

type StampLog struct {
	repositoryStampLog     repository.StampLog
	repositoryCoupon       repository.Coupon
	repositoryCouponStatus repository.CouponStatus
}

func NewStampLog(rsl repository.StampLog, rc repository.Coupon, rcs repository.CouponStatus) *StampLog {
	return &StampLog{
		repositoryStampLog:     rsl,
		repositoryCoupon:       rc,
		repositoryCouponStatus: rcs,
	}
}

func (uc *StampLog) Create(ctx context.Context, in *usecase.StampLogCreateInput) (*usecase.StampLogCreateOutput, error) {
	spotID, err := uuid.Parse(in.SpotID)
	if err != nil {
		return nil, err
	}

	if in.UserID == "" {
		return nil, fmt.Errorf("stamp_log user id is invalid")
	}

	// spotIDから紐づいているクーポンを取得する
	coupon, err := uc.repositoryCoupon.GetBySpotID(ctx, spotID)
	if err != nil {
		return nil, err
	}

	// 何回きたか確認
	visitCnt, err := uc.repositoryStampLog.GetVisitCnt(ctx, in.UserID, spotID)
	if err != nil {
		return nil, err
	}

	// もし初めてきた場合、couponを発行する
	var couponStatus *model.CouponStatus
	if visitCnt == 0 {
		c := &model.CouponStatus{
			ID:      uuid.New(),
			UsedFlg: false,
			Coupon:  coupon,
			User: &model.User{
				ID: model.UserID(in.UserID),
			},
		}
		couponStatus, err = uc.repositoryCouponStatus.Create(ctx, c)
		if err != nil {
			return nil, err
		}
	}

	// stamp logの保存
	s := &model.StampLog{
		ID: uuid.New(),
		Spot: &model.Spot{
			ID: spotID,
		},
		User: &model.User{
			ID: model.UserID(in.UserID),
		},
	}

	stampLog, err := uc.repositoryStampLog.Create(ctx, s)
	if err != nil {
		return nil, err
	}

	return &usecase.StampLogCreateOutput{
		StampLog:     stampLog,
		CouponStatus: couponStatus,
		VisitCnt:     visitCnt,
	}, nil
}

func (uc *StampLog) List(context.Context, *usecase.StampLogListInput) (*usecase.StampLogListOutput, error) {
	// 今回は使わないのでいったんmockのままにしておく
	return mock.NewStampLogListOutput(), nil
}
