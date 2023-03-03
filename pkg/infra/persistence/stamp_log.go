package persistence

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
	"github.com/Doer-org/miyagi2023-server/pkg/infra/dto"
)

type StampLog struct {
	db *sqlx.DB
}

func NewStampLog(db *sqlx.DB) *StampLog {
	return &StampLog{db: db}
}

func (d *StampLog) Create(ctx context.Context, stampLog *model.StampLog) (*model.StampLog, error) {
	dto := dto.NewStampLogFromModel(stampLog)

	query :=
		"INSERT INTO stamp_logs (`id`,`spot_id`,`user_id`)" +
			"VALUES (:id, :spot_id, :user_id)"

	_, err := d.db.NamedExecContext(ctx, query, dto)
	if err != nil {
		return nil, err
	}

	return d.get(ctx, stampLog)
}

func (d *StampLog) get(ctx context.Context, stampLog *model.StampLog) (*model.StampLog, error) {
	var dtoUser dto.User
	queryUser := `
		SELECT * FROM users WHERE id = ? LIMIT 1
	`
	err := d.db.GetContext(ctx, &dtoUser, queryUser, stampLog.User.ID.String())
	if err != nil {
		return nil, err
	}

	var dtoSpot dto.Spot
	querySpot := `
		SELECT * FROM spots WHERE id = ? LIMIT 1
	`
	err = d.db.GetContext(ctx, &dtoSpot, querySpot, stampLog.Spot.ID.String())
	if err != nil {
		return nil, err
	}

	var dtoStampLog dto.StampLog
	queryStampLog := `
		SELECT * FROM stamp_logs WHERE id = ? LIMIT 1
	`
	err = d.db.GetContext(ctx, &dtoSpot, queryStampLog, stampLog.ID.String())
	if err != nil {
		return nil, err
	}
	return dtoStampLog.ToModel(&dtoSpot, &dtoUser)
}

func (d *StampLog) List(context.Context, string) ([]*model.StampLog, error) {
	return nil, nil
}

func (d *StampLog) GetVisitCnt(ctx context.Context, userID string, spotID uuid.UUID) (uint, error) {
	// TODO:yadon ログを取得して訪れた回数を返す
	var dtos []dto.StampLog
	query := `
		SELECT * FROM stamp_logs WHERE user_id = ? AND spot_id = ?
	`
	err := d.db.SelectContext(ctx, &dtos, query, userID, spotID)
	if err != nil {
		return 0, err
	}

	return uint(len(dtos)), nil
}
