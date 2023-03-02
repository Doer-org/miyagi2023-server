package persistence

import (
	"context"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Spot struct {
	db *sqlx.DB
}

func NewSpot(db *sqlx.DB) *Spot {
	return &Spot{db: db}
}

func (d *Spot) Get(context.Context, uuid.UUID) (*model.Spot, error) {
	return nil, nil
}

func (d *Spot) Create(context.Context, *model.Spot) (*model.Spot, error) {
	return nil, nil
}

func (d *Spot) List(context.Context, uuid.UUID) ([]*model.Spot, error) {
	return nil, nil
}
