package persistence

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
)

type User struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *User {
	return &User{db: db}
}

func (d *User) Get(ctx context.Context, id int) (*model.User, error) {
	return d.get(ctx, id)
}

func (d *User) Create(ctx context.Context, user *model.User) (*model.User, error) {
	return nil, nil
}

func (d *User) get(ctx context.Context, id int) (*model.User, error) {
	return nil, nil
}
