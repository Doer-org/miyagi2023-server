package repository

import (
	"context"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
)

type User interface {
	Get(context.Context, model.UserID) (*model.User, error)
	Create(context.Context, *model.User) (*model.User, error)
}
