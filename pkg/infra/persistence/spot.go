package persistence

import (
	"context"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
	"github.com/Doer-org/miyagi2023-server/pkg/infra/dto"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Spot struct {
	db *sqlx.DB
}

func NewSpot(db *sqlx.DB) *Spot {
	return &Spot{db: db}
}

func (d *Spot) Get(ctx context.Context, id uuid.UUID) (*model.Spot, error) {
	return d.get(ctx, id)
}

func (d *Spot) Create(ctx context.Context, spot *model.Spot) (*model.Spot, error) {
	dto := dto.NewSpotDotFromModel(spot)

	query :=
		"INSERT INTO spots (`id`,`name`,`detail`,`like`,`img_url`,`address`)" +
		"VALUES (:id, :name, :detail, :like, :img_url, :address)"

	_, err := d.db.NamedExecContext(ctx, query, dto)
	if err != nil {
		return nil, err
	}

	return d.get(ctx, spot.ID)
}

func (d *Spot) List(ctx context.Context) ([]*model.Spot, error) {
	var dtos []*dto.Spot

	query := `
		SELECT * FROM spots
	`

	err := d.db.SelectContext(ctx, &dtos, query)
	if err != nil {
		return nil, err
	}

	var spots []*model.Spot
	for _, dto := range dtos {
		spot, err := dto.ToModel()
		if err != nil {
			return nil, err
		}
		spots = append(spots, spot)
	}

	return spots, nil
}

func (d *Spot) get(ctx context.Context, id uuid.UUID) (*model.Spot, error) {
	var dto dto.Spot

	query := `
		SELECT * FROM spots WHERE id = ? LIMIT 1
	`

	err := d.db.GetContext(ctx, &dto, query, id.String())
	if err != nil {
		return nil, err
	}

	return dto.ToModel()
}
