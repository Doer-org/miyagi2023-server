package persistence

import (
	"context"
	"math/rand"
	"time"

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

func (d *Spot) GetRandom(ctx context.Context) ([]*model.Spot, error) {
	var dtos []*dto.Spot

	query := `
		SELECT * FROM spots
	`
	err := d.db.SelectContext(ctx, &dtos, query)
	if err != nil {
		return nil, err
	}

	var resspots []*model.Spot
	sizeart := len(dtos)
	used := make([]int, sizeart+1, sizeart+1)

	for i := 0; i <= sizeart; i++ {
		used[i] = -1
	}

	for i := 0; i < 5; i++ {
		for {
			rand.Seed(time.Now().UnixNano())
			randomNumber := rand.Intn(sizeart)
			if used[randomNumber] == -1 {
				used[randomNumber] = 1
				dtospot := dtos[randomNumber]
				resspot, err := dtospot.ToModel()
				if err != nil {
					return nil, err
				}
				resspots = append(resspots, resspot)
				break
			}
		}
	}

	return resspots, nil

}
