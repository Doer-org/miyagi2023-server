package usecase

import (
	"context"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
)

type Spot interface {
	Get(context.Context, *SpotGetInput) (*SpotGetOutput, error)
	Create(context.Context, *SpotCreateInput) (*SpotCreateOutput, error)
	List(context.Context, *SpotListInput) (*SpotListOutput, error)
}

type SpotGetInput struct {
}

type SpotGetOutput struct {
	Spot *model.Spot
}

type SpotCreateInput struct {
}

type SpotCreateOutput struct {
	Spot *model.Spot
}

type SpotListInput struct {
}

type SpotListOutput struct {
	Spots []*model.Spot
}
