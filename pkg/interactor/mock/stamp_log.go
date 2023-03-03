package mock

import (
	"time"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
	"github.com/Doer-org/miyagi2023-server/pkg/usecase"
	"github.com/google/uuid"
)

var stampLog *model.StampLog = &model.StampLog{
	ID:        uuid.New(),
	Spot:      spot,
	User:      user,
	CreatedAt: time.Now(),
}

var stampLogs []*model.StampLog = []*model.StampLog{
	{
		ID:        uuid.New(),
		Spot:      spot,
		User:      user,
		CreatedAt: time.Now(),
	},
	{
		ID:        uuid.New(),
		Spot:      spot,
		User:      user,
		CreatedAt: time.Now(),
	},
}

func NewStampLogCreateOutput() *usecase.StampLogCreateOutput {
	return &usecase.StampLogCreateOutput{
		StampLog: stampLog,
	}
}

func NewStampLogListOutput() *usecase.StampLogListOutput {
	return &usecase.StampLogListOutput{
		StampLogs: stampLogs,
	}
}
