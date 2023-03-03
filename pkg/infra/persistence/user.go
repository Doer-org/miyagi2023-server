package persistence

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
	"github.com/Doer-org/miyagi2023-server/pkg/infra/dto"
)

type User struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *User {
	return &User{db: db}
}

func (d *User) Get(ctx context.Context, id string) (*model.User, error) {
	return d.get(ctx, id)
}

func (d *User) Create(ctx context.Context, user *model.User) (*model.User, error) {
	dto := dto.NewUserDtoFromModel(user)

	query := `
		INSERT INTO users (id,name,age,gender,birthday,address,profile_img,prefecture) 
		VALUES (:id, :name, :age, :gender, :birthday, :address, :profile_img, :prefecture);
	`

	_, err := d.db.NamedExecContext(ctx, query, dto)
	if err != nil {
		return nil, err
	}

	return d.get(ctx,user.ID.String())
}

func (d *User) get(ctx context.Context, id string) (*model.User, error) {
	var dto dto.User

	query := `
		SELECT * FROM users WHERE id = ? LIMIT 1
	`

	err := d.db.GetContext(ctx, &dto, query, id)
	if err != nil {
		return nil, err
	}

	return dto.ToModel(), nil
}
