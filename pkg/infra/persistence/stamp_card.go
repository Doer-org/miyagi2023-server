package persistence

import (
	"context"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type StampCard struct {
	db *sqlx.DB
}

func NewStampCard(db *sqlx.DB) *StampCard {
	return &StampCard{db: db}
}

func (d *StampCard) Get(context.Context, uuid.UUID) (*model.StampCard, error) {
	return nil, nil
}

func (d *StampCard) Create(context.Context, *model.Spot) (*model.Spot, error) {
	return nil, nil
}

func (d *StampCard) List(context.Context) ([]*model.Spot, error) {
	return nil, nil
}
