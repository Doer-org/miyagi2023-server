package mock

import (
	"time"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
	"github.com/Doer-org/miyagi2023-server/pkg/usecase"
)

var user *model.User = &model.User{
	ID:         "0000",
	Name:       "mahiro mock",
	Age:        22,
	Gender:     model.MEN,
	Birthday:   time.Date(2000, 4, 29, 0, 0, 0, 0, time.Local),
	Address:    "hoge",
	ProfilIMG:  "http://example.com",
	Prefecture: model.TOYAMA,
	CreatedAt:  time.Now(),
}

func NewUserGetOutput() *usecase.UserGetOutput {
	return &usecase.UserGetOutput{
		User: user,
	}
}

func NewUserCreateOutput() *usecase.UserCreateOutput {
	return &usecase.UserCreateOutput{
		User: user,
	}
}
