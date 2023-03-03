package interactor

import (
	"context"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/repository"
	"github.com/Doer-org/miyagi2023-server/pkg/interactor/mock"
	"github.com/Doer-org/miyagi2023-server/pkg/usecase"
)

type User struct {
	repository repository.User
}

func NewUser(repository repository.User) *User {
	return &User{repository: repository}
}

func (uc *User) Get(ctx context.Context, in *usecase.UserGetInput) (*usecase.UserGetOutput, error) {
	return mock.NewUserGetOutput(),nil
}

func (uc *User) Create(ctx context.Context, in *usecase.UserCreateInput) (*usecase.UserCreateOutput, error) {
	return mock.NewUserCreateOutput(),nil
}
