package interactor

import (
	"context"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/repository"
	"github.com/Doer-org/miyagi2023-server/pkg/interactor/mock"
	"github.com/Doer-org/miyagi2023-server/pkg/usecase"
)

type Spot struct {
	repository repository.Spot
}

func NewSpot(repository repository.Spot) *Spot {
	return &Spot{repository: repository}
}

func (uc *Spot) Get(ctx context.Context, in *usecase.SpotGetInput) (*usecase.SpotGetOutput, error) {
	return mock.NewSpotGetOutput(),nil
}

func (uc *Spot) Create(ctx context.Context, in *usecase.SpotCreateInput) (*usecase.SpotCreateOutput, error) {
	return mock.NewSpotCreateOutput(),nil
}

func (uc *Spot) List(ctx context.Context, in *usecase.SpotListInput) (*usecase.SpotListOutput, error) {
	return mock.NewSpotListOutput(),nil
}
