package persistence

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
)

type StampLog struct {
	db *sqlx.DB
}

func NewStampLog(db *sqlx.DB) *StampLog {
	return &StampLog{db: db}
}

func (d *StampLog) Create(context.Context, *model.StampLog) (*model.StampLog, error) {
	return nil, nil
}

func (d *StampLog) List(context.Context, string) ([]*model.StampLog, error) {
	return nil, nil
}
