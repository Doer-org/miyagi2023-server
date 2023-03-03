package dto

import (
	"time"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
	"github.com/google/uuid"
)

type CouponStatus struct {
	ID        string `db:"id"`
	UsedFlg   bool `db:"used_flg"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	CouponID string  `db:"coupon_id"`
	UserID      string `db:"user_id"`
}

func NewCouponStatusDotFromModel(m *model.CouponStatus) *CouponStatus {
	return &CouponStatus{
		ID:        m.ID.String(),
		UsedFlg:      m.UsedFlg,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		CouponID: m.Coupon.ID.String(),
		UserID: m.User.ID.String(),
	}
}

func (d *CouponStatus) ToModel(dtoCoupon *Coupon,dtoUser *User) (*model.CouponStatus, error) {
	id, err := uuid.Parse(d.ID)
	if err != nil {
		return nil, err
	}

	coupon,err := dtoCoupon.ToModel(nil)
	if err != nil {
		return nil, err
	}

	return &model.CouponStatus{
		ID:        id,
		UsedFlg:      d.UsedFlg,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
		Coupon: coupon,
		User: dtoUser.ToModel(),
	},nil
}
