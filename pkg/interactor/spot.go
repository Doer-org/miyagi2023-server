package interactor

import (
	"context"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/repository"
	"github.com/Doer-org/miyagi2023-server/pkg/interactor/mock"
	"github.com/Doer-org/miyagi2023-server/pkg/usecase"
)

type Spot struct {
	repostiroy repository.Spot
}

func NewSpot(repository repository.Spot) *Spot {
	return &Spot{repostiroy: repository}
}

func (uc *Spot) Get(context.Context, *usecase.SpotGetInput) (*usecase.SpotGetOutput, error) {
	return mock.NewSpotGetOutput(), nil
}

func (uc *Spot) Create(context.Context, *usecase.SpotCreateInput) (*usecase.SpotCreateOutput, error) {
	return mock.NewSpotCreateOutput(), nil
}

func (uc *Spot) List(context.Context, *usecase.SpotListInput) (*usecase.SpotListOutput, error) {
	return mock.NewSpotListOutput(), nil
}
