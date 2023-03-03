package dto

import (
	"time"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
	"github.com/google/uuid"
)

type StampLog struct {
	ID        string    `db:"id"`
	SpotID    string    `db:"spot_id"`
	UserID    string    `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
}

func NewStampLogFromModel(m *model.StampLog) *StampLog {
	return &StampLog{
		ID:     m.ID.String(),
		SpotID: m.Spot.ID.String(),
		UserID: m.User.ID.String(),
	}
}

func (d *StampLog) ToModel(dtoSpot *Spot, dtoUser *User) (*model.StampLog, error) {
	id, err := uuid.Parse(d.ID)
	if err != nil {
		return nil, err
	}

	spot, err := dtoSpot.ToModel()
	if err != nil {
		return nil, err
	}

	return &model.StampLog{
		ID:   id,
		Spot: spot,
		User: dtoUser.ToModel(),
	}, nil
}
