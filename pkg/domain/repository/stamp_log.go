package repository

import (
	"context"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
	"github.com/google/uuid"
)

type StampLog interface {
	Create(context.Context, *model.StampLog) (*model.StampLog, error)
	List(context.Context, string) ([]*model.StampLog, error)
	GetVisitCnt(ctx context.Context,userID string, spotID uuid.UUID) (uint,error)
}
