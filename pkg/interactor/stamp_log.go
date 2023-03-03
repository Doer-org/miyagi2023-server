package interactor

import (
	"context"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/repository"
	"github.com/Doer-org/miyagi2023-server/pkg/interactor/mock"
	"github.com/Doer-org/miyagi2023-server/pkg/usecase"
)

type StampLog struct {
	repository repository.StampLog
}

func NewStampLog(repository repository.StampLog) *StampLog {
	return &StampLog{repository: repository}
}

func (uc *StampLog) Create(context.Context, *usecase.StampLogCreateInput) (*usecase.StampLogCreateOutput, error) {
	return mock.NewStampLogCreateOutput(), nil
}

func (uc *StampLog) List(context.Context, *usecase.StampLogListInput) (*usecase.StampLogListOutput, error) {
	return mock.NewStampLogListOutput(), nil
}
