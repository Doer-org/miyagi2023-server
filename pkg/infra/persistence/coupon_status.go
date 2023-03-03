package persistence

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
	"github.com/Doer-org/miyagi2023-server/pkg/infra/dto"
)

type CouponStatus struct {
	db *sqlx.DB
}

func NewCouponStatus(db *sqlx.DB) *CouponStatus {
	return &CouponStatus{db: db}
}

func (d *CouponStatus) Get(context.Context, uuid.UUID) (*model.CouponStatus, error) {
	// 今回は使わない
	return nil, nil
}

func (d *CouponStatus) Create(ctx context.Context, couponStatus *model.CouponStatus) (*model.CouponStatus, error) {
	dto := dto.NewCouponStatusDotFromModel(couponStatus)

	query :=
		"INSERT INTO coupon_statuses (`id`,`used_flg`,`coupon_id`,`user_id`)" +
			"VALUES (:id, :used_flg, :coupon_id, :user_id)"

	_, err := d.db.NamedExecContext(ctx, query, dto)
	if err != nil {
		return nil, err
	}

	return d.get(
		ctx, 
		couponStatus.ID,
		couponStatus.Coupon.ID,
		couponStatus.User.ID,
	)
}

func (d *CouponStatus) Use(ctx context.Context, couponID uuid.UUID, userID uuid.UUID) (*model.CouponStatus, error) {
	
	return nil, nil
}

func (d *CouponStatus) ListByUserID(ctx context.Context, couponID uuid.UUID, userID uuid.UUID) (*model.CouponStatus, error) {
	
	return nil, nil
}

func (d *CouponStatus) get(ctx context.Context, id,couponID uuid.UUID,userID model.UserID) (*model.CouponStatus, error) {
	var dtoCouponStatus dto.CouponStatus

	queryCouponStatus := `
		SELECT * FROM coupon_statuses WHERE id = ? LIMIT 1
	`
	err := d.db.GetContext(ctx, &dtoCouponStatus, queryCouponStatus, id.String())
	if err != nil {
		return nil, err
	}

	var dtoCoupon dto.Coupon
	queryCoupon := `
		SELECT * FROM coupons WHERE id = ? LIMIT 1
	`
	err = d.db.GetContext(ctx, &dtoCoupon, queryCoupon, couponID.String())
	if err != nil {
		return nil, err
	}

	var dtoUser dto.User
	queryUser := `
		SELECT * FROM users WHERE id = ? LIMIT 1
	`
	err = d.db.GetContext(ctx, &dtoUser, queryUser, userID.String())
	if err != nil {
		return nil, err
	}

	return dtoCouponStatus.ToModel(&dtoCoupon,&dtoUser)
}
