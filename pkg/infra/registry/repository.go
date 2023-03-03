package registry

import (
	"github.com/jmoiron/sqlx"

	"github.com/Doer-org/miyagi2023-server/pkg/env"
	"github.com/Doer-org/miyagi2023-server/pkg/infra/mysql"
	"github.com/Doer-org/miyagi2023-server/pkg/infra/persistence"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository() (*Repository, error) {
	if env.CONFIG.ENVIRONMENT == "mock" {
		return &Repository{db: nil},nil
	}
	
	db, err := mysql.Init()
	if err != nil {
		return nil, err
	}
	return &Repository{db: db}, nil
}

func (r *Repository) NewUser() *persistence.User {
	return persistence.NewUser(r.db)
}

func (r *Repository) NewSpot() *persistence.Spot {
	return persistence.NewSpot(r.db)
}

func (r *Repository) NewStampLog() *persistence.StampLog {
	return persistence.NewStampLog(r.db)
}

func (r *Repository) NewCoupon() *persistence.Coupon {
	return persistence.NewCoupon(r.db)
}

func (r *Repository) NewCouponStatus() *persistence.CouponStatus {
	return persistence.NewCouponStatus(r.db)
}
