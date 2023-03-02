package interactor

import (
	"context"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/repository"
	"github.com/Doer-org/miyagi2023-server/pkg/interactor/mock"
	"github.com/Doer-org/miyagi2023-server/pkg/usecase"
)

type StampCard struct {
	repository repository.StampCard
}

func NewStampCard(repository repository.StampCard) *StampCard {
	return &StampCard{repository: repository}
}

func (uc *StampCard) Get(context.Context, *usecase.StampCardGetInput) (*usecase.StampCardGetOutput, error) {
	return mock.NewStampCardGetOutput(), nil
}

func (uc *StampCard) Create(context.Context, *usecase.StampCardCreateInput) (*usecase.StampCardCreateOutput, error) {
	return mock.NewStampCardCreateOutput(), nil
}

func (uc *StampCard) List(context.Context, *usecase.StampCardListInput) (*usecase.StampCardListOutput, error) {
	return mock.NewStampCardListOutput(), nil
}
