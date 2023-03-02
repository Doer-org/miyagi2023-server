package usecase

import (
	"context"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
)

type User interface {
	Get(context.Context, *UserGetInput) (*UserGetOutput, error)
	Create(context.Context, *UserCreateInput) (*UserCreateOutput, error)
}

type UserGetInput struct {
	ID string
}

type UserGetOutput struct {
	User *model.User
}

type UserCreateInput struct {
	ID         string
	Name       string
	Age        uint
	Gender     string
	Birthday   string
	Address    string
	ProfilIMG  string
	Prefecture string
}

type UserCreateOutput struct {
	User *model.User
}
