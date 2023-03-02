package mock

import (
	"time"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
	"github.com/Doer-org/miyagi2023-server/pkg/usecase"
	"github.com/google/uuid"
)

var stampCard *model.StampCard = &model.StampCard{
	ID:        uuid.New(),
	Name:      "hoge mock",
	CreatedAt: time.Now(),
	Spots:     spots,
}

var stampCards []*model.StampCard = []*model.StampCard{
	{
		ID:        uuid.New(),
		Name:      "hoge1 mock",
		CreatedAt: time.Now(),
		Spots:     spots,
	},
	{
		ID:        uuid.New(),
		Name:      "hoge2 mock",
		CreatedAt: time.Now(),
		Spots:     spots,
	},
}

func NewStampCardGetOutput() *usecase.StampCardGetOutput {
	return &usecase.StampCardGetOutput{
		StampCard: stampCard,
	}
}

func NewStampCardCreateOutput() *usecase.StampCardCreateOutput {
	return &usecase.StampCardCreateOutput{
		StampCard: stampCard,
	}
}

func NewStampCardListOutput() *usecase.StampCardListOutput {
	return &usecase.StampCardListOutput{
		StampCards: stampCards,
	}
}
