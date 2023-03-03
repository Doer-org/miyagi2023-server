package dto

import (
	"time"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
	"github.com/google/uuid"
)

type Spot struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Detail    string    `db:"detail"`
	Like      uint      `db:"like"`
	ImgURL    string    `db:"img_url"`
	Address   string    `db:"address"`
	CreatedAt time.Time `db:"created_at"`
}

func NewSpotDotFromModel(m *model.Spot) *Spot {
	if m == nil {
		return nil
	}
	return &Spot{
		ID:        m.ID.String(),
		Name:      m.Name,
		Detail:    m.Detail,
		Like:      m.Like,
		ImgURL:    m.ImgURL,
		Address:   m.Address,
		CreatedAt: m.CreatedAt,
	}
}

func (d *Spot) ToModel() (*model.Spot, error) {
	id, err := uuid.Parse(d.ID)
	if err != nil {
		return nil, err
	}

	return &model.Spot{
		ID:        id,
		Name:      d.Name,
		Detail:    d.Detail,
		Like:      d.Like,
		ImgURL:    d.ImgURL,
		Address:   d.Address,
		CreatedAt: d.CreatedAt,
	}, nil
}
