package usecase

import (
	"context"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
)

type StampCard interface {
	Get(context.Context, *StampCardGetInput) (*StampCardGetOutput, error)
	Create(context.Context, *StampCardCreateInput) (*StampCardCreateOutput, error)
	List(context.Context, *StampCardListInput) (*StampCardListOutput, error)
}

type StampCardGetInput struct {
}
type StampCardGetOutput struct {
	StampCard *model.StampCard
}

type StampCardCreateInput struct {
}
type StampCardCreateOutput struct {
	StampCard *model.StampCard
}

type StampCardListInput struct {
}
type StampCardListOutput struct {
	StampCards []*model.StampCard
}
