package dto

import (
	"time"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
)

type User struct {
	ID         string    `db:"id"`
	Name       string    `db:"name"`
	Age        uint      `db:"age"`
	Gender     string    `db:"gender"`
	Birthday   time.Time `db:"birthday"`
	Address    string    `db:"address"`
	ImgURL string    `db:"img_url"`
	Prefecture string    `db:"prefecture"`
	CreatedAt  time.Time `db:"created_at"`
}

func NewUserDtoFromModel(m *model.User) *User {
	return &User{
		ID:         m.ID.String(),
		Name:       m.Name,
		Age:        m.Age,
		Gender:     m.Gender.String(),
		Birthday:   m.Birthday,
		Address:    m.Address,
		ImgURL: m.ImgURL,
		Prefecture: m.Prefecture.String(),
		CreatedAt:  m.CreatedAt,
	}
}

func (d *User) ToModel() *model.User {
	return &model.User{
		ID:         model.UserID(d.ID),
		Name:       d.Name,
		Age:        d.Age,
		Gender:     model.Gender(d.Gender),
		Birthday:   d.Birthday,
		Address:    d.Address,
		ImgURL: d.ImgURL,
		Prefecture: model.Prefecture(d.Prefecture),
		CreatedAt:  d.CreatedAt,
	}
}
