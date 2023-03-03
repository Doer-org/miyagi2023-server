package interactor

import (
	"context"
	"fmt"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
	"github.com/Doer-org/miyagi2023-server/pkg/domain/repository"
	"github.com/Doer-org/miyagi2023-server/pkg/usecase"
	"github.com/google/uuid"
)

type Spot struct {
	repository repository.Spot
}

func NewSpot(repository repository.Spot) *Spot {
	return &Spot{repository: repository}
}

func (uc *Spot) Get(ctx context.Context, in *usecase.SpotGetInput) (*usecase.SpotGetOutput, error) {
	id, err := uuid.Parse(in.ID)
	if err != nil {
		return nil, err
	}

	spot, err := uc.repository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &usecase.SpotGetOutput{
		Spot: spot,
	}, nil
}

func (uc *Spot) Create(ctx context.Context, in *usecase.SpotCreateInput) (*usecase.SpotCreateOutput, error) {
	if in.Name == "" {
		return nil, fmt.Errorf("spot name is invalid")
	}
	if len(in.Name) > 255 {
		return nil, fmt.Errorf("spot name must be 255 characters or less")
	}

	if in.Detail == "" {
		return nil, fmt.Errorf("spot detail is invalid")
	}
	if len(in.Detail) > 255 {
		return nil, fmt.Errorf("spot detail must be 255 characters or less")
	}

	if in.ImgURL == "" {
		return nil, fmt.Errorf("spot img_url is invalid")
	}
	if len(in.Detail) > 255 {
		return nil, fmt.Errorf("spot img_url must be 255 characters or less")
	}

	if in.Address == "" {
		return nil, fmt.Errorf("spot address is invalid")
	}
	if len(in.Address) > 255 {
		return nil, fmt.Errorf("spot address must be 255 characters or less")
	}

	s := &model.Spot{
		ID:      uuid.New(),
		Name:    in.Name,
		Detail:  in.Detail,
		Like:    0,
		ImgURL:  in.ImgURL,
		Address: in.Address,
	}

	newSpot, err := uc.repository.Create(ctx, s)
	if err != nil {
		return nil, err
	}

	return &usecase.SpotCreateOutput{
		Spot: newSpot,
	}, nil
}

func (uc *Spot) List(ctx context.Context, in *usecase.SpotListInput) (*usecase.SpotListOutput, error) {
	spots, err := uc.repository.List(ctx)
	if err != nil {
		return nil, err
	}

	return &usecase.SpotListOutput{
		Spots: spots,
	}, nil
}

func (uc *Spot) GetRandom(ctx context.Context) (*usecase.SpotListOutput, error) {
	spots, err := uc.repository.GetRandom(ctx)
	if err != nil {
		return nil, err
	}

	return &usecase.SpotListOutput{
		Spots: spots,
	}, nil
}
