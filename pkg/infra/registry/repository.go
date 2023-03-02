package registry

import (
	"github.com/jmoiron/sqlx"

	"github.com/Doer-org/miyagi2023-server/pkg/infra/mysql"
	"github.com/Doer-org/miyagi2023-server/pkg/infra/persistence"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository() (*Repository, error) {
	db, err := mysql.Init()
	if err != nil {
		return nil, err
	}
	return &Repository{db: db}, nil
}

func (r *Repository) NewUser() *persistence.User {
	return persistence.NewUser(r.db)
}
