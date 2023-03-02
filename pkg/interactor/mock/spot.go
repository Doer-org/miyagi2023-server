package mock

import (
	"time"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
	"github.com/Doer-org/miyagi2023-server/pkg/usecase"
	"github.com/google/uuid"
)

var spot *model.Spot = &model.Spot{
	ID:        uuid.New(),
	Name:      "hoge mock",
	Detail:    "hogehoge",
	Like:      23,
	CreatedAt: time.Now(),
}

var spots []*model.Spot = []*model.Spot{
	{
		ID:        uuid.New(),
		Name:      "hoge1 mock",
		Detail:    "hogehoge1",
		Like:      123,
		CreatedAt: time.Now(),
	},
	{
		ID:        uuid.New(),
		Name:      "hoge2 mock",
		Detail:    "hogehoge2",
		Like:      223,
		CreatedAt: time.Now(),
	},
	{
		ID:        uuid.New(),
		Name:      "hoge3 mock",
		Detail:    "hogehoge3",
		Like:      323,
		CreatedAt: time.Now(),
	},
}

func NewSpotGetOutput() *usecase.SpotGetOutput {
	return &usecase.SpotGetOutput{
		Spot: spot,
	}
}

func NewSpotCreateOutput() *usecase.SpotCreateOutput {
	return &usecase.SpotCreateOutput{
		Spot: &model.Spot{
			ID:        uuid.New(),
			Name:      "hoge mock",
			Detail:    "hogehoge",
			Like:      23,
			CreatedAt: time.Now(),
		},
	}
}

func NewSpotListOutput() *usecase.SpotListOutput {
	return &usecase.SpotListOutput{
		Spots: spots,
	}
}
