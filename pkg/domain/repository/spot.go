package repository

import (
	"context"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
	"github.com/google/uuid"
)

type Spot interface {
	Get(context.Context, uuid.UUID) (*model.Spot, error)
	Create(context.Context, *model.Spot) (*model.Spot, error)
	List(context.Context) ([]*model.Spot, error)
	GetRandom(context.Context) ([]*model.Spot, error)
}
