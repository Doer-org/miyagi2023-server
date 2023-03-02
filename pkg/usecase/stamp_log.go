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
}
type StampLogCreateOutput struct {
	StampLog *model.StampLog
}

type StampLogListInput struct {
}
type StampLogListOutput struct {
	StampLogs []*model.StampLog
}
