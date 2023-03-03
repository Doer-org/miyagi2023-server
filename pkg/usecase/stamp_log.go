package usecase

import (
	"context"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
)

type StampLog interface {
	Create(context.Context, *StampLogCreateInput) (*StampLogCreateOutput, error)
	List(context.Context, *StampLogListInput) (*StampLogListOutput, error)
}

type StampLogCreateInput struct {
	SpotID string
	UserID string
}
type StampLogCreateOutput struct {
	StampLog     *model.StampLog
	CouponStatus *model.CouponStatus
	VisitCnt     uint
}

type StampLogListInput struct {
}
type StampLogListOutput struct {
	StampLogs []*model.StampLog
}
