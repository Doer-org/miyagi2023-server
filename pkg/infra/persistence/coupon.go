package persistence

import (
	"context"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
	"github.com/Doer-org/miyagi2023-server/pkg/infra/dto"
	"github.com/google/uuid"

	"github.com/jmoiron/sqlx"
)

type Coupon struct {
	db *sqlx.DB
}

func NewCoupon(db *sqlx.DB) *Coupon {
	return &Coupon{db: db}
}

func (d *Coupon) Get(context.Context, uuid.UUID) (*model.Coupon, error) {
	return nil, nil
}

func (d *Coupon) Create(context.Context, uuid.UUID) (*model.Coupon, error) {
	return nil, nil
}

func (d *Coupon) List(context.Context) ([]*model.Coupon, error) {
	return nil, nil
}

func (d *Coupon) GetBySpotID(ctx context.Context,spotID uuid.UUID) (*model.Coupon,error) {
	var dtoCoupon dto.Coupon

	//FIXME: これjoinでできる
	query := `
		SELECT * FROM coupons WHERE spot_id = ? LIMIT 1
	`

	err := d.db.GetContext(ctx, &dtoCoupon, query, spotID.String())
	if err != nil {
		return nil, err
	}

	var dtoSpot dto.Spot
	querySpot := `
		SELECT * FROM spots WHERE id = ? LIMIT 1
	`
	err = d.db.GetContext(ctx, &dtoSpot, querySpot, spotID.String())
	if err != nil {
		return nil, err
	}

	return dtoCoupon.ToModel(&dtoSpot)
}
